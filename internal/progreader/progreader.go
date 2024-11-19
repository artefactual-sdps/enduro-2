package progreader

import (
	"io"
)

type ProgReader struct {
	bytes int
	r     io.Reader
}

func New(r io.Reader) *ProgReader {
	pr := ProgReader{}
	pr.r = io.TeeReader(r, &pr)

	return &pr
}

// Read reads from the progress reader.
func (c *ProgReader) Read(p []byte) (int, error) {
	return c.r.Read(p)
}

// Write adds the length of p to the total number of bytes written on the
// connection.
//
// Write implements the io.Writer interface.
func (c *ProgReader) Write(p []byte) (int, error) {
	n := len(p)
	c.bytes += n
	return n, nil
}

func (c *ProgReader) Bytes() int {
	return c.bytes
}
