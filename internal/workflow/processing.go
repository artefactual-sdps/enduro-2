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
	temporalsdk_temporal "go.temporal.io/sdk/temporal"
	temporalsdk_workflow "go.temporal.io/sdk/workflow"

	"github.com/artefactual-labs/enduro/internal/a3m"
	"github.com/artefactual-labs/enduro/internal/collection"
	sdps_activities "github.com/artefactual-labs/enduro/internal/sdps/activities"
	"github.com/artefactual-labs/enduro/internal/temporal"
	"github.com/artefactual-labs/enduro/internal/validation"
	"github.com/artefactual-labs/enduro/internal/watcher"
	"github.com/artefactual-labs/enduro/internal/workflow/activities"
)

type ProcessingWorkflow struct {
	logger logr.Logger
	colsvc collection.Service
	wsvc   watcher.Service
}

func NewProcessingWorkflow(logger logr.Logger, colsvc collection.Service, wsvc watcher.Service) *ProcessingWorkflow {
	return &ProcessingWorkflow{
		logger: logger,
		colsvc: colsvc,
		wsvc:   wsvc,
	}
}

// TransferInfo is shared state that is passed down to activities. It can be
// useful for hooks that may require quick access to processing state.
// TODO: clean this up, e.g.: it can embed a collection.Collection.
type TransferInfo struct {
	// TempFile is the temporary location where the blob is downloaded.
	//
	// It is populated by the workflow with the result of DownloadActivity.
	TempFile string

	// SIPID given by a3m.
	//
	// It is populated by CreateAIPActivity.
	SIPID string

	// Enduro internal collection ID.
	// The zero value represents a new collection. It can be used to indicate
	// an existing collection in retries.
	//
	// It is populated via the workflow request or createPackageLocalActivity.
	CollectionID uint

	// Name of the watcher that received this blob.
	//
	// It is populated via the workflow request. Expect an empty string when
	// the workflow was started by a batch.
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

	// Batch directory that contains the blob.
	//
	// It is populated via the workflow request.
	BatchDir string

	// StoredAt is the time when the AIP is stored.
	//
	// It is populated by PollIngestActivity as long as Ingest completes.
	StoredAt time.Time

	// Information about the bundle (transfer) that we submit to Archivematica.
	// Full path, relative path...
	//
	// It is populated by BundleActivity.
	Bundle activities.BundleActivityResult
}

// ProcessingWorkflow orchestrates all the activities related to the processing
// of a SIP in Archivematica, including is retrieval, creation of transfer,
// etc...
//
// Retrying this workflow would result in a new Archivematica transfer. We  do
// not have a retry policy in place. The user could trigger a new instance via
// the API.
func (w *ProcessingWorkflow) Execute(ctx temporalsdk_workflow.Context, req *collection.ProcessingWorkflowRequest) error {
	var (
		logger = temporalsdk_workflow.GetLogger(ctx)

		tinfo = &TransferInfo{
			CollectionID:     req.CollectionID,
			WatcherName:      req.WatcherName,
			RetentionPeriod:  req.RetentionPeriod,
			CompletedDir:     req.CompletedDir,
			StripTopLevelDir: req.StripTopLevelDir,
			Key:              req.Key,
			IsDir:            req.IsDir,
			BatchDir:         req.BatchDir,
		}

		// Collection status. All collections start in queued status.
		status = collection.StatusQueued
	)

	// Persist collection as early as possible.
	{
		activityOpts := withLocalActivityOpts(ctx)
		var err error

		if req.CollectionID == 0 {
			err = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, createPackageLocalActivity, w.logger, w.colsvc, &createPackageLocalActivityParams{
				Key:    req.Key,
				Status: status,
			}).Get(activityOpts, &tinfo.CollectionID)
		} else {
			// TODO: investigate better way to reset the collection.
			err = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.colsvc, &updatePackageLocalActivityParams{
				CollectionID: req.CollectionID,
				Key:          req.Key,
				SIPID:        "",
				StoredAt:     time.Time{},
				Status:       status,
			}).Get(activityOpts, nil)
		}

		if err != nil {
			return fmt.Errorf("error persisting collection: %v", err)
		}
	}

	// Ensure that the status of the collection is always updated when this
	// workflow function returns.
	defer func() {
		// Mark as failed unless it completed successfully or it was abandoned.
		if status != collection.StatusDone && status != collection.StatusAbandoned {
			status = collection.StatusError
		}

		// Use disconnected context so it also runs after cancellation.
		dctx, _ := temporalsdk_workflow.NewDisconnectedContext(ctx)
		activityOpts := withLocalActivityOpts(dctx)
		_ = temporalsdk_workflow.ExecuteLocalActivity(activityOpts, updatePackageLocalActivity, w.logger, w.colsvc, &updatePackageLocalActivityParams{
			CollectionID: tinfo.CollectionID,
			Key:          tinfo.Key,
			SIPID:        tinfo.SIPID,
			StoredAt:     tinfo.StoredAt,
			Status:       status,
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

			// We use this timer to identify transfers that exceeded a deadline.
			// We can't rely on workflow.ErrCanceled because the same context
			// error is seen when the session worker dies.
			timer := NewTimer()

			sessErr = w.SessionHandler(sessCtx, attempt, tinfo, req.ValidationConfig, timer)

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
					"err", ctx.Err().Error(),
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

		status = collection.StatusDone
	}

	// Schedule deletion of the original in the watched data source.
	{
		if status == collection.StatusDone {
			if tinfo.RetentionPeriod != nil {
				err := temporalsdk_workflow.NewTimer(ctx, *tinfo.RetentionPeriod).Get(ctx, nil)
				if err != nil {
					logger.Warn("Retention policy timer failed", "err", err.Error())
				} else {
					activityOpts := withActivityOptsForRequest(ctx)
					_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DeleteOriginalActivityName, tinfo.WatcherName, tinfo.BatchDir, tinfo.Key).Get(activityOpts, nil)
				}
			} else if tinfo.CompletedDir != "" {
				activityOpts := withActivityOptsForLocalAction(ctx)
				_ = temporalsdk_workflow.ExecuteActivity(activityOpts, activities.DisposeOriginalActivityName, tinfo.WatcherName, tinfo.CompletedDir, tinfo.BatchDir, tinfo.Key).Get(activityOpts, nil)
			}
		}
	}

	logger.Info(
		"Workflow completed successfully!",
		"collectionID", tinfo.CollectionID,
		"watcher", tinfo.WatcherName,
		"batchDir", tinfo.BatchDir,
		"key", tinfo.Key,
		"status", status.String(),
	)

	return nil
}

