package stream

import "io"

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
		*t = I8(buf[0])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI8() (t int8, err error) {
	r.I8(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I16BE(buf[:2])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI16BE() (t int16, err error) {
	r.I16BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I16LE(buf[:2])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI16LE() (t int16, err error) {
	r.I16LE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I24BE(buf[:3])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI24BE() (t int32, err error) {
	r.I24BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I24LE(buf[:3])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI24LE() (t int32, err error) {
	r.I24LE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I32BE(buf[:4])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI32BE() (t int32, err error) {
	r.I32BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I32LE(buf[:4])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI32LE() (t int32, err error) {
	r.I32LE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I40BE(buf[:5])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI40BE() (t int64, err error) {
	r.I40BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I40LE(buf[:5])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI40LE() (t int64, err error) {
	r.I40LE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I48BE(buf[:6])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI48BE() (t int64, err error) {
	r.I48BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I48LE(buf[:6])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI48LE() (t int64, err error) {
	r.I48LE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I56BE(buf[:7])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI56BE() (t int64, err error) {
	r.I56BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I56LE(buf[:7])
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI56LE() (t int64, err error) {
	r.I56LE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I64BE(buf)
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI64BE() (t int64, err error) {
	r.I64BE(&t)
	err = r.err
	r.err = nil
	return
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
		*t = I64LE(buf)
	}
	r.n = n
	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) GetI64LE() (t int64, err error) {
	r.I64LE(&t)
	err = r.err
	r.err = nil
	return
}
