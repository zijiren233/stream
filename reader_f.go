package stream

import "io"

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) F32BE(t *float32) *Reader {
	if r.err != nil {
		return r
	}
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = io.ReadFull(r.r, buf[:4])
	if r.err == nil {
		*t = F32BE(buf[:4])
	}
	r.total += r.n

	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) F32LE(t *float32) *Reader {
	if r.err != nil {
		return r
	}
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = io.ReadFull(r.r, buf[:4])
	if r.err == nil {
		*t = F32LE(buf[:4])
	}
	r.total += r.n

	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) F64BE(t *float64) *Reader {
	if r.err != nil {
		return r
	}
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = io.ReadFull(r.r, buf[:8])
	if r.err == nil {
		*t = F64BE(buf[:8])
	}
	r.total += r.n

	return r
}

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) F64LE(t *float64) *Reader {
	if r.err != nil {
		return r
	}
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = io.ReadFull(r.r, buf[:8])
	if r.err == nil {
		*t = F64LE(buf[:8])
	}
	r.total += r.n

	return r
}
