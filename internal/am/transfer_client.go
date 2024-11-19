package am

import (
	"context"
	"io"
)

type Deleter interface {
	// Delete removes dest from the transfer location.
	Delete(ctx context.Context, dest string) error
}

type Uploader interface {
	// Upload.
	Upload(ctx context.Context, src io.Reader, dest string) error
}

type TransferClient interface {
	Uploader
	Deleter
}
