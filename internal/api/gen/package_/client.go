// Code generated by goa v3.7.10, DO NOT EDIT.
//
// package client
//
// Command:
// $ goa-v3.7.10 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package package_

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "package" service client.
type Client struct {
	MonitorEndpoint             goa.Endpoint
	ListEndpoint                goa.Endpoint
	ShowEndpoint                goa.Endpoint
	DeleteEndpoint              goa.Endpoint
	CancelEndpoint              goa.Endpoint
	RetryEndpoint               goa.Endpoint
	WorkflowEndpoint            goa.Endpoint
	BulkEndpoint                goa.Endpoint
	BulkStatusEndpoint          goa.Endpoint
	PreservationActionsEndpoint goa.Endpoint
	ConfirmEndpoint             goa.Endpoint
	RejectEndpoint              goa.Endpoint
	MoveEndpoint                goa.Endpoint
	MoveStatusEndpoint          goa.Endpoint
}

// NewClient initializes a "package" service client given the endpoints.
func NewClient(monitor, list, show, delete_, cancel, retry, workflow, bulk, bulkStatus, preservationActions, confirm, reject, move, moveStatus goa.Endpoint) *Client {
	return &Client{
		MonitorEndpoint:             monitor,
		ListEndpoint:                list,
		ShowEndpoint:                show,
		DeleteEndpoint:              delete_,
		CancelEndpoint:              cancel,
		RetryEndpoint:               retry,
		WorkflowEndpoint:            workflow,
		BulkEndpoint:                bulk,
		BulkStatusEndpoint:          bulkStatus,
		PreservationActionsEndpoint: preservationActions,
		ConfirmEndpoint:             confirm,
		RejectEndpoint:              reject,
		MoveEndpoint:                move,
		MoveStatusEndpoint:          moveStatus,
	}
}

// Monitor calls the "monitor" endpoint of the "package" service.
func (c *Client) Monitor(ctx context.Context) (res MonitorClientStream, err error) {
	var ires interface{}
	ires, err = c.MonitorEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(MonitorClientStream), nil
}

// List calls the "list" endpoint of the "package" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res *ListResult, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}

// Show calls the "show" endpoint of the "package" service.
// Show may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *EnduroStoredPackage, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EnduroStoredPackage), nil
}

// Delete calls the "delete" endpoint of the "package" service.
// Delete may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}

// Cancel calls the "cancel" endpoint of the "package" service.
// Cancel may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- "not_running" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Cancel(ctx context.Context, p *CancelPayload) (err error) {
	_, err = c.CancelEndpoint(ctx, p)
	return
}

// Retry calls the "retry" endpoint of the "package" service.
// Retry may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- "not_running" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Retry(ctx context.Context, p *RetryPayload) (err error) {
	_, err = c.RetryEndpoint(ctx, p)
	return
}

// Workflow calls the "workflow" endpoint of the "package" service.
// Workflow may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- error: internal error
func (c *Client) Workflow(ctx context.Context, p *WorkflowPayload) (res *EnduroPackageWorkflowStatus, err error) {
	var ires interface{}
	ires, err = c.WorkflowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EnduroPackageWorkflowStatus), nil
}

// Bulk calls the "bulk" endpoint of the "package" service.
// Bulk may return the following errors:
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Bulk(ctx context.Context, p *BulkPayload) (res *BulkResult, err error) {
	var ires interface{}
	ires, err = c.BulkEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BulkResult), nil
}

// BulkStatus calls the "bulk_status" endpoint of the "package" service.
func (c *Client) BulkStatus(ctx context.Context) (res *BulkStatusResult, err error) {
	var ires interface{}
	ires, err = c.BulkStatusEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*BulkStatusResult), nil
}

// PreservationActions calls the "preservation-actions" endpoint of the
// "package" service.
// PreservationActions may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- error: internal error
func (c *Client) PreservationActions(ctx context.Context, p *PreservationActionsPayload) (res *EnduroPackagePreservationActions, err error) {
	var ires interface{}
	ires, err = c.PreservationActionsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EnduroPackagePreservationActions), nil
}

// Confirm calls the "confirm" endpoint of the "package" service.
// Confirm may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Confirm(ctx context.Context, p *ConfirmPayload) (err error) {
	_, err = c.ConfirmEndpoint(ctx, p)
	return
}

// Reject calls the "reject" endpoint of the "package" service.
// Reject may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Reject(ctx context.Context, p *RejectPayload) (err error) {
	_, err = c.RejectEndpoint(ctx, p)
	return
}

// Move calls the "move" endpoint of the "package" service.
// Move may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- "not_available" (type *goa.ServiceError)
//	- "not_valid" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Move(ctx context.Context, p *MovePayload) (err error) {
	_, err = c.MoveEndpoint(ctx, p)
	return
}

// MoveStatus calls the "move_status" endpoint of the "package" service.
// MoveStatus may return the following errors:
//	- "not_found" (type *PackageNotfound): Package not found
//	- "failed_dependency" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) MoveStatus(ctx context.Context, p *MoveStatusPayload) (res *MoveStatusResult, err error) {
	var ires interface{}
	ires, err = c.MoveStatusEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MoveStatusResult), nil
}
