package stream

import (
	"io"
)

type Reader struct {
	r   io.Reader
	n   int
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: r,
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

func (r *Reader) Reset() {
	r.n = 0
	r.err = nil
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Byte(t *byte) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:1])
	if err != nil {
		r.err = err
	} else {
		*t = buf[0]
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Bytes(t []byte) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, t)
	if err != nil {
		r.err = err
	}
	r.n = n
	return r
}