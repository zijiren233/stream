package stream

import (
	"io"
)

type (
	BitReader struct {
		r   io.Reader
		buf [1]byte
		pos int

		rt io.WriterTo
	}
)

func NewBitReader(r io.Reader) *BitReader {
	reader := &BitReader{r: r, pos: 8}
	if rt, ok := reader.r.(io.WriterTo); ok {
		reader.rt = rt
	}
	return reader
}

// ReadBit reads a single bit from the stream.
func (r *BitReader) ReadBit() (byte, error) {
	// if pos more than 8, represents that the bit reader is aligned
	// and we need to read a byte from the stream
	if r.pos >= 8 {
		_, err := r.r.Read(r.buf[:])
		if err != nil {
			return 0, err
		}
		r.pos = 0
	}
	b := r.buf[0] >> (7 - r.pos) & 1
	r.pos++
	return b, nil
}

// ReadBits reads bits from the stream.
// The bits are written to p, starting at bit start and ending at bit end.
// including start, excluding end.
func (r *BitReader) ReadBits(p []byte, start, end int) error {
	if start < 0 || end < 0 || start >= end || len(p) < (end-start+7)/8 {
		return io.ErrShortBuffer
	}

	// bit reader is already aligned
	if r.pos >= 8 {
		// priority handle aligned read
		if start%8 == 0 && end%8 == 0 {
			_, err := io.ReadFull(r.r, p[start/8:end/8])
			return err
		} else if (end-start)%8 == 0 {
			buf := make([]byte, (end-start)/8)
			_, err := io.ReadFull(r.r, buf)
			if err != nil {
				return err
			}
			// start maybe not aligned to 8
			return CopyBits(buf, p, start)
		}
	}

	// bit reader is not aligned
	for ; start < end; start++ {
		b, err := r.ReadBit()
		if err != nil {
			return err
		}

		if b == 1 {
			p[start/8] |= b << (7 - start%8)
		} else {
			// clear the bit
			p[start/8] &= ^(1 << (7 - start%8))
		}
	}
	return nil
}

// ReadByte reads a single byte from the stream.
func (r *BitReader) ReadByte() (byte, error) {
	if r.pos >= 8 {
		_, err := r.r.Read(r.buf[:])
		if err != nil {
			return 0, err
		}
		return r.buf[0], nil
	}

	err := r.ReadBits(r.buf[:], 0, 8)
	if err != nil {
		return r.buf[0], err
	}
	return r.buf[0], nil
}

// Read reads len(p) bytes from the stream.
func (r *BitReader) Read(p []byte) (int, error) {
	if r.pos >= 8 {
		return r.r.Read(p)
	}

	n := 0
	for i := range p {
		b, err := r.ReadByte()
		if err != nil {
			return n, err
		}
		p[i] = b
		n++
	}
	return n, nil
}

func (r *BitReader) SkipBytes(n int64) (int64, error) {
	if r.pos >= 8 {
		return io.CopyN(io.Discard, r.r, n)
	}

	for i := int64(0); i < n; i++ {
		_, err := r.ReadByte()
		if err != nil {
			return i - 1, err
		}
	}
	return n, nil
}

func (r *BitReader) SkipBits(n int64) (int64, error) {
	if r.pos >= 8 && n%8 == 0 {
		return io.CopyN(io.Discard, r.r, n/8)
	}

	for i := int64(0); i < n; i++ {
		_, err := r.ReadBit()
		if err != nil {
			return i - 1, err
		}
	}
	return n, nil
}
