package stream

func (w *Writer) F32BE(s float32) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutF32BE(buf[:4], s)
	w.n, w.err = w.w.Write(buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F32LE(s float32) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutF32LE(buf[:4], s)
	w.n, w.err = w.w.Write(buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F64BE(s float64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutF64BE(buf[:8], s)
	w.n, w.err = w.w.Write(buf[:8])
	w.total += w.n

	return w
}

func (w *Writer) F64LE(s float64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutF64LE(buf[:8], s)
	w.n, w.err = w.w.Write(buf[:8])
	w.total += w.n

	return w
}
