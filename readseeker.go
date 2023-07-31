package stream

import "io"

type ReadSeeker struct {
	*Reader
	rs io.ReadSeeker
}

func NewReadSeeker(rs io.ReadSeeker) *ReadSeeker {
	return &ReadSeeker{
		Reader: NewReader(rs),
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

func (r *ReadSeeker) F32BE(t *float32) *ReadSeeker {
	r.Reader.F32BE(t)
	return r
}

func (r *ReadSeeker) F32LE(t *float32) *ReadSeeker {
	r.Reader.F32LE(t)
	return r
}

func (r *ReadSeeker) F64BE(t *float64) *ReadSeeker {
	r.Reader.F64BE(t)
	return r
}

func (r *ReadSeeker) F64LE(t *float64) *ReadSeeker {
	r.Reader.F64LE(t)
	return r
}

func (r *ReadSeeker) I16BE(t *int16) *ReadSeeker {
	r.Reader.I16BE(t)
	return r
}

func (r *ReadSeeker) I16LE(t *int16) *ReadSeeker {
	r.Reader.I16LE(t)
	return r
}

func (r *ReadSeeker) I24BE(t *int32) *ReadSeeker {
	r.Reader.I24BE(t)
	return r
}

func (r *ReadSeeker) I24LE(t *int32) *ReadSeeker {
	r.Reader.I24LE(t)
	return r
}

func (r *ReadSeeker) I32BE(t *int32) *ReadSeeker {
	r.Reader.I32BE(t)
	return r
}

func (r *ReadSeeker) I32LE(t *int32) *ReadSeeker {
	r.Reader.I32LE(t)
	return r
}

func (r *ReadSeeker) I40BE(t *int64) *ReadSeeker {
	r.Reader.I40BE(t)
	return r
}

func (r *ReadSeeker) I40LE(t *int64) *ReadSeeker {
	r.Reader.I40LE(t)
	return r
}

func (r *ReadSeeker) I48BE(t *int64) *ReadSeeker {
	r.Reader.I48BE(t)
	return r
}

func (r *ReadSeeker) I48LE(t *int64) *ReadSeeker {
	r.Reader.I48LE(t)
	return r
}

func (r *ReadSeeker) I56BE(t *int64) *ReadSeeker {
	r.Reader.I56BE(t)
	return r
}

func (r *ReadSeeker) I56LE(t *int64) *ReadSeeker {
	r.Reader.I56LE(t)
	return r
}

func (r *ReadSeeker) I64BE(t *int64) *ReadSeeker {
	r.Reader.I64BE(t)
	return r
}

func (r *ReadSeeker) I64LE(t *int64) *ReadSeeker {
	r.Reader.I64LE(t)
	return r
}

func (r *ReadSeeker) I8(t *int8) *ReadSeeker {
	r.Reader.I8(t)
	return r
}

func (r *ReadSeeker) U16BE(t *uint16) *ReadSeeker {
	r.Reader.U16BE(t)
	return r
}

func (r *ReadSeeker) U16LE(t *uint16) *ReadSeeker {
	r.Reader.U16LE(t)
	return r
}

func (r *ReadSeeker) U24BE(t *uint32) *ReadSeeker {
	r.Reader.U24BE(t)
	return r
}

func (r *ReadSeeker) U24LE(t *uint32) *ReadSeeker {
	r.Reader.U24LE(t)
	return r
}

func (r *ReadSeeker) U32BE(t *uint32) *ReadSeeker {
	r.Reader.U32BE(t)
	return r
}

func (r *ReadSeeker) U32LE(t *uint32) *ReadSeeker {
	r.Reader.U32LE(t)
	return r
}

func (r *ReadSeeker) U40BE(t *uint64) *ReadSeeker {
	r.Reader.U40BE(t)
	return r
}

func (r *ReadSeeker) U40LE(t *uint64) *ReadSeeker {
	r.Reader.U40LE(t)
	return r
}

func (r *ReadSeeker) U48BE(t *uint64) *ReadSeeker {
	r.Reader.U48BE(t)
	return r
}

func (r *ReadSeeker) U48LE(t *uint64) *ReadSeeker {
	r.Reader.U48LE(t)
	return r
}

func (r *ReadSeeker) U56BE(t *uint64) *ReadSeeker {
	r.Reader.U56BE(t)
	return r
}

func (r *ReadSeeker) U56LE(t *uint64) *ReadSeeker {
	r.Reader.U56LE(t)
	return r
}

func (r *ReadSeeker) U64BE(t *uint64) *ReadSeeker {
	r.Reader.U64BE(t)
	return r
}

func (r *ReadSeeker) U64LE(t *uint64) *ReadSeeker {
	r.Reader.U64LE(t)
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
