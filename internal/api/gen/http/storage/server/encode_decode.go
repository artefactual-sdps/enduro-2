// Code generated by goa v3.7.12, DO NOT EDIT.
//
// storage HTTP server encoders and decoders
//
// Command:
// $ goa-v3.7.10 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package server

import (
	"context"
	"errors"
	"io"
	"net/http"

	storage "github.com/artefactual-sdps/enduro/internal/api/gen/storage"
	storageviews "github.com/artefactual-sdps/enduro/internal/api/gen/storage/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeSubmitResponse returns an encoder for responses returned by the
// storage submit endpoint.
func EncodeSubmitResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*storage.SubmitResult)
		enc := encoder(ctx, w)
		body := NewSubmitResponseBody(res)
		w.WriteHeader(http.StatusAccepted)
		return enc.Encode(body)
	}
}

// DecodeSubmitRequest returns a decoder for requests sent to the storage
// submit endpoint.
func DecodeSubmitRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body SubmitRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateSubmitRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewSubmitPayload(&body, aipID)

		return payload, nil
	}
}

// EncodeSubmitError returns an encoder for errors returned by the submit
// storage endpoint.
func EncodeSubmitError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSubmitNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewSubmitNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the
// storage update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the storage
// update endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewUpdatePayload(aipID)

		return payload, nil
	}
}

// EncodeUpdateError returns an encoder for errors returned by the update
// storage endpoint.
func EncodeUpdateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDownloadResponse returns an encoder for responses returned by the
// storage download endpoint.
func EncodeDownloadResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]byte)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeDownloadRequest returns a decoder for requests sent to the storage
// download endpoint.
func DecodeDownloadRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewDownloadPayload(aipID)

		return payload, nil
	}
}

// EncodeDownloadError returns an encoder for errors returned by the download
// storage endpoint.
func EncodeDownloadError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *storage.StoragePackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDownloadNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListResponse returns an encoder for responses returned by the storage
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(storageviews.StoredLocationCollection)
		enc := encoder(ctx, w)
		body := NewStoredLocationResponseCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeMoveResponse returns an encoder for responses returned by the storage
// move endpoint.
func EncodeMoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeMoveRequest returns a decoder for requests sent to the storage move
// endpoint.
func DecodeMoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body MoveRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateMoveRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewMovePayload(&body, aipID)

		return payload, nil
	}
}

// EncodeMoveError returns an encoder for errors returned by the move storage
// endpoint.
func EncodeMoveError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *storage.StoragePackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeMoveStatusResponse returns an encoder for responses returned by the
// storage move_status endpoint.
func EncodeMoveStatusResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*storage.MoveStatusResult)
		enc := encoder(ctx, w)
		body := NewMoveStatusResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMoveStatusRequest returns a decoder for requests sent to the storage
// move_status endpoint.
func DecodeMoveStatusRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewMoveStatusPayload(aipID)

		return payload, nil
	}
}

// EncodeMoveStatusError returns an encoder for errors returned by the
// move_status storage endpoint.
func EncodeMoveStatusError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "failed_dependency":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveStatusFailedDependencyResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusFailedDependency)
			return enc.Encode(body)
		case "not_found":
			var res *storage.StoragePackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewMoveStatusNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeRejectResponse returns an encoder for responses returned by the
// storage reject endpoint.
func EncodeRejectResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusAccepted)
		return nil
	}
}

// DecodeRejectRequest returns a decoder for requests sent to the storage
// reject endpoint.
func DecodeRejectRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewRejectPayload(aipID)

		return payload, nil
	}
}

// EncodeRejectError returns an encoder for errors returned by the reject
// storage endpoint.
func EncodeRejectError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_available":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRejectNotAvailableResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusConflict)
			return enc.Encode(body)
		case "not_valid":
			var res *goa.ServiceError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRejectNotValidResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "not_found":
			var res *storage.StoragePackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewRejectNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeShowResponse returns an encoder for responses returned by the storage
// show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*storageviews.StoredStoragePackage)
		enc := encoder(ctx, w)
		body := NewShowResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the storage show
// endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			aipID string

			params = mux.Vars(r)
		)
		aipID = params["aip_id"]
		payload := NewShowPayload(aipID)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show storage
// endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			var res *storage.StoragePackageNotfound
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalStorageviewsStoredLocationViewToStoredLocationResponse builds a value
// of type *StoredLocationResponse from a value of type
// *storageviews.StoredLocationView.
func marshalStorageviewsStoredLocationViewToStoredLocationResponse(v *storageviews.StoredLocationView) *StoredLocationResponse {
	res := &StoredLocationResponse{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}
