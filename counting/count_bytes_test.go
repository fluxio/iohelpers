package counting

import (
	"io"
	"io/ioutil"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReader(t *testing.T) {
	Convey("Bytes read should be counted", t, func() {
		Convey("...fully reading a 0-length reader", func() {
			count := &Reader{Reader: strings.NewReader("")}
			ioutil.ReadAll(count)
			So(count.BytesRead, ShouldEqual, 0)
		})

		Convey("...fully reading a 1-byte reader", func() {
			count := &Reader{Reader: strings.NewReader("a")}
			ioutil.ReadAll(count)
			So(count.BytesRead, ShouldEqual, 1)
		})

		Convey("...fully reading a 12-byte reader", func() {
			count := &Reader{Reader: strings.NewReader("hello, world")}
			ioutil.ReadAll(count)
			So(count.BytesRead, ShouldEqual, 12)
		})

		Convey("...reading a 12-byte reader in chunks", func() {
			count := &Reader{Reader: strings.NewReader("hello, world")}
			buf := make([]byte, 5)
			bytesRead, err := count.Read(buf)
			So(err, ShouldBeNil)
			So(bytesRead, ShouldEqual, 5)
			So(count.BytesRead, ShouldEqual, 5)
			bytesRead, err = count.Read(buf)
			So(err, ShouldBeNil)
			So(bytesRead, ShouldEqual, 5)
			So(count.BytesRead, ShouldEqual, 10)
			bytesRead, err = count.Read(buf)
			So(err, ShouldBeNil)
			So(bytesRead, ShouldEqual, 2)
			So(count.BytesRead, ShouldEqual, 12)
			bytesRead, err = count.Read(buf)
			So(err, ShouldEqual, io.EOF)
			So(bytesRead, ShouldEqual, 0)
			So(count.BytesRead, ShouldEqual, 12)
		})
	})
}
