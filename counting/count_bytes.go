package counting

import (
	"io"
)

// Reader is an io.Reader that counts the bytes read from an underlying
// reader as they are read.
type Reader struct {
	io.Reader
	BytesRead int
}

func (c *Reader) Read(p []byte) (n int, err error) {
	n, err = c.Reader.Read(p)
	c.BytesRead += n
	return
}
