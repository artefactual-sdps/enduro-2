// Package workflow contains an experimental workflow for Archivemica transfers.
//
// It's not generalized since it contains client-specific activities. However,
// the long-term goal is to build a system where workflows and activities are
// dynamically set up based on user input.
package workflow

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	temporalsdk_temporal "go.temporal.io/sdk/temporal"
	temporalsdk_workflow "go.temporal.io/sdk/workflow"

	"github.com/artefactual-sdps/enduro/internal/a3m"
	"github.com/artefactual-sdps/enduro/internal/package_"
	"github.com/artefactual-sdps/enduro/internal/ref"
	"github.com/artefactual-sdps/enduro/internal/temporal"
	"github.com/artefactual-sdps/enduro/internal/watcher"
	"github.com/artefactual-sdps/enduro/internal/workflow/activities"
)

type ProcessingWorkflow struct {
	logger logr.Logger
	pkgsvc package_.Service
	wsvc   watcher.Service
}

func NewProcessingWorkflow(logger logr.Logger, pkgsvc package_.Service, wsvc watcher.Service) *ProcessingWorkflow {
	return &ProcessingWorkflow{
		logger: logger,
		pkgsvc: pkgsvc,
		wsvc:   wsvc,
	}
}

// TransferInfo is shared state that is passed down to activities. It can be
// useful for hooks that may require quick access to processing state.
// TODO: clean this up, e.g.: it can embed a package_.package_.
type TransferInfo struct {
	// TempFile is the temporary location where the blob is downloaded.
	//
	// It is populated by the workflow with the result of DownloadActivity.
	TempFile string

	// SIPID given by a3m.
	//
	// It is populated by CreateAIPActivity.
	SIPID string

	// Enduro internal package ID.
	// The zero value represents a new package_. It can be used to indicate
	// an existing package in retries.
	//
	// It is populated via the workflow request or createPackageLocalActivity.
	PackageID uint

	// Name of the watcher that received this blob.
	//
	// It is populated via the workflow request.
	WatcherName string

	// Retention period.
	// Period of time to schedule the deletion of the original blob from the
	// watched data source. nil means no deletion.
	//
	// It is populated via the workflow request.
	RetentionPeriod *time.Duration

	// Directory where the transfer is moved to once processing has completed
	// successfully.
	//
	// It is populated via the workflow request.
	CompletedDir string

	// Path to the compressed AIP generated by the preservation system.
	//
	// It is populated once the preservation system creates the AIP.
	AIPPath string

	// Whether the top-level directory is meant to be stripped.
	//
	// It is populated via the workflow request.
	StripTopLevelDir bool

	// Key of the blob.
	//
	// It is populated via the workflow request.
	Key string

	// Whether the blob is a directory (fs watcher)
	//
	// It is populated via the workflow request.
	IsDir bool

	// StoredAt is the time when the AIP is stored.
	//
	// It is populated by PollIngestActivity as long as Ingest completes.
	StoredAt time.Time

	// Information about the bundle (transfer) that we submit to Archivematica.
	// Full path, relative path...
	//
	// It is populated by BundleActivity.
	Bundle activities.BundleActivityResult

	// Identifier of the preservation action that creates the AIP
	//
	// It is populated by createPreservationActionLocalActivity .
	PreservationActionID uint
}

