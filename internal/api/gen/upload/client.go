// Code generated by goa v3.8.5, DO NOT EDIT.
//
// upload client
//
// Command:
// $ goa-v3.8.5 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package upload

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "upload" service client.
type Client struct {
	UploadEndpoint goa.Endpoint
}

// NewClient initializes a "upload" service client given the endpoints.
func NewClient(upload goa.Endpoint) *Client {
	return &Client{
		UploadEndpoint: upload,
	}
}

// Upload calls the "upload" endpoint of the "upload" service.
// Upload may return the following errors:
//   - "invalid_media_type" (type *goa.ServiceError): Error returned when the Content-Type header does not define a multipart request.
//   - "invalid_multipart_request" (type *goa.ServiceError): Error returned when the request body is not a valid multipart content.
//   - "internal_error" (type *goa.ServiceError): Fault while processing upload.
//   - error: internal error
func (c *Client) Upload(ctx context.Context, p *UploadPayload, req io.ReadCloser) (err error) {
	_, err = c.UploadEndpoint(ctx, &UploadRequestData{Payload: p, Body: req})
	return
}
