package stream

import (
	"io"
)

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
		*t = U8(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU8() (t uint8, err error) {
	r.U8(&t)
	err = r.err
	return
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
		*t = U16BE(buf[:2])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU16BE() (t uint16, err error) {
	r.U16BE(&t)
	err = r.err
	return
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
		*t = U16LE(buf[:2])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU16LE() (t uint16, err error) {
	r.U16LE(&t)
	err = r.err
	return
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
		*t = U24BE(buf[:3])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU24BE() (t uint32, err error) {
	r.U24BE(&t)
	err = r.err
	return
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
		*t = U24LE(buf[:3])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU24LE() (t uint32, err error) {
	r.U24LE(&t)
	err = r.err
	return
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
		*t = U32BE(buf[:4])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU32BE() (t uint32, err error) {
	r.U32BE(&t)
	err = r.err
	return
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
		*t = U32LE(buf[:4])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU32LE() (t uint32, err error) {
	r.U32LE(&t)
	err = r.err
	return
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
		*t = U40BE(buf[:5])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU40BE() (t uint64, err error) {
	r.U40BE(&t)
	err = r.err
	return
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
		*t = U40LE(buf[:5])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU40LE() (t uint64, err error) {
	r.U40LE(&t)
	err = r.err
	return
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
		*t = U48BE(buf[:6])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU48BE() (t uint64, err error) {
	r.U48BE(&t)
	err = r.err
	return
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
		*t = U48LE(buf[:6])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU48LE() (t uint64, err error) {
	r.U48LE(&t)
	err = r.err
	return
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
		*t = U56BE(buf[:7])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU56BE() (t uint64, err error) {
	r.U56BE(&t)
	err = r.err
	return
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
		*t = U56LE(buf[:7])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU56LE() (t uint64, err error) {
	r.U56LE(&t)
	err = r.err
	return
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
		*t = U64BE(buf)
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU64BE() (t uint64, err error) {
	r.U64BE(&t)
	err = r.err
	return
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
		*t = U64LE(buf)
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetU64LE() (t uint64, err error) {
	r.U64LE(&t)
	err = r.err
	return
}
