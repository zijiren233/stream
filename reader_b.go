package stream

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Bool(t *bool) *Reader {
	if r.err != nil {
		return r
	}
	if r.closed {
		r.err = ErrAlreadyClosed
		return r
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	r.n, r.err = r.r.Read(buf[:1])
	if r.err == nil {
		*t = Bool(buf[:1])
	}
	r.total += r.n

	return r
}
