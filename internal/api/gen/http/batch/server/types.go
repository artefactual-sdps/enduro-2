// Code generated by goa v3.5.5, DO NOT EDIT.
//
// batch HTTP server types
//
// Command:
// $ goa-v3.5.5 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	batch "github.com/artefactual-labs/enduro/internal/api/gen/batch"
	goa "goa.design/goa/v3/pkg"
)

// SubmitRequestBody is the type of the "batch" service "submit" endpoint HTTP
// request body.
type SubmitRequestBody struct {
	Path            *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
	CompletedDir    *string `form:"completed_dir,omitempty" json:"completed_dir,omitempty" xml:"completed_dir,omitempty"`
	RetentionPeriod *string `form:"retention_period,omitempty" json:"retention_period,omitempty" xml:"retention_period,omitempty"`
}

// SubmitResponseBody is the type of the "batch" service "submit" endpoint HTTP
// response body.
type SubmitResponseBody struct {
	WorkflowID string `form:"workflow_id" json:"workflow_id" xml:"workflow_id"`
	RunID      string `form:"run_id" json:"run_id" xml:"run_id"`
}

// StatusResponseBody is the type of the "batch" service "status" endpoint HTTP
// response body.
type StatusResponseBody struct {
	Running    bool    `form:"running" json:"running" xml:"running"`
	Status     *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	RunID      *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
}

// HintsResponseBody is the type of the "batch" service "hints" endpoint HTTP
// response body.
type HintsResponseBody struct {
	// A list of known values of completedDir used by existing watchers.
	CompletedDirs []string `form:"completed_dirs,omitempty" json:"completed_dirs,omitempty" xml:"completed_dirs,omitempty"`
}

// SubmitNotAvailableResponseBody is the type of the "batch" service "submit"
// endpoint HTTP response body for the "not_available" error.
type SubmitNotAvailableResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// SubmitNotValidResponseBody is the type of the "batch" service "submit"
// endpoint HTTP response body for the "not_valid" error.
type SubmitNotValidResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// NewSubmitResponseBody builds the HTTP response body from the result of the
// "submit" endpoint of the "batch" service.
func NewSubmitResponseBody(res *batch.BatchResult) *SubmitResponseBody {
	body := &SubmitResponseBody{
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewStatusResponseBody builds the HTTP response body from the result of the
// "status" endpoint of the "batch" service.
func NewStatusResponseBody(res *batch.BatchStatusResult) *StatusResponseBody {
	body := &StatusResponseBody{
		Running:    res.Running,
		Status:     res.Status,
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewHintsResponseBody builds the HTTP response body from the result of the
// "hints" endpoint of the "batch" service.
func NewHintsResponseBody(res *batch.BatchHintsResult) *HintsResponseBody {
	body := &HintsResponseBody{}
	if res.CompletedDirs != nil {
		body.CompletedDirs = make([]string, len(res.CompletedDirs))
		for i, val := range res.CompletedDirs {
			body.CompletedDirs[i] = val
		}
	}
	return body
}

// NewSubmitNotAvailableResponseBody builds the HTTP response body from the
// result of the "submit" endpoint of the "batch" service.
func NewSubmitNotAvailableResponseBody(res *goa.ServiceError) *SubmitNotAvailableResponseBody {
	body := &SubmitNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSubmitNotValidResponseBody builds the HTTP response body from the result
// of the "submit" endpoint of the "batch" service.
func NewSubmitNotValidResponseBody(res *goa.ServiceError) *SubmitNotValidResponseBody {
	body := &SubmitNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewSubmitPayload builds a batch service submit endpoint payload.
func NewSubmitPayload(body *SubmitRequestBody) *batch.SubmitPayload {
	v := &batch.SubmitPayload{
		Path:            *body.Path,
		CompletedDir:    body.CompletedDir,
		RetentionPeriod: body.RetentionPeriod,
	}

	return v
}

// ValidateSubmitRequestBody runs the validations defined on SubmitRequestBody
func ValidateSubmitRequestBody(body *SubmitRequestBody) (err error) {
	if body.Path == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("path", "body"))
	}
	return
}
