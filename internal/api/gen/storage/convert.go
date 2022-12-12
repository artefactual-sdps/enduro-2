// Code generated by goa v3.9.1, DO NOT EDIT.
//
// storage service type conversion functions
//
// Command:
// $ goa-v3.10.2 gen github.com/artefactual-sdps/enduro/internal/api/design -o
// internal/api

package storage

import (
	types "github.com/artefactual-sdps/enduro/internal/storage/types"
)

// ConvertToS3Config creates an instance of S3Config initialized from t.
func (t *S3Config) ConvertToS3Config() *types.S3Config {
	v := &types.S3Config{
		Bucket: t.Bucket,
		Region: t.Region,
	}
	if t.Endpoint != nil {
		v.Endpoint = *t.Endpoint
	}
	if t.PathStyle != nil {
		v.PathStyle = *t.PathStyle
	}
	if t.Profile != nil {
		v.Profile = *t.Profile
	}
	if t.Key != nil {
		v.Key = *t.Key
	}
	if t.Secret != nil {
		v.Secret = *t.Secret
	}
	if t.Token != nil {
		v.Token = *t.Token
	}
	return v
}

// ConvertToSFTPConfig creates an instance of SFTPConfig initialized from t.
func (t *SFTPConfig) ConvertToSFTPConfig() *types.SFTPConfig {
	v := &types.SFTPConfig{
		Address:   t.Address,
		Username:  t.Username,
		Password:  t.Password,
		Directory: t.Directory,
	}
	return v
}
