package stream

func (w *Writer) I8(s int8) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI8(buf[:1], s)
	w.n, w.err = w.w.Write(buf[:1])
	return w
}

func (w *Writer) I16BE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI16BE(buf[:2], s)
	w.n, w.err = w.w.Write(buf[:2])
	return w
}

func (w *Writer) I16LE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI16LE(buf[:2], s)
	w.n, w.err = w.w.Write(buf[:2])
	return w
}

func (w *Writer) I24BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI24BE(buf[:3], s)
	w.n, w.err = w.w.Write(buf[:3])
	return w
}

func (w *Writer) I24LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI24LE(buf[:3], s)
	w.n, w.err = w.w.Write(buf[:3])
	return w
}

func (w *Writer) I32BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI32BE(buf[:4], s)
	w.n, w.err = w.w.Write(buf[:4])
	return w
}

func (w *Writer) I32LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI32LE(buf[:4], s)
	w.n, w.err = w.w.Write(buf[:4])
	return w
}

func (w *Writer) I40BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI40BE(buf[:5], s)
	w.n, w.err = w.w.Write(buf[:5])
	return w
}

func (w *Writer) I40LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI40LE(buf[:5], s)
	w.n, w.err = w.w.Write(buf[:5])
	return w
}

func (w *Writer) I48BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI48BE(buf[:6], s)
	w.n, w.err = w.w.Write(buf[:6])
	return w
}

func (w *Writer) I48LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI48LE(buf[:6], s)
	w.n, w.err = w.w.Write(buf[:6])
	return w
}

func (w *Writer) I56BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI56BE(buf[:7], s)
	w.n, w.err = w.w.Write(buf[:7])
	return w
}

func (w *Writer) I56LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI56LE(buf[:7], s)
	w.n, w.err = w.w.Write(buf[:7])
	return w
}

func (w *Writer) I64BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI64BE(buf[:8], s)
	w.n, w.err = w.w.Write(buf)
	return w
}

func (w *Writer) I64LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.closed {
		w.err = ErrAlreadyClosed
		return w
	}
	buf := *bufBytesPool.Get().(*[]byte)
	defer bufBytesPool.Put(&buf)
	PutI64LE(buf[:8], s)
	w.n, w.err = w.w.Write(buf)
	return w
}
