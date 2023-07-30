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

func (w *WriteSeeker) Seek(offset int64, whence int) *WriteSeeker {
	if w.err != nil {
		return w
	}
	_, w.err = w.ws.Seek(offset, whence)
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

func (w *WriteSeeker) F32BE(t float32) *WriteSeeker {
	w.Writer.F32BE(t)
	return w
}

func (w *WriteSeeker) F32LE(t float32) *WriteSeeker {
	w.Writer.F32LE(t)
	return w
}

func (w *WriteSeeker) F64BE(t float64) *WriteSeeker {
	w.Writer.F64BE(t)
	return w
}

func (w *WriteSeeker) I16BE(t int16) *WriteSeeker {
	w.Writer.I16BE(t)
	return w
}

func (w *WriteSeeker) I16LE(t int16) *WriteSeeker {
	w.Writer.I16LE(t)
	return w
}

func (w *WriteSeeker) I24BE(t int32) *WriteSeeker {
	w.Writer.I24BE(t)
	return w
}

func (w *WriteSeeker) I24LE(t int32) *WriteSeeker {
	w.Writer.I24LE(t)
	return w
}

func (w *WriteSeeker) I32BE(t int32) *WriteSeeker {
	w.Writer.I32BE(t)
	return w
}

func (w *WriteSeeker) I32LE(t int32) *WriteSeeker {
	w.Writer.I32LE(t)
	return w
}

func (w *WriteSeeker) I40BE(t int64) *WriteSeeker {
	w.Writer.I40BE(t)
	return w
}

func (w *WriteSeeker) I40LE(t int64) *WriteSeeker {
	w.Writer.I40LE(t)
	return w
}

func (w *WriteSeeker) I48BE(t int64) *WriteSeeker {
	w.Writer.I48BE(t)
	return w
}

func (w *WriteSeeker) I48LE(t int64) *WriteSeeker {
	w.Writer.I48LE(t)
	return w
}

func (w *WriteSeeker) I56BE(t int64) *WriteSeeker {
	w.Writer.I56BE(t)
	return w
}

func (w *WriteSeeker) I56LE(t int64) *WriteSeeker {
	w.Writer.I56LE(t)
	return w
}

func (w *WriteSeeker) I64BE(t int64) *WriteSeeker {
	w.Writer.I64BE(t)
	return w
}

func (w *WriteSeeker) I64LE(t int64) *WriteSeeker {
	w.Writer.I64LE(t)
	return w
}

func (w *WriteSeeker) I8(t int8) *WriteSeeker {
	w.Writer.I8(t)
	return w
}

func (w *WriteSeeker) U16BE(t uint16) *WriteSeeker {
	w.Writer.U16BE(t)
	return w
}

func (w *WriteSeeker) U16LE(t uint16) *WriteSeeker {
	w.Writer.U16LE(t)
	return w
}

func (w *WriteSeeker) U24BE(t uint32) *WriteSeeker {
	w.Writer.U24BE(t)
	return w
}

func (w *WriteSeeker) U24LE(t uint32) *WriteSeeker {
	w.Writer.U24LE(t)
	return w
}

func (w *WriteSeeker) U32BE(t uint32) *WriteSeeker {
	w.Writer.U32BE(t)
	return w
}

func (w *WriteSeeker) U32LE(t uint32) *WriteSeeker {
	w.Writer.U32LE(t)
	return w
}

func (w *WriteSeeker) U40BE(t uint64) *WriteSeeker {
	w.Writer.U40BE(t)
	return w
}

func (w *WriteSeeker) U40LE(t uint64) *WriteSeeker {
	w.Writer.U40LE(t)
	return w
}

func (w *WriteSeeker) U48BE(t uint64) *WriteSeeker {
	w.Writer.U48BE(t)
	return w
}

func (w *WriteSeeker) U48LE(t uint64) *WriteSeeker {
	w.Writer.U48LE(t)
	return w
}

func (w *WriteSeeker) U56BE(t uint64) *WriteSeeker {
	w.Writer.U56BE(t)
	return w
}

func (w *WriteSeeker) U56LE(t uint64) *WriteSeeker {
	w.Writer.U56LE(t)
	return w
}

func (w *WriteSeeker) U64BE(t uint64) *WriteSeeker {
	w.Writer.U64BE(t)
	return w
}

func (w *WriteSeeker) U64LE(t uint64) *WriteSeeker {
	w.Writer.U64LE(t)
	return w
}

func (w *WriteSeeker) U8(t uint8) *WriteSeeker {
	w.Writer.U8(t)
	return w
}
