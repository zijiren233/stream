package stream

import (
	"bufio"
	"io"
)

type Reader struct {
	r   io.Reader
	n   int
	err error

	bufferSize int
}

type ReaderConf func(*Reader)

func WithReaderBufferSize(size int) ReaderConf {
	return func(r *Reader) {
		r.bufferSize = size
	}
}

func NewReader(r io.Reader, conf ...ReaderConf) *Reader {
	reader := new(Reader)

	for _, c := range conf {
		c(reader)
	}

	if reader.bufferSize != 0 {
		reader.r = bufio.NewReaderSize(r, reader.bufferSize)
	} else {
		reader.r = r
	}

	return reader
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
