package am

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FSTransferClient implements a TransferClient for the local filesystem.
type FSTransferClient struct{}

var _ TransferClient = (*FSTransferClient)(nil)

func NewFSTransferClient() FSTransferClient {
	return FSTransferClient{}
}

func (c FSTransferClient) Upload(
	ctx context.Context,
	src io.Reader,
	dest string,
) error {
	f, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("upload: create dest: %v", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, src); err != nil {
		return fmt.Errorf("upload: %v", err)
	}

	return nil
}

func (c FSTransferClient) Delete(ctx context.Context, dest string) error {
	if err := os.Remove(filepath.Clean(dest)); err != nil {
		return err
	}

	return nil
}
