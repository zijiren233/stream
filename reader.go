package stream

import (
	"bufio"
	"io"
)

type Reader struct {
	r     io.Reader
	n     int
	total int
	err   error

	buffered bool
	closed   bool
}

type ReaderConf func(*Reader)

func WithReaderBuffer(buffered bool) ReaderConf {
	return func(r *Reader) {
		r.buffered = buffered
	}
}

func NewReader(r io.Reader, conf ...ReaderConf) *Reader {
	reader := new(Reader)

	for _, c := range conf {
		c(reader)
	}

	if reader.buffered {
		buf := bufReaderPool.Get().(*bufio.Reader)
		buf.Reset(r)
		reader.r = buf
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

func (r *Reader) Total() int {
	return r.total
}

func (r *Reader) Reset() {
	r.n = 0
	r.total = 0
	r.err = nil
}

func (r *Reader) Close() error {
	if r.closed {
		return ErrAlreadyClosed
	}
	r.closed = true
	if r.buffered {
		bufReaderPool.Put(r.r.(*bufio.Reader))
	}
	r.r = nil

	return nil
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Byte(t *byte) *Reader {
	if r.err != nil {
		return r
	}
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = io.ReadFull(r.r, buf[:1])
	if r.err == nil {
		*t = buf[0]
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
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = io.ReadFull(r.r, t)
	r.total += r.n

	return r
}
