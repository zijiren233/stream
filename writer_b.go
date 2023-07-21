package stream

func (w *Writer) Bool(s bool) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutBool(buf[:1], s)
	w.n, w.err = w.w.Write(buf[:1])
	w.total += w.n

	return w
}
