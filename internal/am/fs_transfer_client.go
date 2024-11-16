package am

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/artefactual-sdps/enduro/internal/async"
)

// FSTransferClient implements a TransferClient for the local filesystem.
type FSTransferClient struct {
	RootDir string
}

var _ TransferClient = (*FSTransferClient)(nil)

func NewFSTransferClient(root string) FSTransferClient {
	return FSTransferClient{RootDir: root}
}

func (c FSTransferClient) Upload(
	ctx context.Context,
	src io.Reader,
	dest string,
) (remotePath string, upload async.Upload, err error) {
	remotePath = filepath.Join(c.RootDir, dest)

	f, err := os.Create(remotePath)
	if err != nil {
		return remotePath, nil, fmt.Errorf("upload: create dest: %v", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, src); err != nil {
		return remotePath, nil, fmt.Errorf("upload: %v", err)
	}

	return remotePath, nil, nil
}

func (c FSTransferClient) Delete(ctx context.Context, dest string) error {
	if err := os.Remove(filepath.Join(c.RootDir, filepath.Clean(dest))); err != nil {
		return err
	}

	return nil
}
