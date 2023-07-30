package stream

import (
	"io"
)

type Writer struct {
	w     io.Writer
	n     int
	total int
	err   error
	buf   []byte
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w:   w,
		buf: make([]byte, 8),
	}
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

func (w *Writer) Byte(s byte) *Writer {
	if w.err != nil {
		return w
	}
	w.n, w.err = w.w.Write([]byte{s})
	w.total += w.n

	return w
}

func (w *Writer) Bytes(s []byte) *Writer {
	if w.err != nil {
		return w
	}
	w.n, w.err = w.w.Write(s)
	w.total += w.n

	return w
}

func (w *Writer) I8(s int8) *Writer {
	if w.err != nil {
		return w
	}
	WriteI8(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) I16BE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	WriteI16BE(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) I16LE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	WriteI16LE(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) I24BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	WriteI24BE(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) I24LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	WriteI24LE(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) I32BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	WriteI32BE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) I32LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	WriteI32LE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) I40BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI40BE(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) I40LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI40LE(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) I48BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI48BE(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) I48LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI48LE(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) I56BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI56BE(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) I56LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI56LE(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) I64BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI64BE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) I64LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	WriteI64LE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) U8(s uint8) *Writer {
	if w.err != nil {
		return w
	}
	WriteU8(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) U16BE(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	WriteU16BE(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) U16LE(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	WriteU16LE(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) U24BE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	WriteU24BE(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) U24LE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	WriteU24LE(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) U32BE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	WriteU32BE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) U32LE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	WriteU32LE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) U40BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU40BE(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) U40LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU40LE(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) U48BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU48BE(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) U48LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU48LE(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) U56BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU56BE(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) U56LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU56LE(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) U64BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU64BE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) U64LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	WriteU64LE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) F32BE(s float32) *Writer {
	if w.err != nil {
		return w
	}
	WriteF32BE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F32LE(s float32) *Writer {
	if w.err != nil {
		return w
	}
	WriteF32LE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F64BE(s float64) *Writer {
	if w.err != nil {
		return w
	}
	WriteF64BE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf[:8])
	w.total += w.n

	return w
}

func (w *Writer) F64LE(s float64) *Writer {
	if w.err != nil {
		return w
	}
	WriteF64LE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf[:8])
	w.total += w.n

	return w
}

func (w *Writer) Bool(s bool) *Writer {
	if w.err != nil {
		return w
	}
	WriteBool(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) String(s string) *Writer {
	if w.err != nil {
		return w
	}

	w.n, w.err = w.w.Write(StringToBytes(s))
	w.total += w.n

	return w
}
