package am

import (
	"context"
	"io"

	"github.com/artefactual-sdps/enduro/internal/async"
)

type TransferClient interface {
	// Delete removes dest from the transfer location.
	Delete(ctx context.Context, dest string) error
	// Upload asynchronously copies data from the src reader to the specified
	// dest within the transfer location.
	Upload(ctx context.Context, src io.Reader, dest string) (remotePath string, upload async.Upload, err error)
}
