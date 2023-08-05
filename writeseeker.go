package stream

import "io"

type WriteSeeker struct {
	*Writer
	ws io.WriteSeeker
}

func NewWriteSeeker(ws io.WriteSeeker, o Order) *WriteSeeker {
	return &WriteSeeker{
		Writer: NewWriter(ws, o),
		ws:     ws,
	}
}

func (w *WriteSeeker) Seek(offset int64, whence int) *WriteSeeker {
	if w.err != nil {
		return w
	}
	_, w.err = w.ws.Seek(offset, whence)
	return w
}

func (w *WriteSeeker) Write(data any) *WriteSeeker {
	w.Writer.Write(data)
	return w
}

func (w *WriteSeeker) Bool(t bool) *WriteSeeker {
	w.Writer.Bool(t)
	return w
}

func (w *WriteSeeker) Byte(t byte) *WriteSeeker {
	w.Writer.Byte(t)
	return w
}

func (w *WriteSeeker) Bytes(t []byte) *WriteSeeker {
	w.Writer.Bytes(t)
	return w
}

func (w *WriteSeeker) F32(t float32) *WriteSeeker {
	w.Writer.F32(t)
	return w
}

func (w *WriteSeeker) F64(t float64) *WriteSeeker {
	w.Writer.F64(t)
	return w
}

func (w *WriteSeeker) I16(t int16) *WriteSeeker {
	w.Writer.I16(t)
	return w
}

func (w *WriteSeeker) I24(t int32) *WriteSeeker {
	w.Writer.I24(t)
	return w
}

func (w *WriteSeeker) I32(t int32) *WriteSeeker {
	w.Writer.I32(t)
	return w
}

func (w *WriteSeeker) I40(t int64) *WriteSeeker {
	w.Writer.I40(t)
	return w
}

func (w *WriteSeeker) I48(t int64) *WriteSeeker {
	w.Writer.I48(t)
	return w
}

func (w *WriteSeeker) I56(t int64) *WriteSeeker {
	w.Writer.I56(t)
	return w
}

func (w *WriteSeeker) I64(t int64) *WriteSeeker {
	w.Writer.I64(t)
	return w
}

func (w *WriteSeeker) I8(t int8) *WriteSeeker {
	w.Writer.I8(t)
	return w
}

func (w *WriteSeeker) U16(t uint16) *WriteSeeker {
	w.Writer.U16(t)
	return w
}

func (w *WriteSeeker) U24(t uint32) *WriteSeeker {
	w.Writer.U24(t)
	return w
}

func (w *WriteSeeker) U32(t uint32) *WriteSeeker {
	w.Writer.U32(t)
	return w
}

func (w *WriteSeeker) U40(t uint64) *WriteSeeker {
	w.Writer.U40(t)
	return w
}

func (w *WriteSeeker) U48(t uint64) *WriteSeeker {
	w.Writer.U48(t)
	return w
}

func (w *WriteSeeker) U56(t uint64) *WriteSeeker {
	w.Writer.U56(t)
	return w
}

func (w *WriteSeeker) U64(t uint64) *WriteSeeker {
	w.Writer.U64(t)
	return w
}

func (w *WriteSeeker) U8(t uint8) *WriteSeeker {
	w.Writer.U8(t)
	return w
}

func (w *WriteSeeker) String(t string) *WriteSeeker {
	w.Writer.String(t)
	return w
}
