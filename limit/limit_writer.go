package limit

import (
	"errors"
	"io"
)

// Writer limits the number of bytes written to the underlying writer, and
// then returns ErrMaxLenExceeded if there are any attempts to write beyond the
// limit.
type Writer struct {
	io.Writer
	Max     int
	Written int
}

var ErrMaxLenExceeded = errors.New("max length exceeded")

func (l Writer) Write(p []byte) (n int, err error) {
	allowed := l.Max - l.Written
	if allowed == 0 && len(p) > 0 {
		return 0, ErrMaxLenExceeded
	}
	cut := false
	if len(p) > allowed {
		p = p[:allowed]
		cut = true
	}
	n, err = l.Writer.Write(p)
	l.Written += n
	if cut && err == nil {
		err = ErrMaxLenExceeded
	}
	return n, err
}
