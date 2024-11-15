package async

// Upload provides information about an upload happening asynchronously in
// a separate goroutine.
type Upload interface {
	// Bytes returns the number of bytes copied to the SFTP destination.
	Bytes() int64
	// Close closes SFTP connection used for the upload. Close must be called
	// when the upload is complete to prevent memory leaks.
	Close() error
	// Done returns a channel that receives a true value when the upload is
	// complete.  A done signal should not be sent on error.
	Done() chan bool
	// Done returns a channel that receives an error if the upload encounters
	// an error.
	Err() chan error
	// Write implements the io.Writer interface and adds len(p) to the count of
	// bytes uploaded.
	Write(p []byte) (int, error)
}
