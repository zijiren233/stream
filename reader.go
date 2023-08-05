package stream

import (
	"io"
	"reflect"
)

type Reader struct {
	r     io.Reader
	o     Order
	n     int
	total int
	err   error
	buf   []byte
}

func NewReader(r io.Reader, o Order) *Reader {
	return &Reader{
		r:   r,
		o:   o,
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

func (r *Reader) Read(data any) *Reader {
	if r.err != nil {
		return r
	}
	preTotal := r.total

	switch data := data.(type) {
	case *bool:
		r.Bool(data)
	case []bool:
		for i := range data {
			r.Bool(&data[i])
		}
	case *int8:
		r.I8(data)
	case []int8:
		for i := range data {
			r.I8(&data[i])
		}
	case *uint8:
		r.U8(data)
	case []uint8:
		for i := range data {
			r.U8(&data[i])
		}
	case *int16:
		r.I16(data)
	case []int16:
		for i := range data {
			r.I16(&data[i])
		}
	case *uint16:
		r.U16(data)
	case []uint16:
		for i := range data {
			r.U16(&data[i])
		}
	case *int32:
		r.I32(data)
	case []int32:
		for i := range data {
			r.I32(&data[i])
		}
	case *uint32:
		r.U32(data)
	case []uint32:
		for i := range data {
			r.U32(&data[i])
		}
	case *int64:
		r.I64(data)
	case []int64:
		for i := range data {
			r.I64(&data[i])
		}
	case *uint64:
		r.U64(data)
	case []uint64:
		for i := range data {
			r.U64(&data[i])
		}
	case *float32:
		r.F32(data)
	case []float32:
		for i := range data {
			r.F32(&data[i])
		}
	case *float64:
		r.F64(data)
	case []float64:
		for i := range data {
			r.F64(&data[i])
		}
	default:
		if v, ok := data.([]byte); ok {
			r.Bytes(v)
			break
		}
		v := reflect.Indirect(reflect.ValueOf(data))
		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				field := reflect.Indirect(v.Field(i))
				if field.CanAddr() && field.CanSet() {
					if r.Read(field.Addr().Interface()).err != nil {
						break
					}
				} else {
					r.err = FormatUnsupportedTypeError(field.Type().String())
					break
				}
			}
		} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
			for i := 0; i < v.Len(); i++ {
				elem := reflect.Indirect(v.Index(i))
				if elem.CanAddr() && elem.CanSet() {
					if r.Read(elem.Addr().Interface()).err != nil {
						break
					}
				} else {
					r.err = FormatUnsupportedTypeError(elem.Type().String())
					break
				}
			}
		} else {
			r.err = FormatUnsupportedTypeError(v.Type().String())
		}
	}

	r.n = r.total - preTotal

	return r
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

func (r *Reader) SkipN(n int) *Reader {
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
		*t = ReadI8(r.buf[0])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI8() (t int8, err error) {
	err = r.I8(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I16(t *int16) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:2])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI16BE(r.buf[:2])
		} else {
			*t = readI16LE(r.buf[:2])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI16() (t int16, err error) {
	err = r.I16(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I24(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:3])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI24BE(r.buf[:3])
		} else {
			*t = readI24LE(r.buf[:3])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI24() (t int32, err error) {
	err = r.I24(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I32(t *int32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI32BE(r.buf[:4])
		} else {
			*t = readI32LE(r.buf[:4])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI32() (t int32, err error) {
	err = r.I32(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I40(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:5])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI40BE(r.buf[:5])
		} else {
			*t = readI40LE(r.buf[:5])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI40() (t int64, err error) {
	err = r.I40(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I48(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:6])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI48BE(r.buf[:6])
		} else {
			*t = readI48LE(r.buf[:6])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI48() (t int64, err error) {
	err = r.I48(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I56(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:7])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI56BE(r.buf[:7])
		} else {
			*t = readI56LE(r.buf[:7])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI56() (t int64, err error) {
	err = r.I56(&t).Error()
	r.err = nil

	return
}

func (r *Reader) I64(t *int64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf)
	if r.err == nil {
		if r.o == BigEndian {
			*t = readI64BE(r.buf)
		} else {
			*t = readI64LE(r.buf)
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadI64() (t int64, err error) {
	err = r.I64(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U8(t *uint8) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:1])
	if r.err == nil {
		*t = ReadU8(r.buf[0])
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU8() (t uint8, err error) {
	err = r.U8(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U16(t *uint16) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:2])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU16BE(r.buf[:2])
		} else {
			*t = readU16LE(r.buf[:2])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU16() (t uint16, err error) {
	err = r.U16(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U24(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:3])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU24BE(r.buf[:3])
		} else {
			*t = readU24LE(r.buf[:3])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU24() (t uint32, err error) {
	err = r.U24(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U32(t *uint32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU32BE(r.buf[:4])
		} else {
			*t = readU32LE(r.buf[:4])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU32() (t uint32, err error) {
	err = r.U32(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U40(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:5])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU40BE(r.buf[:5])
		} else {
			*t = readU40LE(r.buf[:5])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU40() (t uint64, err error) {
	err = r.U40(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U48(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:6])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU48BE(r.buf[:6])
		} else {
			*t = readU48LE(r.buf[:6])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU48() (t uint64, err error) {
	err = r.U48(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U56(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:7])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU56BE(r.buf[:7])
		} else {
			*t = readU56LE(r.buf[:7])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU56() (t uint64, err error) {
	err = r.U56(&t).Error()
	r.err = nil

	return
}

func (r *Reader) U64(t *uint64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf)
	if r.err == nil {
		if r.o == BigEndian {
			*t = readU64BE(r.buf)
		} else {
			*t = readU64LE(r.buf)
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadU64() (t uint64, err error) {
	err = r.U64(&t).Error()
	r.err = nil

	return
}

func (r *Reader) F32(t *float32) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:4])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readF32BE(r.buf[:4])
		} else {
			*t = readF32LE(r.buf[:4])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadF32() (t float32, err error) {
	err = r.F32(&t).Error()
	r.err = nil

	return
}

func (r *Reader) F64(t *float64) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = io.ReadFull(r.r, r.buf[:8])
	if r.err == nil {
		if r.o == BigEndian {
			*t = readF64BE(r.buf[:8])
		} else {
			*t = readF64LE(r.buf[:8])
		}
	}
	r.total += r.n

	return r
}

func (r *Reader) ReadF64() (t float64, err error) {
	err = r.F64(&t).Error()
	r.err = nil

	return
}

func (r *Reader) Bool(t *bool) *Reader {
	if r.err != nil {
		return r
	}
	r.n, r.err = r.r.Read(r.buf[:1])
	if r.err == nil {
		*t = ReadBool(r.buf[0])
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
