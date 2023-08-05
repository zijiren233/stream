package stream

import (
	"io"
)

type Reader struct {
	r     io.Reader
	n     int
	total int
	err   error
	buf   []byte
}

func NewReader(r io.Reader) *Reader {
	return &Reader{
		r:   r,
		buf: make([]byte, 8),
	}
}

func (r *Reader) Error() error {
	return r.err
}

// N returns the number of bytes read at last read.
// Instead of the total.
func (r *Reader) N() int {
	return r.n
}

func (r *Reader) Total() int {
	return r.total
}

func (r *Reader) Reset() {
	r.n = 0
	r.total = 0
	r.err = nil
}

func (r *Reader) Byte(t *byte) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:1])
	if r.err == nil {
		*t = r.buf[0]
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadByte() (t byte, err error) {
	err = r.Byte(&t).Error()
	r.err = nil

	return
}

func (r *Reader) Bytes(t []byte) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, t)
	r.total += r.n

	return r
}

func (r *Reader) ReadBytes(n int) (t []byte, err error) {
	t = make([]byte, n)
	err = r.Bytes(t).Error()
	r.err = nil

	return
}

func (r *Reader) SkipBytes(n int) *Reader {
	if r.err != nil {
		return r
	}
	var tn int64
	tn, r.err = io.CopyN(io.Discard, r.r, int64(n))
	r.n = int(tn)
	r.total += r.n

	return r
}

func (r *Reader) I8(t *int8) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:1])
	if r.err == nil {
		*t = ReadI8(r.buf[:1])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI8() (t int8, err error) {
	err = r.I8(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I16BE(t *int16) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:2])
	if r.err == nil {
		*t = ReadI16BE(r.buf[:2])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI16BE() (t int16, err error) {
	err = r.I16BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I16LE(t *int16) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:2])
	if r.err == nil {
		*t = ReadI16LE(r.buf[:2])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI16LE() (t int16, err error) {
	err = r.I16LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I24BE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:3])
	if r.err == nil {
		*t = ReadI24BE(r.buf[:3])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI24BE() (t int32, err error) {
	err = r.I24BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I24LE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:3])
	if r.err == nil {
		*t = ReadI24LE(r.buf[:3])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI24LE() (t int32, err error) {
	err = r.I24LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I32BE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		*t = ReadI32BE(r.buf[:4])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI32BE() (t int32, err error) {
	err = r.I32BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I32LE(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		*t = ReadI32LE(r.buf[:4])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI32LE() (t int32, err error) {
	err = r.I32LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I40BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:5])
	if r.err == nil {
		*t = ReadI40BE(r.buf[:5])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI40BE() (t int64, err error) {
	err = r.I40BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I40LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:5])
	if r.err == nil {
		*t = ReadI40LE(r.buf[:5])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI40LE() (t int64, err error) {
	err = r.I40LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I48BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:6])
	if r.err == nil {
		*t = ReadI48BE(r.buf[:6])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI48BE() (t int64, err error) {
	err = r.I48BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I48LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:6])
	if r.err == nil {
		*t = ReadI48LE(r.buf[:6])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI48LE() (t int64, err error) {
	err = r.I48LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I56BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:7])
	if r.err == nil {
		*t = ReadI56BE(r.buf[:7])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI56BE() (t int64, err error) {
	err = r.I56BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I56LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:7])
	if r.err == nil {
		*t = ReadI56LE(r.buf[:7])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI56LE() (t int64, err error) {
	err = r.I56LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I64BE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf)
	if r.err == nil {
		*t = ReadI64BE(r.buf)
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI64BE() (t int64, err error) {
	err = r.I64BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I64LE(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf)
	if r.err == nil {
		*t = ReadI64LE(r.buf)
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI64LE() (t int64, err error) {
	err = r.I64LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U8(t *uint8) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:1])
	if r.err == nil {
		*t = ReadU8(r.buf[:1])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU8() (t uint8, err error) {
	err = r.U8(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U16BE(t *uint16) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:2])
	if r.err == nil {
		*t = ReadU16BE(r.buf[:2])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU16BE() (t uint16, err error) {
	err = r.U16BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U16LE(t *uint16) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:2])
	if r.err == nil {
		*t = ReadU16LE(r.buf[:2])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU16LE() (t uint16, err error) {
	err = r.U16LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U24BE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:3])
	if r.err == nil {
		*t = ReadU24BE(r.buf[:3])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU24BE() (t uint32, err error) {
	err = r.U24BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U24LE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:3])
	if r.err == nil {
		*t = ReadU24LE(r.buf[:3])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU24LE() (t uint32, err error) {
	err = r.U24LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U32BE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		*t = ReadU32BE(r.buf[:4])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU32BE() (t uint32, err error) {
	err = r.U32BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U32LE(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		*t = ReadU32LE(r.buf[:4])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU32LE() (t uint32, err error) {
	err = r.U32LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U40BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:5])
	if r.err == nil {
		*t = ReadU40BE(r.buf[:5])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU40BE() (t uint64, err error) {
	err = r.U40BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U40LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:5])
	if r.err == nil {
		*t = ReadU40LE(r.buf[:5])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU40LE() (t uint64, err error) {
	err = r.U40LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U48BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:6])
	if r.err == nil {
		*t = ReadU48BE(r.buf[:6])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU48BE() (t uint64, err error) {
	err = r.U48BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U48LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:6])
	if r.err == nil {
		*t = ReadU48LE(r.buf[:6])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU48LE() (t uint64, err error) {
	err = r.U48LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U56BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:7])
	if r.err == nil {
		*t = ReadU56BE(r.buf[:7])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU56BE() (t uint64, err error) {
	err = r.U56BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U56LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:7])
	if r.err == nil {
		*t = ReadU56LE(r.buf[:7])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU56LE() (t uint64, err error) {
	err = r.U56LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U64BE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf)
	if r.err == nil {
		*t = ReadU64BE(r.buf)
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU64BE() (t uint64, err error) {
	err = r.U64BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U64LE(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf)
	if r.err == nil {
		*t = ReadU64LE(r.buf)
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU64LE() (t uint64, err error) {
	err = r.U64LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) F32BE(t *float32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		*t = ReadF32BE(r.buf[:4])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadF32BE() (t float32, err error) {
	err = r.F32BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) F32LE(t *float32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		*t = ReadF32LE(r.buf[:4])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadF32LE() (t float32, err error) {
	err = r.F32LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) F64BE(t *float64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:8])
	if r.err == nil {
		*t = ReadF64BE(r.buf[:8])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadF64BE() (t float64, err error) {
	err = r.F64BE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) F64LE(t *float64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:8])
	if r.err == nil {
		*t = ReadF64LE(r.buf[:8])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadF64LE() (t float64, err error) {
	err = r.F64LE(&t).Error()
	r.err = nil

	return
}

func (r *Reader) Bool(t *bool) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = r.r.Read(r.buf[:1])
	if r.err == nil {
		*t = ReadBool(r.buf[:1])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadBool() (t bool, err error) {
	err = r.Bool(&t).Error()
	r.err = nil

	return
}

func (r *Reader) String(t *string, n int) *Reader {
	var b = make([]byte, n)
	r.Bytes(b)
	*t = BytesToString(b)

	return r
}

func (r *Reader) ReadString(n int) (t string, err error) {
	err = r.String(&t, n).Error()
	r.err = nil

	return
}
