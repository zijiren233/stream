package stream

func (w *Writer) I8(s int8) *Writer {
	if w.err != nil {
		return w
	}
	PutI8(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) I16BE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	PutI16BE(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) I16LE(s int16) *Writer {
	if w.err != nil {
		return w
	}
	PutI16LE(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) I24BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	PutI24BE(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) I24LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	PutI24LE(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) I32BE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	PutI32BE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) I32LE(s int32) *Writer {
	if w.err != nil {
		return w
	}
	PutI32LE(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) I40BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI40BE(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) I40LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI40LE(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) I48BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI48BE(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) I48LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI48LE(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) I56BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI56BE(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) I56LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI56LE(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) I64BE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI64BE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) I64LE(s int64) *Writer {
	if w.err != nil {
		return w
	}
	PutI64LE(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}
