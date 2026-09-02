package groupcache

import (
	"io"
)

type ByteView struct {
	b []byte
	s string
}

func (v ByteView) Len() int { _ = "STUB: not implemented"; return 0 }

func (v ByteView) ByteSlice() []byte { _ = "STUB: not implemented"; return nil }

func (v ByteView) String() string { _ = "STUB: not implemented"; return "" }

func (v ByteView) At(i int) byte { _ = "STUB: not implemented"; return 0 }

func (v ByteView) Slice(from, to int) ByteView { _ = "STUB: not implemented"; return *new(ByteView) }

func (v ByteView) SliceFrom(from int) ByteView { _ = "STUB: not implemented"; return *new(ByteView) }

func (v ByteView) Copy(dest []byte) int { _ = "STUB: not implemented"; return 0 }

func (v ByteView) Equal(b2 ByteView) bool { _ = "STUB: not implemented"; return false }

func (v ByteView) EqualString(s string) bool { _ = "STUB: not implemented"; return false }

func (v ByteView) EqualBytes(b2 []byte) bool { _ = "STUB: not implemented"; return false }

func (v ByteView) Reader() io.ReadSeeker { _ = "STUB: not implemented"; return *new(io.ReadSeeker) }

func (v ByteView) ReadAt(p []byte, off int64) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (v ByteView) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
