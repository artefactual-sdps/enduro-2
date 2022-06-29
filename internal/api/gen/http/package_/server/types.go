// Code generated by goa v3.7.6, DO NOT EDIT.
//
// package HTTP server types
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

import (
	package_ "github.com/artefactual-labs/enduro/internal/api/gen/package_"
	package_views "github.com/artefactual-labs/enduro/internal/api/gen/package_/views"
	goa "goa.design/goa/v3/pkg"
)

// BulkRequestBody is the type of the "package" service "bulk" endpoint HTTP
// request body.
type BulkRequestBody struct {
	Operation *string `form:"operation,omitempty" json:"operation,omitempty" xml:"operation,omitempty"`
	Status    *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	Size      *uint   `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
}

// ConfirmRequestBody is the type of the "package" service "confirm" endpoint
// HTTP request body.
type ConfirmRequestBody struct {
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
}

// MonitorResponseBody is the type of the "package" service "monitor" endpoint
// HTTP response body.
type MonitorResponseBody struct {
	// Identifier of package
	ID uint `form:"id" json:"id" xml:"id"`
	// Type of the event
	Type string `form:"type" json:"type" xml:"type"`
	// Package
	Item *EnduroStoredPackageResponseBody `form:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
}

// ListResponseBody is the type of the "package" service "list" endpoint HTTP
// response body.
type ListResponseBody struct {
	Items      EnduroStoredPackageCollectionResponseBody `form:"items" json:"items" xml:"items"`
	NextCursor *string                                   `form:"next_cursor,omitempty" json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
}

// ShowResponseBody is the type of the "package" service "show" endpoint HTTP
// response body.
type ShowResponseBody struct {
	// Identifier of package
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the package
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Location of the package
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Status of the package
	Status string `form:"status" json:"status" xml:"status"`
	// Identifier of processing workflow
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	// Identifier of latest processing workflow run
	RunID *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
	// Identifier of Archivematica AIP
	AipID *string `form:"aip_id,omitempty" json:"aip_id,omitempty" xml:"aip_id,omitempty"`
	// Creation datetime
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// Start datetime
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// Completion datetime
	CompletedAt *string `form:"completed_at,omitempty" json:"completed_at,omitempty" xml:"completed_at,omitempty"`
}

// WorkflowResponseBody is the type of the "package" service "workflow"
// endpoint HTTP response body.
type WorkflowResponseBody struct {
	Status  *string                                            `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	History EnduroPackageWorkflowHistoryResponseBodyCollection `form:"history,omitempty" json:"history,omitempty" xml:"history,omitempty"`
}

// BulkResponseBody is the type of the "package" service "bulk" endpoint HTTP
// response body.
type BulkResponseBody struct {
	WorkflowID string `form:"workflow_id" json:"workflow_id" xml:"workflow_id"`
	RunID      string `form:"run_id" json:"run_id" xml:"run_id"`
}

// BulkStatusResponseBody is the type of the "package" service "bulk_status"
// endpoint HTTP response body.
type BulkStatusResponseBody struct {
	Running    bool    `form:"running" json:"running" xml:"running"`
	StartedAt  *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	ClosedAt   *string `form:"closed_at,omitempty" json:"closed_at,omitempty" xml:"closed_at,omitempty"`
	Status     *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	RunID      *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
}

