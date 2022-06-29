// Code generated by goa v3.7.6, DO NOT EDIT.
//
// storage HTTP client CLI support package
//
// Command:
// $ goa-v3.7.6 gen github.com/artefactual-labs/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"

	storage "github.com/artefactual-labs/enduro/internal/api/gen/storage"
)

// BuildSubmitPayload builds the payload for the storage submit endpoint from
// CLI flags.
func BuildSubmitPayload(storageSubmitBody string, storageSubmitAipID string) (*storage.SubmitPayload, error) {
	var err error
	var body SubmitRequestBody
	{
		err = json.Unmarshal([]byte(storageSubmitBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"name\": \"Eius fugit natus iste ad.\"\n   }'")
		}
	}
	var aipID string
	{
		aipID = storageSubmitAipID
	}
	v := &storage.SubmitPayload{
		Name: body.Name,
	}
	v.AipID = aipID

	return v, nil
}

// BuildUpdatePayload builds the payload for the storage update endpoint from
// CLI flags.
func BuildUpdatePayload(storageUpdateAipID string) (*storage.UpdatePayload, error) {
	var aipID string
	{
		aipID = storageUpdateAipID
	}
	v := &storage.UpdatePayload{}
	v.AipID = aipID

	return v, nil
}

// BuildDownloadPayload builds the payload for the storage download endpoint
// from CLI flags.
func BuildDownloadPayload(storageDownloadAipID string) (*storage.DownloadPayload, error) {
	var aipID string
	{
		aipID = storageDownloadAipID
	}
	v := &storage.DownloadPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildMovePayload builds the payload for the storage move endpoint from CLI
// flags.
func BuildMovePayload(storageMoveBody string, storageMoveAipID string) (*storage.MovePayload, error) {
	var err error
	var body MoveRequestBody
	{
		err = json.Unmarshal([]byte(storageMoveBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"location\": \"Et voluptatem quia pariatur voluptates ut blanditiis.\"\n   }'")
		}
	}
	var aipID string
	{
		aipID = storageMoveAipID
	}
	v := &storage.MovePayload{
		Location: body.Location,
	}
	v.AipID = aipID

	return v, nil
}

// BuildMoveStatusPayload builds the payload for the storage move_status
// endpoint from CLI flags.
func BuildMoveStatusPayload(storageMoveStatusAipID string) (*storage.MoveStatusPayload, error) {
	var aipID string
	{
		aipID = storageMoveStatusAipID
	}
	v := &storage.MoveStatusPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildRejectPayload builds the payload for the storage reject endpoint from
// CLI flags.
func BuildRejectPayload(storageRejectAipID string) (*storage.RejectPayload, error) {
	var aipID string
	{
		aipID = storageRejectAipID
	}
	v := &storage.RejectPayload{}
	v.AipID = aipID

	return v, nil
}

// BuildShowPayload builds the payload for the storage show endpoint from CLI
// flags.
func BuildShowPayload(storageShowAipID string) (*storage.ShowPayload, error) {
	var aipID string
	{
		aipID = storageShowAipID
	}
	v := &storage.ShowPayload{}
	v.AipID = aipID

	return v, nil
}
