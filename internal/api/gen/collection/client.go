// Code generated by goa v3.5.5, DO NOT EDIT.
//
// collection client
//
// Command:
// $ goa-v3.5.5 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package collection

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "collection" service client.
type Client struct {
	MonitorEndpoint             goa.Endpoint
	ListEndpoint                goa.Endpoint
	ShowEndpoint                goa.Endpoint
	DeleteEndpoint              goa.Endpoint
	CancelEndpoint              goa.Endpoint
	RetryEndpoint               goa.Endpoint
	WorkflowEndpoint            goa.Endpoint
	DownloadEndpoint            goa.Endpoint
	BulkEndpoint                goa.Endpoint
	BulkStatusEndpoint          goa.Endpoint
	PreservationActionsEndpoint goa.Endpoint
}

// NewClient initializes a "collection" service client given the endpoints.
func NewClient(monitor, list, show, delete_, cancel, retry, workflow, download, bulk, bulkStatus, preservationActions goa.Endpoint) *Client {
	return &Client{
		MonitorEndpoint:             monitor,
		ListEndpoint:                list,
		ShowEndpoint:                show,
		DeleteEndpoint:              delete_,
		CancelEndpoint:              cancel,
		RetryEndpoint:               retry,
		WorkflowEndpoint:            workflow,
		DownloadEndpoint:            download,
		BulkEndpoint:                bulk,
		BulkStatusEndpoint:          bulkStatus,
		PreservationActionsEndpoint: preservationActions,
	}
}

// Monitor calls the "monitor" endpoint of the "collection" service.
func (c *Client) Monitor(ctx context.Context) (res MonitorClientStream, err error) {
	var ires interface{}
	ires, err = c.MonitorEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(MonitorClientStream), nil
}

// List calls the "list" endpoint of the "collection" service.
func (c *Client) List(ctx context.Context, p *ListPayload) (res *ListResult, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListResult), nil
}

// Show calls the "show" endpoint of the "collection" service.
// Show may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- error: internal error
func (c *Client) Show(ctx context.Context, p *ShowPayload) (res *EnduroStoredCollection, err error) {
	var ires interface{}
	ires, err = c.ShowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EnduroStoredCollection), nil
}

// Delete calls the "delete" endpoint of the "collection" service.
// Delete may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- error: internal error
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}

// Cancel calls the "cancel" endpoint of the "collection" service.
// Cancel may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- "not_running" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Cancel(ctx context.Context, p *CancelPayload) (err error) {
	_, err = c.CancelEndpoint(ctx, p)
	return
}

// Retry calls the "retry" endpoint of the "collection" service.
// Retry may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- "not_running" (type *goa.ServiceError)
//	- error: internal error
func (c *Client) Retry(ctx context.Context, p *RetryPayload) (err error) {
	_, err = c.RetryEndpoint(ctx, p)
	return
}

// Workflow calls the "workflow" endpoint of the "collection" service.
// Workflow may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- error: internal error
func (c *Client) Workflow(ctx context.Context, p *WorkflowPayload) (res *EnduroCollectionWorkflowStatus, err error) {
	var ires interface{}
	ires, err = c.WorkflowEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EnduroCollectionWorkflowStatus), nil
}

// Download calls the "download" endpoint of the "collection" service.
// Download may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- error: internal error
func (c *Client) Download(ctx context.Context, p *DownloadPayload) (res []byte, err error) {
	var ires interface{}
	ires, err = c.DownloadEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]byte), nil
}

// Bulk calls the "bulk" endpoint of the "collection" service.
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

// BulkStatus calls the "bulk_status" endpoint of the "collection" service.
func (c *Client) BulkStatus(ctx context.Context) (res *BulkStatusResult, err error) {
	var ires interface{}
	ires, err = c.BulkStatusEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*BulkStatusResult), nil
}

// PreservationActions calls the "preservation-actions" endpoint of the
// "collection" service.
// PreservationActions may return the following errors:
//	- "not_found" (type *CollectionNotfound): Collection not found
//	- error: internal error
func (c *Client) PreservationActions(ctx context.Context, p *PreservationActionsPayload) (res *EnduroCollectionPreservationActions, err error) {
	var ires interface{}
	ires, err = c.PreservationActionsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*EnduroCollectionPreservationActions), nil
}
