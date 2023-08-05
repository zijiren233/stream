package stream

import "io"

type ReadSeeker struct {
	*Reader
	rs io.ReadSeeker
}

func NewReadSeeker(rs io.ReadSeeker, o Order) *ReadSeeker {
	return &ReadSeeker{
		Reader: NewReader(rs, o),
		rs:     rs,
	}
}

func (r *ReadSeeker) Seek(offset int64, whence int) *ReadSeeker {
	if r.err != nil {
		return r
	}
	_, r.err = r.rs.Seek(offset, whence)
	return r
}

func (r *ReadSeeker) Read(data any) *ReadSeeker {
	r.Reader.Read(data)
	return r
}

func (r *ReadSeeker) Bool(t *bool) *ReadSeeker {
	r.Reader.Bool(t)
	return r
}

func (r *ReadSeeker) Byte(t *byte) *ReadSeeker {
	r.Reader.Byte(t)
	return r
}

func (r *ReadSeeker) Bytes(t []byte) *ReadSeeker {
	r.Reader.Bytes(t)
	return r
}

func (r *ReadSeeker) SkipN(n int) *ReadSeeker {
	r.Reader.SkipN(n)
	return r
}

func (r *ReadSeeker) F32(t *float32) *ReadSeeker {
	r.Reader.F32(t)
	return r
}

func (r *ReadSeeker) F64(t *float64) *ReadSeeker {
	r.Reader.F64(t)
	return r
}

func (r *ReadSeeker) I16(t *int16) *ReadSeeker {
	r.Reader.I16(t)
	return r
}

func (r *ReadSeeker) I24(t *int32) *ReadSeeker {
	r.Reader.I24(t)
	return r
}

func (r *ReadSeeker) I32(t *int32) *ReadSeeker {
	r.Reader.I32(t)
	return r
}

func (r *ReadSeeker) I40(t *int64) *ReadSeeker {
	r.Reader.I40(t)
	return r
}

func (r *ReadSeeker) I48(t *int64) *ReadSeeker {
	r.Reader.I48(t)
	return r
}

func (r *ReadSeeker) I56(t *int64) *ReadSeeker {
	r.Reader.I56(t)
	return r
}

func (r *ReadSeeker) I64(t *int64) *ReadSeeker {
	r.Reader.I64(t)
	return r
}

func (r *ReadSeeker) I8(t *int8) *ReadSeeker {
	r.Reader.I8(t)
	return r
}

func (r *ReadSeeker) U16(t *uint16) *ReadSeeker {
	r.Reader.U16(t)
	return r
}

func (r *ReadSeeker) U24(t *uint32) *ReadSeeker {
	r.Reader.U24(t)
	return r
}

func (r *ReadSeeker) U32(t *uint32) *ReadSeeker {
	r.Reader.U32(t)
	return r
}

func (r *ReadSeeker) U40(t *uint64) *ReadSeeker {
	r.Reader.U40(t)
	return r
}

func (r *ReadSeeker) U48(t *uint64) *ReadSeeker {
	r.Reader.U48(t)
	return r
}

func (r *ReadSeeker) U56(t *uint64) *ReadSeeker {
	r.Reader.U56(t)
	return r
}

func (r *ReadSeeker) U64(t *uint64) *ReadSeeker {
	r.Reader.U64(t)
	return r
}

func (r *ReadSeeker) U8(t *uint8) *ReadSeeker {
	r.Reader.U8(t)
	return r
}

func (r *ReadSeeker) String(t *string, n int) *ReadSeeker {
	r.Reader.String(t, n)
	return r
}
