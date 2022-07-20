// Code generated by goa v3.7.12, DO NOT EDIT.
//
// batch HTTP client CLI support package
//
// Command:
// $ goa-v3.7.10 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package client

import (
	"encoding/json"
	"fmt"

	batch "github.com/artefactual-sdps/enduro/internal/api/gen/batch"
)

// BuildSubmitPayload builds the payload for the batch submit endpoint from CLI
// flags.
func BuildSubmitPayload(batchSubmitBody string) (*batch.SubmitPayload, error) {
	var err error
	var body SubmitRequestBody
	{
		err = json.Unmarshal([]byte(batchSubmitBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"completed_dir\": \"Vitae odit sunt.\",\n      \"path\": \"Illum quasi.\",\n      \"retention_period\": \"Dolor et suscipit.\"\n   }'")
		}
	}
	v := &batch.SubmitPayload{
		Path:            body.Path,
		CompletedDir:    body.CompletedDir,
		RetentionPeriod: body.RetentionPeriod,
	}

	return v, nil
}
