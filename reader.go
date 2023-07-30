package stream

import (
	"io"
)

type Reader struct {
	r     io.Reader
	n     int
	total int
	err   error
	buf   []byte
}

func NewReader(r io.Reader) *Reader {
	reader := new(Reader)

	reader.r = r

	return &Reader{
		r:   r,
		buf: make([]byte, 8),
	}
}

func (r *Reader) Error() error {
	return r.err
}

// N returns the number of bytes read at last read.
// Instead of the total.
func (r *Reader) N() int {
	return r.n
}

func (r *Reader) Total() int {
	return r.total
}

func (r *Reader) Reset() {
	r.n = 0
	r.total = 0
	r.err = nil
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Byte(t *byte) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:1])
	if r.err == nil {
		*t = r.buf[0]
	}
	r.total += r.n

	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Bytes(t []byte) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, t)
	r.total += r.n

	return r
}
