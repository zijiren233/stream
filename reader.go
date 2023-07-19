package stream

import (
	"io"
)

type Reader struct {
	r   io.Reader
	n   int
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: r,
	}
}

func (r *Reader) Error() error {
	return r.err
}

// N returns the number of bytes read at last read.
// Instead of the total.
func (r *Reader) N() int {
	return r.n
}

func (r *Reader) Reset() {
	r.n = 0
	r.err = nil
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Byte(t *byte) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:1])
	if err != nil {
		r.err = err
	} else {
		*t = buf[0]
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Bytes(t []byte) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, t)
	if err != nil {
		r.err = err
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I8(t *int8) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:1])
	if err != nil {
		r.err = err
	} else {
		*t = int8(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I16BE(t *int16) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:2])
	if err != nil {
		r.err = err
	} else {
		*t = int16(buf[0])<<8 | int16(buf[1])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I16LE(t *int16) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:2])
	if err != nil {
		r.err = err
	} else {
		*t = int16(buf[1])<<8 | int16(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I24BE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:3])
	if err != nil {
		r.err = err
	} else {
		*t = int32(buf[0])<<16 | int32(buf[1])<<8 | int32(buf[2])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I24LE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:3])
	if err != nil {
		r.err = err
	} else {
		*t = int32(buf[2])<<16 | int32(buf[1])<<8 | int32(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I32BE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:4])
	if err != nil {
		r.err = err
	} else {
		*t = int32(buf[0])<<24 | int32(buf[1])<<16 | int32(buf[2])<<8 | int32(buf[3])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I32LE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:4])
	if err != nil {
		r.err = err
	} else {
		*t = int32(buf[3])<<24 | int32(buf[2])<<16 | int32(buf[1])<<8 | int32(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I40BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:5])
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[0])<<32 | int64(buf[1])<<24 | int64(buf[2])<<16 | int64(buf[3])<<8 | int64(buf[4])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I40LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:5])
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[4])<<32 | int64(buf[3])<<24 | int64(buf[2])<<16 | int64(buf[1])<<8 | int64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I48BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:6])
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[0])<<40 | int64(buf[1])<<32 | int64(buf[2])<<24 | int64(buf[3])<<16 | int64(buf[4])<<8 | int64(buf[5])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I48LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:6])
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[5])<<40 | int64(buf[4])<<32 | int64(buf[3])<<24 | int64(buf[2])<<16 | int64(buf[1])<<8 | int64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I56BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:7])
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[0])<<48 | int64(buf[1])<<40 | int64(buf[2])<<32 | int64(buf[3])<<24 | int64(buf[4])<<16 | int64(buf[5])<<8 | int64(buf[6])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I56LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:7])
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[6])<<48 | int64(buf[5])<<40 | int64(buf[4])<<32 | int64(buf[3])<<24 | int64(buf[2])<<16 | int64(buf[1])<<8 | int64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I64BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf)
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[0])<<56 | int64(buf[1])<<48 | int64(buf[2])<<40 | int64(buf[3])<<32 | int64(buf[4])<<24 | int64(buf[5])<<16 | int64(buf[6])<<8 | int64(buf[7])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) I64LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf)
	if err != nil {
		r.err = err
	} else {
		*t = int64(buf[7])<<56 | int64(buf[6])<<48 | int64(buf[5])<<40 | int64(buf[4])<<32 | int64(buf[3])<<24 | int64(buf[2])<<16 | int64(buf[1])<<8 | int64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U8(t *uint8) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:1])
	if err != nil {
		r.err = err
	} else {
		*t = buf[0]
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U16BE(t *uint16) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:2])
	if err != nil {
		r.err = err
	} else {
		*t = uint16(buf[0])<<8 | uint16(buf[1])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U16LE(t *uint16) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:2])
	if err != nil {
		r.err = err
	} else {
		*t = uint16(buf[1])<<8 | uint16(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U24BE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:3])
	if err != nil {
		r.err = err
	} else {
		*t = uint32(buf[0])<<16 | uint32(buf[1])<<8 | uint32(buf[2])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U24LE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:3])
	if err != nil {
		r.err = err
	} else {
		*t = uint32(buf[2])<<16 | uint32(buf[1])<<8 | uint32(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U32BE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:4])
	if err != nil {
		r.err = err
	} else {
		*t = uint32(buf[0])<<24 | uint32(buf[1])<<16 | uint32(buf[2])<<8 | uint32(buf[3])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U32LE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:4])
	if err != nil {
		r.err = err
	} else {
		*t = uint32(buf[3])<<24 | uint32(buf[2])<<16 | uint32(buf[1])<<8 | uint32(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U40BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:5])
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[0])<<32 | uint64(buf[1])<<24 | uint64(buf[2])<<16 | uint64(buf[3])<<8 | uint64(buf[4])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U40LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:5])
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[4])<<32 | uint64(buf[3])<<24 | uint64(buf[2])<<16 | uint64(buf[1])<<8 | uint64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U48BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:6])
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[0])<<40 | uint64(buf[1])<<32 | uint64(buf[2])<<24 | uint64(buf[3])<<16 | uint64(buf[4])<<8 | uint64(buf[5])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U48LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:6])
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[5])<<40 | uint64(buf[4])<<32 | uint64(buf[3])<<24 | uint64(buf[2])<<16 | uint64(buf[1])<<8 | uint64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U56BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:7])
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[0])<<48 | uint64(buf[1])<<40 | uint64(buf[2])<<32 | uint64(buf[3])<<24 | uint64(buf[4])<<16 | uint64(buf[5])<<8 | uint64(buf[6])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U56LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf[:7])
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[6])<<48 | uint64(buf[5])<<40 | uint64(buf[4])<<32 | uint64(buf[3])<<24 | uint64(buf[2])<<16 | uint64(buf[1])<<8 | uint64(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U64BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf)
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[0])<<56 | uint64(buf[1])<<48 | uint64(buf[2])<<40 | uint64(buf[3])<<32 | uint64(buf[4])<<24 | uint64(buf[5])<<16 | uint64(buf[6])<<8 | uint64(buf[7])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) U64LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	n, err := io.ReadFull(r.r, buf)
	if err != nil {
		r.err = err
	} else {
		*t = uint64(buf[7])<<56 | uint64(buf[6])<<48 | uint64(buf[5])<<40 | uint64(buf[4])<<32 | uint64(buf[3])<<24 | uint64(buf[2])<<16 | uint64(buf[1])<<8 | uint64(buf[0])
	}
	r.n = n
	return r
}