// PreservationActionsResponseBody is the type of the "package" service
// "preservation-actions" endpoint HTTP response body.
type PreservationActionsResponseBody struct {
	Actions EnduroPackagePreservationActionsActionResponseBodyCollection `form:"actions,omitempty" json:"actions,omitempty" xml:"actions,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "package" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing package
	ID uint `form:"id" json:"id" xml:"id"`
}

// DeleteNotFoundResponseBody is the type of the "package" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing package
	ID uint `form:"id" json:"id" xml:"id"`
}

// CancelNotRunningResponseBody is the type of the "package" service "cancel"
// endpoint HTTP response body for the "not_running" error.
type CancelNotRunningResponseBody struct {
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

// CancelNotFoundResponseBody is the type of the "package" service "cancel"
// endpoint HTTP response body for the "not_found" error.
type CancelNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing package
	ID uint `form:"id" json:"id" xml:"id"`
}

// RetryNotRunningResponseBody is the type of the "package" service "retry"
// endpoint HTTP response body for the "not_running" error.
type RetryNotRunningResponseBody struct {
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

// RetryNotFoundResponseBody is the type of the "package" service "retry"
// endpoint HTTP response body for the "not_found" error.
type RetryNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing package
	ID uint `form:"id" json:"id" xml:"id"`
}

// WorkflowNotFoundResponseBody is the type of the "package" service "workflow"
// endpoint HTTP response body for the "not_found" error.
type WorkflowNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing package
	ID uint `form:"id" json:"id" xml:"id"`
}

// BulkNotAvailableResponseBody is the type of the "package" service "bulk"
// endpoint HTTP response body for the "not_available" error.
type BulkNotAvailableResponseBody struct {
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

// BulkNotValidResponseBody is the type of the "package" service "bulk"
// endpoint HTTP response body for the "not_valid" error.
type BulkNotValidResponseBody struct {
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

// PreservationActionsNotFoundResponseBody is the type of the "package" service
// "preservation-actions" endpoint HTTP response body for the "not_found" error.
type PreservationActionsNotFoundResponseBody struct {
	// Message of error
	Message string `form:"message" json:"message" xml:"message"`
	// Identifier of missing package
	ID uint `form:"id" json:"id" xml:"id"`
}

// ConfirmNotAvailableResponseBody is the type of the "package" service
// "confirm" endpoint HTTP response body for the "not_available" error.
type ConfirmNotAvailableResponseBody struct {
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

// ConfirmNotValidResponseBody is the type of the "package" service "confirm"
// endpoint HTTP response body for the "not_valid" error.
type ConfirmNotValidResponseBody struct {
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

// RejectNotAvailableResponseBody is the type of the "package" service "reject"
// endpoint HTTP response body for the "not_available" error.
type RejectNotAvailableResponseBody struct {
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

// RejectNotValidResponseBody is the type of the "package" service "reject"
// endpoint HTTP response body for the "not_valid" error.
type RejectNotValidResponseBody struct {
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

// EnduroStoredPackageResponseBody is used to define fields on response body
// types.
type EnduroStoredPackageResponseBody struct {
	// Identifier of package
	ID uint `form:"id" json:"id" xml:"id"`
	// Name of the package
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Location of the package
	Location *string `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	// Status of the package
	Status string `form:"status" json:"status" xml:"status"`
	// Identifier of processing workflow
	WorkflowID *string `form:"workflow_id,omitempty" json:"workflow_id,omitempty" xml:"workflow_id,omitempty"`
	// Identifier of latest processing workflow run
	RunID *string `form:"run_id,omitempty" json:"run_id,omitempty" xml:"run_id,omitempty"`
	// Identifier of Archivematica AIP
	AipID *string `form:"aip_id,omitempty" json:"aip_id,omitempty" xml:"aip_id,omitempty"`
	// Creation datetime
	CreatedAt string `form:"created_at" json:"created_at" xml:"created_at"`
	// Start datetime
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// Completion datetime
	CompletedAt *string `form:"completed_at,omitempty" json:"completed_at,omitempty" xml:"completed_at,omitempty"`
}

// EnduroStoredPackageCollectionResponseBody is used to define fields on
// response body types.
type EnduroStoredPackageCollectionResponseBody []*EnduroStoredPackageResponseBody

// EnduroPackageWorkflowHistoryResponseBodyCollection is used to define fields
// on response body types.
type EnduroPackageWorkflowHistoryResponseBodyCollection []*EnduroPackageWorkflowHistoryResponseBody

// EnduroPackageWorkflowHistoryResponseBody is used to define fields on
// response body types.
type EnduroPackageWorkflowHistoryResponseBody struct {
	// Identifier of package
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Type of the event
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Contents of the event
	Details interface{} `form:"details,omitempty" json:"details,omitempty" xml:"details,omitempty"`
}

