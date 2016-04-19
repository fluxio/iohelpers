package counting

import "io"

// Closer implements ReadCloser over an underlying io.Reader with a
// trivial Close that just increments a counter.  Closer is intended for
// use in tests, and is not thread safe.
type Closer struct {
	io.Reader
	Count int
}

// Close increments a counter.  If the backing io.Reader is also an io.Closer,
// we also try to close it.
func (c *Closer) Close() error {
	c.Count++
	if cl, ok := c.Reader.(io.Closer); ok {
		return cl.Close()
	}
	return nil
}
