package stream

import (
	"bufio"
	"io"
)

type Writer struct {
	w     io.Writer
	n     int
	total int
	err   error

	buffered bool
	closed   bool
}

type WriterConf func(*Writer)

func WithWriterBuffer(buffered bool) WriterConf {
	return func(w *Writer) {
		w.buffered = buffered
	}
}

func NewWriter(w io.Writer, conf ...WriterConf) *Writer {
	writer := new(Writer)

	for _, c := range conf {
		c(writer)
	}

	if writer.buffered {
		buf := bufWriterPool.Get().(*bufio.Writer)
		buf.Reset(w)
		writer.w = buf
	} else {
		writer.w = w
	}

	return writer
}

func (w *Writer) Flush() *Writer {
	if w.err != nil {
		return w
	}
	if w.buffered {
		w.err = w.w.(*bufio.Writer).Flush()
	}
	return w
}

func (w *Writer) Error() error {
	return w.err
}

// N returns the number of bytes write at last write.
// Instead of the total.
func (w *Writer) N() int {
	return w.n
}

func (w *Writer) Total() int {
	return w.total
}

func (w *Writer) Reset() *Writer {
	w.n = 0
	w.total = 0
	w.err = nil
	return w
}

func (w *Writer) Close() *Writer {
	if w.closed {
		return nil
	}
	w.closed = true
	if w.buffered {
		writer := w.w.(*bufio.Writer)
		w.err = writer.Flush()
		bufWriterPool.Put(writer)
	}
	w.w = nil
	return w
}

func (w *Writer) Byte(s byte) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	w.n, w.err = w.w.Write([]byte{s})
	w.total += w.n

	return w
}

func (w *Writer) Bytes(s []byte) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	w.n, w.err = w.w.Write(s)
	w.total += w.n

	return w
}