// ProcessingWorkflow orchestrates all the activities related to the processing
// of a SIP in Archivematica, including is retrieval, creation of transfer,
// etc...
//
// Retrying this workflow would result in a new Archivematica transfer. We  do
// not have a retry policy in place. The user could trigger a new instance via
// the API.
func (w *ProcessingWorkflow) Execute(ctx temporalsdk_workflow.Context, req *package_.ProcessingWorkflowRequest) error {
	var (
		logger = temporalsdk_workflow.GetLogger(ctx)

		tinfo = &TransferInfo{
			PackageID:        req.PackageID,
			WatcherName:      req.WatcherName,
			RetentionPeriod:  req.RetentionPeriod,
			CompletedDir:     req.CompletedDir,
			StripTopLevelDir: req.StripTopLevelDir,
			Key:              req.Key,
			IsDir:            req.IsDir,
		}

		// Package status. All packages start in queued status.
		status = package_.StatusQueued

		// Create AIP preservation action status.
		paStatus = package_.ActionStatusUnspecified
	)

	// Persist package as early as possible.
	{
		activityOpts := withLocalActivityOpts(ctx)
		var err error

		if req.PackageID == 0 {
			err = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, createPackageLocalActivity, w.logger, w.pkgsvc, &createPackageLocalActivityParams{
				Key:    req.Key,
				Status: status,
			}).Get(activityOpts, &tinfo.PackageID)
		} else {
			// TODO: investigate better way to reset the package_.
			err = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.pkgsvc, &updatePackageLocalActivityParams{
				PackageID: req.PackageID,
				Key:       req.Key,
				SIPID:     "",
				StoredAt:  temporalsdk_workflow.Now(ctx).UTC(),
				Status:    status,
			}).Get(activityOpts, nil)
		}

		if err != nil {
			return fmt.Errorf("error persisting package: %v", err)
		}
	}

	// Ensure that the status of the package and the preservation action is always updated when this
	// workflow function returns.
	defer func() {
		// Mark as failed unless it completed successfully or it was abandoned.
		if status != package_.StatusDone && status != package_.StatusAbandoned {
			status = package_.StatusError
		}

		// Use disconnected context so it also runs after cancellation.
		dctx, _ := temporalsdk_workflow.NewDisconnectedContext(ctx)
		activityOpts := withLocalActivityOpts(dctx)
		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.pkgsvc, &updatePackageLocalActivityParams{
			PackageID: tinfo.PackageID,
			Key:       tinfo.Key,
			SIPID:     tinfo.SIPID,
			StoredAt:  tinfo.StoredAt,
			Status:    status,
		}).Get(activityOpts, nil)

		if paStatus != package_.ActionStatusDone {
			paStatus = package_.ActionStatusError
		}

		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, completePreservationActionLocalActivity, w.pkgsvc, &completePreservationActionLocalActivityParams{
			PreservationActionID: tinfo.PreservationActionID,
			Status:               paStatus,
			CompletedAt:          temporalsdk_workflow.Now(dctx).UTC(),
		}).Get(activityOpts, nil)
	}()

	// Activities running within a session.
	{
		var sessErr error
		maxAttempts := 5

		for attempt := 1; attempt <= maxAttempts; attempt++ {
			activityOpts := temporalsdk_workflow.WithActivityOptions(ctx, temporalsdk_workflow.ActivityOptions{
				StartToCloseTimeout: time.Minute,
				TaskQueue:           temporal.A3mWorkerTaskQueue,
			})
			sessCtx, err := temporalsdk_workflow.CreateSession(activityOpts, &temporalsdk_workflow.SessionOptions{
				CreationTimeout:  forever,
				ExecutionTimeout: forever,
			})
			if err != nil {
				return fmt.Errorf("error creating session: %v", err)
			}

			sessErr = w.SessionHandler(sessCtx, attempt, tinfo)

			// We want to retry the session if it has been canceled as a result
			// of losing the worker but not otherwise. This scenario seems to be
			// identifiable when we have an error but the root context has not
			// been canceled.
			if sessErr != nil && (errors.Is(sessErr, temporalsdk_workflow.ErrSessionFailed) || temporalsdk_temporal.IsCanceledError(sessErr)) {
				// Root context canceled, hence workflow canceled.
				if ctx.Err() == temporalsdk_workflow.ErrCanceled {
					return nil
				}

				logger.Error(
					"Session failed, will retry shortly (10s)...",
					"err", ctx.Err(),
					"attemptFailed", attempt,
					"attemptsLeft", maxAttempts-attempt,
				)

				_ = temporalsdk_workflow.Sleep(ctx, time.Second*10)

				continue
			}
			break
		}

		if sessErr != nil {
			return sessErr
		}

		status = package_.StatusDone

		paStatus = package_.ActionStatusDone
	}

	// Schedule deletion of the original in the watched data source.
	{
		if status == package_.StatusDone {
			if tinfo.RetentionPeriod != nil {
				err := temporalsdk_workflow.NewTimer(ctx, *tinfo.RetentionPeriod).Get(ctx, nil)
				if err != nil {
					logger.Warn("Retention policy timer failed", "err", err.Error())
				} else {
					activityOpts := withActivityOptsForRequest(ctx)
					_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DeleteOriginalActivityName, tinfo.WatcherName, tinfo.Key).Get(activityOpts, nil)
				}
			} else if tinfo.CompletedDir != "" {
				activityOpts := withActivityOptsForLocalAction(ctx)
				_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DisposeOriginalActivityName, tinfo.WatcherName, tinfo.CompletedDir, tinfo.Key).Get(activityOpts, nil)
			}
		}
	}

	logger.Info(
		"Workflow completed successfully!",
		"packageID", tinfo.PackageID,
		"watcher", tinfo.WatcherName,
		"key", tinfo.Key,
		"status", status.String(),
	)

	return nil
}

