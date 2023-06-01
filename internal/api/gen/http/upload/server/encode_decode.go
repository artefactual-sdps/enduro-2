// Code generated by goa v3.11.3, DO NOT EDIT.
//
// upload HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"errors"
	"net/http"
	"strings"

	upload "github.com/artefactual-sdps/enduro/internal/api/gen/upload"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeUploadResponse returns an encoder for responses returned by the upload
// upload endpoint.
func EncodeUploadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeUploadRequest returns a decoder for requests sent to the upload upload
// endpoint.
func DecodeUploadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		var (
			contentType string
			oauthToken  *string
			err         error
		)
		contentTypeRaw := r.Header.Get("Content-Type")
		if contentTypeRaw != "" {
			contentType = contentTypeRaw
		} else {
			contentType = "multipart/form-data; boundary=goa"
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("contentType", contentType, "multipart/[^;]+; boundary=.+"))
		oauthTokenRaw := r.Header.Get("Authorization")
		if oauthTokenRaw != "" {
			oauthToken = &oauthTokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewUploadPayload(contentType, oauthToken)
		if payload.OauthToken != nil {
			if strings.Contains(*payload.OauthToken, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.OauthToken, " ", 2)[1]
				payload.OauthToken = &cred
			}
		}

		return payload, nil
	}
}

// EncodeUploadError returns an encoder for errors returned by the upload
// upload endpoint.
func EncodeUploadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "invalid_media_type":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadInvalidMediaTypeResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "invalid_multipart_request":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadInvalidMultipartRequestResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "internal_error":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body any
			if formatter != nil {
				body = formatter(ctx, res)
			} else {
				body = NewUploadInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		case "unauthorized":
			var res upload.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
