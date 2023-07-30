package stream

func (w *Writer) F32BE(s float32) *Writer {
	if w.err != nil {
		return w
	}
	PutF32BE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F32LE(s float32) *Writer {
	if w.err != nil {
		return w
	}
	PutF32LE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F64BE(s float64) *Writer {
	if w.err != nil {
		return w
	}
	PutF64BE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf[:8])
	w.total += w.n

	return w
}

func (w *Writer) F64LE(s float64) *Writer {
	if w.err != nil {
		return w
	}
	PutF64LE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf[:8])
	w.total += w.n

	return w
}