// SessionHandler runs activities that belong to the same session.
func (w *ProcessingWorkflow) SessionHandler(sessCtx temporalsdk_workflow.Context, attempt int, tinfo *TransferInfo) error {
	defer temporalsdk_workflow.CompleteSession(sessCtx)

	packageStartedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	// Set in-progress status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusInProgressLocalActivity, w.pkgsvc, tinfo.PackageID, packageStartedAt).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Persist the preservation action for creating the AIP.
	{
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationActionLocalActivity, w.pkgsvc, &createPreservationActionLocalActivityParams{
				WorkflowID: temporalsdk_workflow.GetInfo(ctx).WorkflowExecution.ID,
				Type:       package_.ActionTypeCreateAIP,
				Status:     package_.ActionStatusInProgress,
				StartedAt:  packageStartedAt,
				PackageID:  tinfo.PackageID,
			}).Get(ctx, &tinfo.PreservationActionID)
			if err != nil {
				return err
			}
		}
	}

	// Download.
	{
		if tinfo.WatcherName != "" && !tinfo.IsDir {
			// TODO: even if TempFile is defined, we should confirm that the file is
			// locally available in disk, just in case we're in the context of a
			// session retry where a different worker is doing the work. In that
			// case, the activity whould be executed again.
			if tinfo.TempFile == "" {
				activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
				err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DownloadActivityName, tinfo.WatcherName, tinfo.Key).Get(activityOpts, &tinfo.TempFile)
				if err != nil {
					return err
				}
			}
		}
	}

	// Bundle.
	{
		if tinfo.Bundle == (activities.BundleActivityResult{}) {
			activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.BundleActivityName, &activities.BundleActivityParams{
				WatcherName:      tinfo.WatcherName,
				TransferDir:      "/home/a3m/.local/share/a3m/share",
				Key:              tinfo.Key,
				IsDir:            tinfo.IsDir,
				TempFile:         tinfo.TempFile,
				StripTopLevelDir: tinfo.StripTopLevelDir,
			}).Get(activityOpts, &tinfo.Bundle)
			if err != nil {
				return err
			}
		}
	}

	// Delete local temporary files.
	defer func() {
		// TODO: call clean up here to enforce that we always destroy TempDir.
		if tinfo.Bundle.FullPathBeforeStrip != "" {
			activityOpts := withActivityOptsForRequest(sessCtx)
			_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.CleanUpActivityName, &activities.CleanUpActivityParams{
				FullPath: tinfo.Bundle.FullPathBeforeStrip,
			}).Get(activityOpts, nil)
		}
	}()

	{
		err := w.transferA3m(sessCtx, tinfo)
		if err != nil {
			return err
		}
	}

	// Persist SIPID.
	{
		activityOpts := withLocalActivityOpts(sessCtx)
		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.pkgsvc, &updatePackageLocalActivityParams{
			PackageID: tinfo.PackageID,
			Key:       tinfo.Key,
			SIPID:     tinfo.SIPID,
			StoredAt:  tinfo.StoredAt,
			Status:    package_.StatusInProgress,
		}).Get(activityOpts, nil)
	}

	// Identifier of the preservation task for upload to sips bucket.
	var uploadPreservationTaskID uint

	uploadStartedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	// Add preservation task for upload to review bucket.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationTaskLocalActivity, w.pkgsvc, &createPreservationTaskLocalActivityParams{
			TaskID:               uuid.NewString(),
			Name:                 "Move AIP",
			Status:               package_.TaskStatusInProgress,
			StartedAt:            uploadStartedAt,
			Note:                 "Moving to review bucket",
			PreservationActionID: tinfo.PreservationActionID,
		}).Get(ctx, &uploadPreservationTaskID)
		if err != nil {
			return err
		}
	}

	// Upload AIP to MinIO.
	{
		activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
			StartToCloseTimeout: time.Hour * 24,
			RetryPolicy: &temporalsdk_temporal.RetryPolicy{
				InitialInterval:    time.Second,
				BackoffCoefficient: 2,
				MaximumAttempts:    3,
			},
		})
		err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.UploadActivityName, &activities.UploadActivityParams{
			AIPPath: tinfo.AIPPath,
			AIPID:   tinfo.SIPID,
			Name:    tinfo.Key,
		}).Get(activityOpts, nil)
		if err != nil {
			return err
		}
	}

	uploadCompletedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	// Complete preservation task for upload to review bucket.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
			ID:          uploadPreservationTaskID,
			Status:      package_.TaskStatusDone,
			CompletedAt: uploadCompletedAt,
			Note:        ref.New("Moved to review bucket"),
		}).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Set package to pending status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusLocalActivity, w.pkgsvc, tinfo.PackageID, package_.StatusPending).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Set preservation action to pending status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setPreservationActonStatusLocalActivity, w.pkgsvc, tinfo.PreservationActionID, package_.ActionStatusPending).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Identifier of the preservation task for package review
	var reviewPreservationTaskID uint

	reviewStartedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	// Add preservation task for package review
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationTaskLocalActivity, w.pkgsvc, &createPreservationTaskLocalActivityParams{
			TaskID:               uuid.NewString(),
			Name:                 "Review AIP",
			Status:               package_.TaskStatusPending,
			StartedAt:            reviewStartedAt,
			Note:                 "Awaiting user decision",
			PreservationActionID: tinfo.PreservationActionID,
		}).Get(ctx, &reviewPreservationTaskID)
		if err != nil {
			return err
		}
	}

	review := w.waitForReview(sessCtx)
	reviewCompletedAt := temporalsdk_workflow.Now(sessCtx).UTC()

	// Set package to in progress status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusLocalActivity, w.pkgsvc, tinfo.PackageID, package_.StatusInProgress).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Set preservation action to in progress status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setPreservationActonStatusLocalActivity, w.pkgsvc, tinfo.PreservationActionID, package_.ActionStatusInProgress).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	if review.Accepted {
		// Record package confirmation in review preservation task
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
				ID:          reviewPreservationTaskID,
				Status:      package_.TaskStatusDone,
				CompletedAt: reviewCompletedAt,
				Note:        ref.New("Reviewed and accepted"),
			}).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Identifier of the preservation task for permanent storage move.
		var movePreservationTaskID uint

		moveStartedAt := temporalsdk_workflow.Now(sessCtx).UTC()

		// Add preservation task for permanent storage move.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, createPreservationTaskLocalActivity, w.pkgsvc, &createPreservationTaskLocalActivityParams{
				TaskID:               uuid.NewString(),
				Name:                 "Move AIP",
				Status:               package_.TaskStatusInProgress,
				StartedAt:            moveStartedAt,
				Note:                 "Moving to permanent storage",
				PreservationActionID: tinfo.PreservationActionID,
			}).Get(ctx, &movePreservationTaskID)
			if err != nil {
				return err
			}
		}

		// Move package to permanent storage
		{
			activityOpts := withActivityOptsForRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.MoveToPermanentStorageActivityName, &activities.MoveToPermanentStorageActivityParams{
				AIPID:      tinfo.SIPID,
				LocationID: *review.LocationID,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}

		// Poll package move to permanent storage
		{
			activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.PollMoveToPermanentStorageActivityName, &activities.PollMoveToPermanentStorageActivityParams{
				AIPID: tinfo.SIPID,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}

		moveCompletedAt := temporalsdk_workflow.Now(sessCtx).UTC()

		// Complete preservation task for permanent storage move.
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
				ID:          movePreservationTaskID,
				Status:      package_.TaskStatusDone,
				CompletedAt: moveCompletedAt,
				Note:        ref.New(fmt.Sprintf("Moved to location %s", *review.LocationID)),
			}).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Set package location
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setLocationIDLocalActivity, w.pkgsvc, tinfo.PackageID, *review.LocationID).Get(ctx, nil)
			if err != nil {
				return err
			}
		}
	} else {
		// Record package rejection in review preservation task
		{
			ctx := withLocalActivityOpts(sessCtx)
			err := temporalsdk_workflow.ExecuteLocalActivity(ctx, completePreservationTaskLocalActivity, w.pkgsvc, &completePreservationTaskLocalActivityParams{
				ID:          reviewPreservationTaskID,
				Status:      package_.TaskStatusDone,
				CompletedAt: reviewCompletedAt,
				Note:        ref.New("Reviewed and rejected"),
			}).Get(ctx, nil)
			if err != nil {
				return err
			}
		}

		// Reject package
		{
			activityOpts := withActivityOptsForRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.RejectPackageActivityName, &activities.RejectPackageActivityParams{
				AIPID: tinfo.SIPID,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *ProcessingWorkflow) waitForReview(ctx temporalsdk_workflow.Context) *package_.ReviewPerformedSignal {
	var review package_.ReviewPerformedSignal
	signalChan := temporalsdk_workflow.GetSignalChannel(ctx, package_.ReviewPerformedSignalName)
	selector := temporalsdk_workflow.NewSelector(ctx)
	selector.AddReceive(signalChan, func(channel temporalsdk_workflow.ReceiveChannel, more bool) {
		_ = channel.Receive(ctx, &review)
	})
	selector.Select(ctx)
	return &review
}

func (w *ProcessingWorkflow) transferA3m(sessCtx temporalsdk_workflow.Context, tinfo *TransferInfo) error {
	activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
		StartToCloseTimeout: time.Hour * 24,
		HeartbeatTimeout:    time.Second * 5,
		RetryPolicy: &temporalsdk_temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})

	params := &a3m.CreateAIPActivityParams{
		Path:                 tinfo.Bundle.FullPath,
		PreservationActionID: tinfo.PreservationActionID,
	}

	result := a3m.CreateAIPActivityResult{}
	err := temporalsdk_workflow.ExecuteActivity(activityOpts, a3m.CreateAIPActivityName, params).Get(sessCtx, &result)

	tinfo.SIPID = result.UUID
	tinfo.AIPPath = result.Path
	tinfo.StoredAt = temporalsdk_workflow.Now(sessCtx).UTC()

	return err
}