// EnduroPackagePreservationActionsActionResponseBodyCollection is used to
// define fields on response body types.
type EnduroPackagePreservationActionsActionResponseBodyCollection []*EnduroPackagePreservationActionsActionResponseBody

// EnduroPackagePreservationActionsActionResponseBody is used to define fields
// on response body types.
type EnduroPackagePreservationActionsActionResponseBody struct {
	ID          uint    `form:"id" json:"id" xml:"id"`
	ActionID    string  `form:"action_id" json:"action_id" xml:"action_id"`
	Name        string  `form:"name" json:"name" xml:"name"`
	Status      string  `form:"status" json:"status" xml:"status"`
	StartedAt   string  `form:"started_at" json:"started_at" xml:"started_at"`
	CompletedAt *string `form:"completed_at,omitempty" json:"completed_at,omitempty" xml:"completed_at,omitempty"`
}

// NewMonitorResponseBody builds the HTTP response body from the result of the
// "monitor" endpoint of the "package" service.
func NewMonitorResponseBody(res *package_views.EnduroMonitorUpdateView) *MonitorResponseBody {
	body := &MonitorResponseBody{
		ID:   *res.ID,
		Type: *res.Type,
	}
	if res.Item != nil {
		body.Item = marshalPackageViewsEnduroStoredPackageViewToEnduroStoredPackageResponseBody(res.Item)
	}
	return body
}

// NewListResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "package" service.
func NewListResponseBody(res *package_.ListResult) *ListResponseBody {
	body := &ListResponseBody{
		NextCursor: res.NextCursor,
	}
	if res.Items != nil {
		body.Items = make([]*EnduroStoredPackageResponseBody, len(res.Items))
		for i, val := range res.Items {
			body.Items[i] = marshalPackageEnduroStoredPackageToEnduroStoredPackageResponseBody(val)
		}
	}
	return body
}

// NewShowResponseBody builds the HTTP response body from the result of the
// "show" endpoint of the "package" service.
func NewShowResponseBody(res *package_views.EnduroStoredPackageView) *ShowResponseBody {
	body := &ShowResponseBody{
		ID:          *res.ID,
		Name:        res.Name,
		Location:    res.Location,
		Status:      *res.Status,
		WorkflowID:  res.WorkflowID,
		RunID:       res.RunID,
		AipID:       res.AipID,
		CreatedAt:   *res.CreatedAt,
		StartedAt:   res.StartedAt,
		CompletedAt: res.CompletedAt,
	}
	return body
}

// NewWorkflowResponseBody builds the HTTP response body from the result of the
// "workflow" endpoint of the "package" service.
func NewWorkflowResponseBody(res *package_views.EnduroPackageWorkflowStatusView) *WorkflowResponseBody {
	body := &WorkflowResponseBody{
		Status: res.Status,
	}
	if res.History != nil {
		body.History = make([]*EnduroPackageWorkflowHistoryResponseBody, len(res.History))
		for i, val := range res.History {
			body.History[i] = marshalPackageViewsEnduroPackageWorkflowHistoryViewToEnduroPackageWorkflowHistoryResponseBody(val)
		}
	}
	return body
}

