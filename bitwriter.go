package stream

import (
	"io"
)

type BitWriter struct {
	w   io.Writer
	buf [1]byte
	pos uint8
}

func NewBitWriter(w io.Writer) *BitWriter {
	return &BitWriter{
		w: w,
	}
}

func (w *BitWriter) WriteBit(b byte) error {
	if b != 0 {
		w.buf[0] |= 1 << (7 - w.pos)
	} else {
		w.buf[0] &= ^(1 << (7 - w.pos))
	}
	w.pos++
	if w.pos == 8 {
		_, err := w.w.Write(w.buf[:])
		if err != nil {
			return err
		}
		w.pos = 0
	}
	return nil
}

func (w *BitWriter) WriteBits(b []byte, start, end int) error {
	if start < 0 || end > len(b)*8 {
		return io.ErrShortBuffer
	}

	if w.pos == 0 && start%8 == 0 && end%8 == 0 {
		_, err := w.w.Write(b[start/8 : end/8])
		return err
	}

	for i := start; i < end; i++ {
		w.WriteBit(b[i/8] & (1 << (7 - i%8)))
	}
	return nil
}

func (w *BitWriter) WriteByte(b byte) error {
	return w.WriteBits([]byte{b}, 0, 8)
}

func (w *BitWriter) Write(p []byte) (int, error) {
	if w.pos == 0 {
		return w.w.Write(p)
	}

	n := 0
	for _, b := range p {
		err := w.WriteByte(b)
		if err != nil {
			return n, err
		}
		n++
	}
	return n, nil
}
