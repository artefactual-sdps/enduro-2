// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the storage service.
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package server

// SubmitStoragePath returns the URL path to the storage service submit HTTP endpoint.
func SubmitStoragePath() string {
	return "/storage/submit"
}

// UpdateStoragePath returns the URL path to the storage service update HTTP endpoint.
func UpdateStoragePath() string {
	return "/storage/update"
}
