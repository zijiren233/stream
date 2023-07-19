package stream

import (
	"bufio"
	"io"

	"github.com/zijiren233/stream/pool"
)

type Writer struct {
	w   io.Writer
	n   int
	err error

	bufferSize int
}

type WriterConf func(*Writer)

func WithBufferSize(size int) WriterConf {
	return func(w *Writer) {
		w.bufferSize = size
	}
}

func NewWriter(w io.Writer, conf ...WriterConf) *Writer {
	writer := new(Writer)

	for _, c := range conf {
		c(writer)
	}

	if writer.bufferSize != 0 {
		writer.w = bufio.NewWriterSize(w, writer.bufferSize)
	} else {
		writer.w = w
	}

	return writer
}

func (w *Writer) Flush() *Writer {
	if w.err != nil {
		return w
	}
	if w.bufferSize != 0 {
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

func (w *Writer) Reset() {
	w.n = 0
	w.err = nil
}

func (w *Writer) Byte(s byte) *Writer {
	if w.err != nil {
		return w
	}
	buf := *pool.BufPool.Get().(*[]byte)
	defer pool.BufPool.Put(&buf)
	n, err := w.w.Write([]byte{s})
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) Bytes(s []byte) *Writer {
	if w.err != nil {
		return w
	}
	buf := *pool.BufPool.Get().(*[]byte)
	defer pool.BufPool.Put(&buf)
	n, err := w.w.Write(s)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}
