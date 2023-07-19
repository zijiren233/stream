package stream

import "io"

type Writer struct {
	w   io.Writer
	n   int
	err error
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: w,
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

func (w *Writer) Reset() {
	w.n = 0
	w.err = nil
}

func (w *Writer) Byte(s byte) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
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
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := w.w.Write(s)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I8(s int8) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	n, err := w.w.Write(buf[:1])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I16BE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 8)
	buf[1] = byte(s)
	n, err := w.w.Write(buf[:2])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I16LE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	n, err := w.w.Write(buf[:2])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I24BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 16)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s)
	n, err := w.w.Write(buf[:3])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I24LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	n, err := w.w.Write(buf[:3])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I32BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 24)
	buf[1] = byte(s >> 16)
	buf[2] = byte(s >> 8)
	buf[3] = byte(s)
	n, err := w.w.Write(buf[:4])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I32LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	n, err := w.w.Write(buf[:4])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I40BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 32)
	buf[1] = byte(s >> 24)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 8)
	buf[4] = byte(s)
	n, err := w.w.Write(buf[:5])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I40LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	n, err := w.w.Write(buf[:5])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I48BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 40)
	buf[1] = byte(s >> 32)
	buf[2] = byte(s >> 24)
	buf[3] = byte(s >> 16)
	buf[4] = byte(s >> 8)
	buf[5] = byte(s)
	n, err := w.w.Write(buf[:6])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I48LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	buf[5] = byte(s >> 40)
	n, err := w.w.Write(buf[:6])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I56BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 48)
	buf[1] = byte(s >> 40)
	buf[2] = byte(s >> 32)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 16)
	buf[5] = byte(s >> 8)
	buf[6] = byte(s)
	n, err := w.w.Write(buf[:7])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I56LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	buf[5] = byte(s >> 40)
	buf[6] = byte(s >> 48)
	n, err := w.w.Write(buf[:7])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I64BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 56)
	buf[1] = byte(s >> 48)
	buf[2] = byte(s >> 40)
	buf[3] = byte(s >> 32)
	buf[4] = byte(s >> 24)
	buf[5] = byte(s >> 16)
	buf[6] = byte(s >> 8)
	buf[7] = byte(s)
	n, err := w.w.Write(buf)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) I64LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	buf[5] = byte(s >> 40)
	buf[6] = byte(s >> 48)
	buf[7] = byte(s >> 56)
	n, err := w.w.Write(buf)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U8(s uint8) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = s
	n, err := w.w.Write(buf[:1])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U16BE(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 8)
	buf[1] = byte(s)
	n, err := w.w.Write(buf[:2])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U16LE(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	n, err := w.w.Write(buf[:2])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U24BE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 16)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s)
	n, err := w.w.Write(buf[:3])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U24LE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	n, err := w.w.Write(buf[:3])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U32BE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 24)
	buf[1] = byte(s >> 16)
	buf[2] = byte(s >> 8)
	buf[3] = byte(s)
	n, err := w.w.Write(buf[:4])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U32LE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	n, err := w.w.Write(buf[:4])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U40BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 32)
	buf[1] = byte(s >> 24)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 8)
	buf[4] = byte(s)
	n, err := w.w.Write(buf[:5])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U40LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	n, err := w.w.Write(buf[:5])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U48BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 40)
	buf[1] = byte(s >> 32)
	buf[2] = byte(s >> 24)
	buf[3] = byte(s >> 16)
	buf[4] = byte(s >> 8)
	buf[5] = byte(s)
	n, err := w.w.Write(buf[:6])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U48LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	buf[5] = byte(s >> 40)
	n, err := w.w.Write(buf[:6])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U56BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 48)
	buf[1] = byte(s >> 40)
	buf[2] = byte(s >> 32)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 16)
	buf[5] = byte(s >> 8)
	buf[6] = byte(s)
	n, err := w.w.Write(buf[:7])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U56LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	buf[5] = byte(s >> 40)
	buf[6] = byte(s >> 48)
	n, err := w.w.Write(buf[:7])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U64BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s >> 56)
	buf[1] = byte(s >> 48)
	buf[2] = byte(s >> 40)
	buf[3] = byte(s >> 32)
	buf[4] = byte(s >> 24)
	buf[5] = byte(s >> 16)
	buf[6] = byte(s >> 8)
	buf[7] = byte(s)
	n, err := w.w.Write(buf)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U64LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	buf[0] = byte(s)
	buf[1] = byte(s >> 8)
	buf[2] = byte(s >> 16)
	buf[3] = byte(s >> 24)
	buf[4] = byte(s >> 32)
	buf[5] = byte(s >> 40)
	buf[6] = byte(s >> 48)
	buf[7] = byte(s >> 56)
	n, err := w.w.Write(buf)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}
