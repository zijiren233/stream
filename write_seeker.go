package stream

import "io"

type WriteSeeker struct {
	*Writer
	ws io.WriteSeeker
}

func NewWriteSeeker(ws io.WriteSeeker) *WriteSeeker {
	return &WriteSeeker{
		Writer: NewWriter(ws),
		ws:     ws,
	}
}

func (w *WriteSeeker) Seek(offset int64, whence int) (int64, error) {
	return w.ws.Seek(offset, whence)
}
