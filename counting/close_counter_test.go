package counting

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCloser(t *testing.T) {
	cc := &Closer{Reader: strings.NewReader("hello")}
	if cc.Count != 0 {
		t.Errorf("close count should be 0 initially")
	}
	readBytes, err := ioutil.ReadAll(cc)
	if err != nil {
		t.Fatalf("Unexpected error reading from Closer: %v", err)
	}
	if string(readBytes) != "hello" {
		t.Errorf("Closer unexpectedly mangled read bytes: %q", readBytes)
	}
	for i := 0; i < 3; i++ {
		err = cc.Close()
		if err != nil {
			t.Fatalf("Unexpected error closing Closer: %v", err)
		}
		if cc.Count != i+1 {
			t.Fatalf("Closer should increment; expected %d, got: %d",
				i+1, cc.Count)
		}
	}

	// Test stacked Closer.
	// This exercises the case where Reader is also an io.Closer.
	cc2 := &Closer{Reader: strings.NewReader("goodbye")}
	cc3 := &Closer{Reader: cc2}
	err = cc3.Close()
	if err != nil {
		t.Fatalf("Unexpected error closing Closer: %v", err)
	}
	if cc3.Count != 1 {
		t.Errorf("Outer Closer should have 1 close; got %d", cc3.Count)
	}
	if cc2.Count != 1 {
		t.Errorf("Inner Closer should have 1 close; got %d", cc2.Count)
	}
}
