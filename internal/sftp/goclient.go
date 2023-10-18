package sftp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/dolmen-go/contextio"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// GoClient implements the SFTP service using native Go SSH and SFTP packages.
type GoClient struct {
	cfg Config

	ssh  *ssh.Client
	sftp *sftp.Client
}

var _ Client = (*GoClient)(nil)

// NewGoClient returns a new GoSFTP client with the given configuration.
func NewGoClient(cfg Config) *GoClient {
	cfg.SetDefaults()

	return &GoClient{cfg: cfg}
}

// Upload writes the data from src to the remote file at dest and returns the
// number of bytes written.  A new SFTP connection is opened before writing, and
// closed when the upload is complete or cancelled.
//
// Upload is not thread safe.
func (c *GoClient) Upload(ctx context.Context, src io.Reader, dest string) (int64, error) {
	if err := c.dial(); err != nil {
		return 0, err
	}
	defer c.close()

	// Note: Some SFTP servers don't support O_RDWR mode.
	w, err := c.sftp.OpenFile(dest, (os.O_WRONLY | os.O_CREATE | os.O_TRUNC))
	if err != nil {
		return 0, fmt.Errorf("SFTP: open remote file %q: %v", dest, err)
	}
	defer w.Close()

	// Use contextio to stop the upload if a context cancellation signal is
	// received.
	bytes, err := io.Copy(contextio.NewWriter(ctx, w), contextio.NewReader(ctx, src))
	if err != nil {
		return 0, fmt.Errorf("SFTP: upload to %q: %v", dest, err)
	}

	return bytes, nil
}

// Dial connects to an SSH host then creates an SFTP client on the connection.
// When the clients are no longer needed, close() must be called to prevent
// leaks.
func (c *GoClient) dial() error {
	var err error

	c.ssh, err = sshConnect(c.cfg)
	if err != nil {
		return fmt.Errorf("SSH: %v", err)
	}

	c.sftp, err = sftp.NewClient(c.ssh)
	if err != nil {
		return fmt.Errorf("start SFTP subsystem: %v", err)
	}

	return nil
}

// Close closes the SFTP client first, then the SSH client.
func (c *GoClient) close() error {
	var errs error

	if c.sftp != nil {
		if err := c.sftp.Close(); err != nil {
			errs = errors.Join(err, errs)
		}
	}

	if c.ssh != nil {
		if err := c.ssh.Close(); err != nil {
			errs = errors.Join(err, errs)
		}
	}

	return errs
}
