package stream

import "io"

type ReadSeeker struct {
	*Reader
	rs io.ReadSeeker
}

func NewReadSeeker(rs io.ReadSeeker) *ReadSeeker {
	return &ReadSeeker{
		Reader: NewReader(rs),
		rs:     rs,
	}
}

func (r *ReadSeeker) Seek(offset int64, whence int) (int64, error) {
	return r.rs.Seek(offset, whence)
}
