// Code generated by goa v3.11.3, DO NOT EDIT.
//
// swagger HTTP server
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"net/http"

	swagger "github.com/artefactual-sdps/enduro/internal/api/gen/swagger"
	goahttp "goa.design/goa/v3/http"
	"goa.design/plugins/v3/cors"
)

// Server lists the swagger service endpoint HTTP handlers.
type Server struct {
	Mounts                        []*MountPoint
	CORS                          http.Handler
	InternalAPIGenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the swagger service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *swagger.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemInternalAPIGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemInternalAPIGenHTTPOpenapiJSON == nil {
		fileSystemInternalAPIGenHTTPOpenapiJSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"CORS", "OPTIONS", "/swagger/swagger.json"},
			{"internal/api/gen/http/openapi.json", "GET", "/swagger/swagger.json"},
		},
		CORS:                          NewCORSHandler(),
		InternalAPIGenHTTPOpenapiJSON: http.FileServer(fileSystemInternalAPIGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "swagger" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return swagger.MethodNames[:] }

// Mount configures the mux to serve the swagger endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCORSHandler(mux, h.CORS)
	MountInternalAPIGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/internal/api/gen/http/openapi.json", h.InternalAPIGenHTTPOpenapiJSON))
}

// Mount configures the mux to serve the swagger endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountInternalAPIGenHTTPOpenapiJSON configures the mux to serve GET request
// made to "/swagger/swagger.json".
func MountInternalAPIGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/swagger/swagger.json", HandleSwaggerOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service swagger.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleSwaggerOrigin(h)
	mux.Handle("OPTIONS", "/swagger/swagger.json", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandleSwaggerOrigin applies the CORS response headers corresponding to the
// origin for the service swagger.
func HandleSwaggerOrigin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
