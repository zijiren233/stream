package stream

func (w *Writer) U8(s uint8) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU8(buf[:1], s)
	n, err := w.w.Write(buf[:1])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U16BE(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU16BE(buf[:2], s)
	n, err := w.w.Write(buf[:2])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U16LE(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU16LE(buf[:2], s)
	n, err := w.w.Write(buf[:2])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U24BE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU24BE(buf[:3], s)
	n, err := w.w.Write(buf[:3])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U24LE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU24LE(buf[:3], s)
	n, err := w.w.Write(buf[:3])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U32BE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU32BE(buf[:4], s)
	n, err := w.w.Write(buf[:4])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U32LE(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU32LE(buf[:4], s)
	n, err := w.w.Write(buf[:4])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U40BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU40BE(buf[:5], s)
	n, err := w.w.Write(buf[:5])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U40LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU40LE(buf[:5], s)
	n, err := w.w.Write(buf[:5])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U48BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU48BE(buf[:6], s)
	n, err := w.w.Write(buf[:6])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U48LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU48LE(buf[:6], s)
	n, err := w.w.Write(buf[:6])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U56BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU56BE(buf[:7], s)
	n, err := w.w.Write(buf[:7])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U56LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU56LE(buf[:7], s)
	n, err := w.w.Write(buf[:7])
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U64BE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU64BE(buf[:8], s)
	n, err := w.w.Write(buf)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}

func (w *Writer) U64LE(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	buf := *bufPool.Get().(*[]byte)
	defer bufPool.Put(&buf)
	PutU64LE(buf[:8], s)
	n, err := w.w.Write(buf)
	if err != nil {
		w.err = err
	}
	w.n = n
	return w
}
