package stream

// The error is EOF only if no bytes were read.
// If an EOF happens after reading some but not all the bytes,
// Will returns ErrUnexpectedEOF.
func (r *Reader) Bool(t *bool) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = r.r.Read(r.buf[:1])
	if r.err == nil {
		*t = Bool(r.buf[:1])
	}
	r.total += r.n

	return r
}