// SessionHandler runs activities that belong to the same session.
func (w *ProcessingWorkflow) SessionHandler(sessCtx temporalsdk_workflow.Context, attempt int, tinfo *TransferInfo, validationConfig validation.Config, timer *Timer) error {
	defer temporalsdk_workflow.CompleteSession(sessCtx)

	// Set in-progress status.
	{
		ctx := withLocalActivityOpts(sessCtx)
		err := temporalsdk_workflow.ExecuteLocalActivity(ctx, setStatusInProgressLocalActivity, w.colsvc, tinfo.CollectionID, time.Now().UTC()).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	// Download.
	{
		if tinfo.WatcherName != "" && !tinfo.IsDir {
			// TODO: even if TempFile is defined, we should confirm that the file is
			// locally available in disk, just in case we're in the context of a
			// session retry where a different working is doing the work. In that
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
				BatchDir:         tinfo.BatchDir,
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

	// Validate package.
	{
		if tinfo.Bundle != (activities.BundleActivityResult{}) {
			activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
				StartToCloseTimeout: time.Minute * 5,
				RetryPolicy: &temporalsdk_temporal.RetryPolicy{
					MaximumAttempts: 1,
				},
			})
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, sdps_activities.ValidatePackageActivityName, tinfo.Bundle.FullPath).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}
	}

	// Validate transfer.
	{
		if validationConfig.IsEnabled() && tinfo.Bundle != (activities.BundleActivityResult{}) {
			activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
				StartToCloseTimeout: time.Minute * 5,
				RetryPolicy: &temporalsdk_temporal.RetryPolicy{
					MaximumAttempts: 1,
				},
			})
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.ValidateTransferActivityName, &activities.ValidateTransferActivityParams{
				Config: validationConfig,
				Path:   tinfo.Bundle.FullPath,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}
	}

	{
		err := w.transferA3m(sessCtx, tinfo)
		if err != nil {
			return err
		}
	}

	// Upload AIP to MinIO.
	{
		// tinfo.AIPPath
		activityOpts := temporalsdk_workflow.WithActivityOptions(sessCtx, temporalsdk_workflow.ActivityOptions{
			StartToCloseTimeout: time.Hour * 24,
			RetryPolicy: &temporalsdk_temporal.RetryPolicy{
				InitialInterval:    time.Second,
				BackoffCoefficient: 2,
				MaximumAttempts:    3,
			},
		})
		err := temporalsdk_workflow.ExecuteActivity(activityOpts, activities.UploadActivityName, tinfo.AIPPath).Get(activityOpts, nil)
		if err != nil {
			return err
		}
	}

	// Index content in OpenSearch.
	{
		if tinfo.Bundle != (activities.BundleActivityResult{}) {
			activityOpts := withActivityOptsForLongLivedRequest(sessCtx)
			err := temporalsdk_workflow.ExecuteActivity(activityOpts, sdps_activities.IndexActivityName, &sdps_activities.IndexActivityParams{
				Path:        tinfo.Bundle.FullPath,
				SearchIndex: sdps_activities.ESIndexName,
			}).Get(activityOpts, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
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
		Path:         tinfo.Bundle.FullPath,
		CollectionID: tinfo.CollectionID,
	}

	result := a3m.CreateAIPActivityResult{}
	err := temporalsdk_workflow.ExecuteActivity(activityOpts, a3m.CreateAIPActivityName, params).Get(sessCtx, &result)

	tinfo.SIPID = result.UUID
	tinfo.AIPPath = result.Path
	tinfo.StoredAt = temporalsdk_workflow.Now(sessCtx).UTC()

	return err
}
