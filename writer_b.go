package stream

func (w *Writer) Bool(s bool) *Writer {
	if w.err != nil {
		return w
	}
	PutBool(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}
