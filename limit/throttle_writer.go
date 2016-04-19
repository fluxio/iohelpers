package limit

import (
	"io"
	"time"
)

type Quota struct {
	Limit  uint64
	Within time.Duration
}

type Callback func()

type ThrottleWriter struct {
	io.Writer
	Quota

	Callback Callback
	Start    time.Time
	Count    uint64
}

// Throttle Writer silently throws away messages over the limit.
func (t ThrottleWriter) Write(p []byte) (n int, err error) {

	now := time.Now().UTC()
	if now.Sub(t.Start) < t.Within {
		if t.Count > t.Limit {
			if t.Callback != nil {
				t.Callback()
			}
			return 0, nil
		}
		t.Count++
	} else {
		t.Start = now
		t.Count = 0
	}

	return t.Writer.Write(p)
}
