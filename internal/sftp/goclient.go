package sftp

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/dolmen-go/contextio"
	"github.com/go-logr/logr"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// GoClient implements the SFTP service using native Go SSH and SFTP packages.
type GoClient struct {
	cfg    Config
	logger logr.Logger
}

var _ Client = (*GoClient)(nil)

// NewGoClient returns a new GoSFTP client with the given configuration.
func NewGoClient(logger logr.Logger, cfg Config) *GoClient {
	cfg.SetDefaults()

	return &GoClient{cfg: cfg, logger: logger}
}

// Upload writes the data from src to the remote file at dest and returns the
// number of bytes written.  A new SFTP connection is opened before writing, and
// closed when the upload is complete or cancelled.
func (c *GoClient) Upload(ctx context.Context, src io.Reader, dest string) (int64, string, error) {
	conn, err := c.dial(ctx)
	if err != nil {
		return 0, "", err
	}
	defer conn.close()

	// SFTP assumes that "/" is used as the directory separator. See:
	// https://datatracker.ietf.org/doc/html/draft-ietf-secsh-filexfer-02#section-6.2
	remotePath := strings.TrimSuffix(c.cfg.RemoteDir, "/") + "/" + dest

	// Note: Some SFTP servers don't support O_RDWR mode.
	w, err := conn.sftp.OpenFile(remotePath, (os.O_WRONLY | os.O_CREATE | os.O_TRUNC))
	if err != nil {
		return 0, "", fmt.Errorf("SFTP: open remote file %q: %v", dest, err)
	}
	defer w.Close()

	// Use contextio to stop the upload if a context cancellation signal is
	// received.
	bytes, err := io.Copy(contextio.NewWriter(ctx, w), contextio.NewReader(ctx, src))
	if err != nil {
		return 0, "", fmt.Errorf("SFTP: upload to %q: %v", dest, err)
	}

	return bytes, remotePath, nil
}

// Delete removes the data from dest. A new SFTP connection is opened before
// removing the file, and closed when the delete is complete.
func (c *GoClient) Delete(ctx context.Context, dest string) error {
	conn, err := c.dial(ctx)
	if err != nil {
		return fmt.Errorf("SFTP: unable to dial: %w", err)
	}
	defer conn.close()

	// SFTP assumes that "/" is used as the directory separator. See:
	// https://datatracker.ietf.org/doc/html/draft-ietf-secsh-filexfer-02#section-6.2
	remotePath := strings.TrimSuffix(c.cfg.RemoteDir, "/") + "/" + dest

	if err := conn.sftp.Remove(remotePath); err != nil {
		head := fmt.Sprintf("SFTP: unable to remove file %q", dest)
		if errors.Is(err, fs.ErrNotExist) || errors.Is(err, fs.ErrPermission) {
			return fmt.Errorf("%s: %w", head, err)
		}
		if statusErr, ok := err.(*sftp.StatusError); ok {
			return fmt.Errorf("%s: %s", head, formatStatusError(statusErr))
		}
		return fmt.Errorf("%s: %v", head, err)
	}

	return nil
}

// Dial connects to an SSH host, creates an SFTP client on the connection, then
// returns conn. When conn is no longer needed, conn.close() must be called to
// prevent leaks.
func (c *GoClient) dial(ctx context.Context) (*connection, error) {
	var (
		conn connection
		err  error
	)

	conn.ssh, err = sshConnect(ctx, c.logger, c.cfg)
	if err != nil {
		return nil, fmt.Errorf("SSH: %v", err)
	}

	conn.sftp, err = sftp.NewClient(conn.ssh)
	if err != nil {
		_ = conn.ssh.Close()
		return nil, fmt.Errorf("start SFTP subsystem: %v", err)
	}

	return &conn, nil
}

var statusCodeRegex = regexp.MustCompile(`\(SSH_[A-Z_]+\)$`)

// formatStatusError extracts/formats the SFTP status error code and message.
func formatStatusError(err *sftp.StatusError) string {
	var (
		code    string
		codeMsg = err.FxCode()
	)

	// Find the first match in the error, removing surrounding parentheses.
	matches := statusCodeRegex.FindStringSubmatch(err.Error())
	if len(matches) > 0 {
		code = matches[0][1 : len(matches[0])-1]
	} else {
		code = strconv.FormatUint(uint64(err.Code), 10)
	}

	return fmt.Sprintf("%s (%s)", codeMsg, code)
}

type connection struct {
	ssh  *ssh.Client
	sftp *sftp.Client
}

// close closes the SFTP client first, then the SSH client.
func (conn *connection) close() error {
	var errs error

	if conn.sftp != nil {
		if err := conn.sftp.Close(); err != nil {
			errs = errors.Join(err, errs)
		}
	}

	if conn.ssh != nil {
		if err := conn.ssh.Close(); err != nil {
			errs = errors.Join(err, errs)
		}
	}

	return errs
}
