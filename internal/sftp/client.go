package sftp

import (
	"context"
	"fmt"
	"io"

	"github.com/artefactual-sdps/enduro/internal/async"
)

// AuthError represents an SFTP authentication error.
type AuthError struct {
	Message string
}

// Error implements the error interface.
func (e *AuthError) Error() string {
	return fmt.Sprintf("auth: %s", e.Message)
}

// NewAuthError returns a pointer to a new AuthError from the underlying |e|
// error message.
func NewAuthError(e error) error {
	return &AuthError{Message: e.Error()}
}

// A Client manages the transmission of data over SFTP.
//
// Implementations of the Client interface handle the connection details,
// authentication, and other intricacies associated with different SFTP
// servers and protocols.
type Client interface {
	// Delete removes dest from the SFTP server.
	Delete(ctx context.Context, dest string) error
	// Upload asynchronously copies data from the src reader to the specified
	// dest on the SFTP server.
	Upload(ctx context.Context, src io.Reader, dest string) (remotePath string, upload async.Upload, err error)
}
