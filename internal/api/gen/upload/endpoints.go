// Code generated by goa v3.8.5, DO NOT EDIT.
//
// upload endpoints
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

// Endpoints wraps the "upload" service endpoints.
type Endpoints struct {
	Upload goa.Endpoint
}

// UploadRequestData holds both the payload and the HTTP request body reader of
// the "upload" method.
type UploadRequestData struct {
	// Payload is the method payload.
	Payload *UploadPayload
	// Body streams the HTTP request body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "upload" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Upload: NewUploadEndpoint(s),
	}
}

// Use applies the given middleware to all the "upload" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Upload = m(e.Upload)
}

// NewUploadEndpoint returns an endpoint function that calls the method
// "upload" of service "upload".
func NewUploadEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		ep := req.(*UploadRequestData)
		return nil, s.Upload(ctx, ep.Payload, ep.Body)
	}
}