// NewBulkResponseBody builds the HTTP response body from the result of the
// "bulk" endpoint of the "package" service.
func NewBulkResponseBody(res *package_.BulkResult) *BulkResponseBody {
	body := &BulkResponseBody{
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewBulkStatusResponseBody builds the HTTP response body from the result of
// the "bulk_status" endpoint of the "package" service.
func NewBulkStatusResponseBody(res *package_.BulkStatusResult) *BulkStatusResponseBody {
	body := &BulkStatusResponseBody{
		Running:    res.Running,
		StartedAt:  res.StartedAt,
		ClosedAt:   res.ClosedAt,
		Status:     res.Status,
		WorkflowID: res.WorkflowID,
		RunID:      res.RunID,
	}
	return body
}

// NewPreservationActionsResponseBody builds the HTTP response body from the
// result of the "preservation-actions" endpoint of the "package" service.
func NewPreservationActionsResponseBody(res *package_views.EnduroPackagePreservationActionsView) *PreservationActionsResponseBody {
	body := &PreservationActionsResponseBody{}
	if res.Actions != nil {
		body.Actions = make([]*EnduroPackagePreservationActionsActionResponseBody, len(res.Actions))
		for i, val := range res.Actions {
			body.Actions[i] = marshalPackageViewsEnduroPackagePreservationActionsActionViewToEnduroPackagePreservationActionsActionResponseBody(val)
		}
	}
	return body
}

// NewShowNotFoundResponseBody builds the HTTP response body from the result of
// the "show" endpoint of the "package" service.
func NewShowNotFoundResponseBody(res *package_.PackageNotfound) *ShowNotFoundResponseBody {
	body := &ShowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "package" service.
func NewDeleteNotFoundResponseBody(res *package_.PackageNotfound) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewCancelNotRunningResponseBody builds the HTTP response body from the
// result of the "cancel" endpoint of the "package" service.
func NewCancelNotRunningResponseBody(res *goa.ServiceError) *CancelNotRunningResponseBody {
	body := &CancelNotRunningResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCancelNotFoundResponseBody builds the HTTP response body from the result
// of the "cancel" endpoint of the "package" service.
func NewCancelNotFoundResponseBody(res *package_.PackageNotfound) *CancelNotFoundResponseBody {
	body := &CancelNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewRetryNotRunningResponseBody builds the HTTP response body from the result
// of the "retry" endpoint of the "package" service.
func NewRetryNotRunningResponseBody(res *goa.ServiceError) *RetryNotRunningResponseBody {
	body := &RetryNotRunningResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRetryNotFoundResponseBody builds the HTTP response body from the result
// of the "retry" endpoint of the "package" service.
func NewRetryNotFoundResponseBody(res *package_.PackageNotfound) *RetryNotFoundResponseBody {
	body := &RetryNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewWorkflowNotFoundResponseBody builds the HTTP response body from the
// result of the "workflow" endpoint of the "package" service.
func NewWorkflowNotFoundResponseBody(res *package_.PackageNotfound) *WorkflowNotFoundResponseBody {
	body := &WorkflowNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewBulkNotAvailableResponseBody builds the HTTP response body from the
// result of the "bulk" endpoint of the "package" service.
func NewBulkNotAvailableResponseBody(res *goa.ServiceError) *BulkNotAvailableResponseBody {
	body := &BulkNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewBulkNotValidResponseBody builds the HTTP response body from the result of
// the "bulk" endpoint of the "package" service.
func NewBulkNotValidResponseBody(res *goa.ServiceError) *BulkNotValidResponseBody {
	body := &BulkNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewPreservationActionsNotFoundResponseBody builds the HTTP response body
// from the result of the "preservation-actions" endpoint of the "package"
// service.
func NewPreservationActionsNotFoundResponseBody(res *package_.PackageNotfound) *PreservationActionsNotFoundResponseBody {
	body := &PreservationActionsNotFoundResponseBody{
		Message: res.Message,
		ID:      res.ID,
	}
	return body
}

// NewConfirmNotAvailableResponseBody builds the HTTP response body from the
// result of the "confirm" endpoint of the "package" service.
func NewConfirmNotAvailableResponseBody(res *goa.ServiceError) *ConfirmNotAvailableResponseBody {
	body := &ConfirmNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewConfirmNotValidResponseBody builds the HTTP response body from the result
// of the "confirm" endpoint of the "package" service.
func NewConfirmNotValidResponseBody(res *goa.ServiceError) *ConfirmNotValidResponseBody {
	body := &ConfirmNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRejectNotAvailableResponseBody builds the HTTP response body from the
// result of the "reject" endpoint of the "package" service.
func NewRejectNotAvailableResponseBody(res *goa.ServiceError) *RejectNotAvailableResponseBody {
	body := &RejectNotAvailableResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewRejectNotValidResponseBody builds the HTTP response body from the result
// of the "reject" endpoint of the "package" service.
func NewRejectNotValidResponseBody(res *goa.ServiceError) *RejectNotValidResponseBody {
	body := &RejectNotValidResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewListPayload builds a package service list endpoint payload.
func NewListPayload(name *string, aipID *string, earliestCreatedTime *string, latestCreatedTime *string, location *string, status *string, cursor *string) *package_.ListPayload {
	v := &package_.ListPayload{}
	v.Name = name
	v.AipID = aipID
	v.EarliestCreatedTime = earliestCreatedTime
	v.LatestCreatedTime = latestCreatedTime
	v.Location = location
	v.Status = status
	v.Cursor = cursor

	return v
}

// NewShowPayload builds a package service show endpoint payload.
func NewShowPayload(id uint) *package_.ShowPayload {
	v := &package_.ShowPayload{}
	v.ID = id

	return v
}

// NewDeletePayload builds a package service delete endpoint payload.
func NewDeletePayload(id uint) *package_.DeletePayload {
	v := &package_.DeletePayload{}
	v.ID = id

	return v
}

// NewCancelPayload builds a package service cancel endpoint payload.
func NewCancelPayload(id uint) *package_.CancelPayload {
	v := &package_.CancelPayload{}
	v.ID = id

	return v
}

// NewRetryPayload builds a package service retry endpoint payload.
func NewRetryPayload(id uint) *package_.RetryPayload {
	v := &package_.RetryPayload{}
	v.ID = id

	return v
}

// NewWorkflowPayload builds a package service workflow endpoint payload.
func NewWorkflowPayload(id uint) *package_.WorkflowPayload {
	v := &package_.WorkflowPayload{}
	v.ID = id

	return v
}

// NewBulkPayload builds a package service bulk endpoint payload.
func NewBulkPayload(body *BulkRequestBody) *package_.BulkPayload {
	v := &package_.BulkPayload{
		Operation: *body.Operation,
		Status:    *body.Status,
	}
	if body.Size != nil {
		v.Size = *body.Size
	}
	if body.Size == nil {
		v.Size = 100
	}

	return v
}

// NewPreservationActionsPayload builds a package service preservation-actions
// endpoint payload.
func NewPreservationActionsPayload(id uint) *package_.PreservationActionsPayload {
	v := &package_.PreservationActionsPayload{}
	v.ID = id

	return v
}

// NewConfirmPayload builds a package service confirm endpoint payload.
func NewConfirmPayload(body *ConfirmRequestBody, id uint) *package_.ConfirmPayload {
	v := &package_.ConfirmPayload{
		Location: *body.Location,
	}
	v.ID = id

	return v
}

// NewRejectPayload builds a package service reject endpoint payload.
func NewRejectPayload(id uint) *package_.RejectPayload {
	v := &package_.RejectPayload{}
	v.ID = id

	return v
}

// ValidateBulkRequestBody runs the validations defined on BulkRequestBody
func ValidateBulkRequestBody(body *BulkRequestBody) (err error) {
	if body.Operation == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("operation", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Operation != nil {
		if !(*body.Operation == "retry" || *body.Operation == "cancel" || *body.Operation == "abandon") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operation", *body.Operation, []interface{}{"retry", "cancel", "abandon"}))
		}
	}
	if body.Status != nil {
		if !(*body.Status == "new" || *body.Status == "in progress" || *body.Status == "done" || *body.Status == "error" || *body.Status == "unknown" || *body.Status == "queued" || *body.Status == "pending" || *body.Status == "abandoned") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"new", "in progress", "done", "error", "unknown", "queued", "pending", "abandoned"}))
		}
	}
	return
}

// ValidateConfirmRequestBody runs the validations defined on ConfirmRequestBody
func ValidateConfirmRequestBody(body *ConfirmRequestBody) (err error) {
	if body.Location == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("location", "body"))
	}
	return
}
