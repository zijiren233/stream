package stream

import (
	"io"
	"reflect"
)

type Writer struct {
	w     io.Writer
	o     Order
	n     int
	total int
	err   error
	buf   []byte
}

func NewWriter(w io.Writer, o Order) *Writer {
	return &Writer{
		w:   w,
		o:   o,
		buf: make([]byte, 8),
	}
}

func (w *Writer) Error() error {
	return w.err
}

// N returns the number of bytes write at last write.
// Instead of the total.
func (w *Writer) N() int {
	return w.n
}

func (w *Writer) Total() int {
	return w.total
}

func (w *Writer) Reset() *Writer {
	w.n = 0
	w.total = 0
	w.err = nil
	return w
}

func (w *Writer) Write(data any) *Writer {
	if w.err != nil {
		return w
	}

	preTotal := w.total

	switch data := data.(type) {
	case *bool:
		w.Bool(*data)
	case bool:
		w.Bool(data)
	case []bool:
		for i := range data {
			w.Bool(data[i])
		}
	case *int8:
		w.I8(*data)
	case int8:
		w.I8(data)
	case []int8:
		for i := range data {
			w.I8(data[i])
		}
	case *uint8:
		w.Byte(*data)
	case uint8:
		w.Byte(data)
	case []uint8:
		w.Bytes(data)
	case *int16:
		w.I16(*data)
	case int16:
		w.I16(data)
	case []int16:
		for i := range data {
			w.I16(data[i])
		}
	case *uint16:
		w.U16(*data)
	case uint16:
		w.U16(data)
	case []uint16:
		for i := range data {
			w.U16(data[i])
		}
	case *int32:
		w.I32(*data)
	case int32:
		w.I32(data)
	case []int32:
		for i := range data {
			w.I32(data[i])
		}
	case *uint32:
		w.U32(*data)
	case uint32:
		w.U32(data)
	case []uint32:
		for i := range data {
			w.U32(data[i])
		}
	case *int64:
		w.I64(*data)
	case int64:
		w.I64(data)
	case []int64:
		for i := range data {
			w.I64(data[i])
		}
	case *uint64:
		w.U64(*data)
	case uint64:
		w.U64(data)
	case []uint64:
		for i := range data {
			w.U64(data[i])
		}
	case *float32:
		w.F32(*data)
	case float32:
		w.F32(data)
	case []float32:
		for i := range data {
			w.F32(data[i])
		}
	case *float64:
		w.F64(*data)
	case float64:
		w.F64(data)
	case []float64:
		for i := range data {
			w.F64(data[i])
		}
	default:
		v := reflect.Indirect(reflect.ValueOf(data))
		if v.Kind() == reflect.Struct {
			for i := 0; i < v.NumField(); i++ {
				field := reflect.Indirect(v.Field(i))
				if field.CanAddr() {
					if w.Write(field.Addr().Interface()).err != nil {
						break
					}
				} else if field.CanInterface() {
					if w.Write(field.Interface()).err != nil {
						break
					}
				} else {
					w.err = FormatUnsupportedTypeError(field.Type().String())
					break
				}
			}
		} else {
			w.err = FormatUnsupportedTypeError(v.Type().String())
		}
	}

	w.n = w.total - preTotal

	return w
}

func (w *Writer) Byte(s byte) *Writer {
	if w.err != nil {
		return w
	}
	w.n, w.err = w.w.Write([]byte{s})
	w.total += w.n

	return w
}

func (w *Writer) Bytes(s []byte) *Writer {
	if w.err != nil {
		return w
	}
	w.n, w.err = w.w.Write(s)
	w.total += w.n

	return w
}

func (w *Writer) I8(s int8) *Writer {
	if w.err != nil {
		return w
	}
	WriteI8(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) I16(s int16) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI16BE(w.buf[:2], s)
	} else {
		WriteI16LE(w.buf[:2], s)
	}
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) I24(s int32) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI24BE(w.buf[:3], s)
	} else {
		WriteI24LE(w.buf[:3], s)
	}
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) I32(s int32) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI32BE(w.buf[:4], s)
	} else {
		WriteI32LE(w.buf[:4], s)
	}
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) I40(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI40BE(w.buf[:5], s)
	} else {
		WriteI40LE(w.buf[:5], s)
	}
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) I48(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI48BE(w.buf[:6], s)
	} else {
		WriteI48LE(w.buf[:6], s)
	}
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) I56(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI56BE(w.buf[:7], s)
	} else {
		WriteI56LE(w.buf[:7], s)
	}
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) I64(s int64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteI64BE(w.buf[:8], s)
	} else {
		WriteI64LE(w.buf[:8], s)
	}
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) U8(s uint8) *Writer {
	if w.err != nil {
		return w
	}
	WriteU8(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) U16(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU16BE(w.buf[:2], s)
	} else {
		WriteU16LE(w.buf[:2], s)
	}
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) U24(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU24BE(w.buf[:3], s)
	} else {
		WriteU24LE(w.buf[:3], s)
	}
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) U32(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU32BE(w.buf[:4], s)
	} else {
		WriteU32LE(w.buf[:4], s)
	}
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) U40(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU40BE(w.buf[:5], s)
	} else {
		WriteU40LE(w.buf[:5], s)
	}
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) U48(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU48BE(w.buf[:6], s)
	} else {
		WriteU48LE(w.buf[:6], s)
	}
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) U56(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU56BE(w.buf[:7], s)
	} else {
		WriteU56LE(w.buf[:7], s)
	}
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) U64(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteU64BE(w.buf[:8], s)
	} else {
		WriteU64LE(w.buf[:8], s)
	}
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) F32(s float32) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteF32BE(w.buf[:4], s)
	} else {
		WriteF32LE(w.buf[:4], s)
	}
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F64(s float64) *Writer {
	if w.err != nil {
		return w
	}
	if w.o == BigEndian {
		WriteF64BE(w.buf[:8], s)
	} else {
		WriteF64LE(w.buf[:8], s)
	}
	w.n, w.err = w.w.Write(w.buf[:8])
	w.total += w.n

	return w
}

func (w *Writer) Bool(s bool) *Writer {
	if w.err != nil {
		return w
	}
	WriteBool(w.buf[:1], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) String(s string) *Writer {
	if w.err != nil {
		return w
	}

	w.n, w.err = w.w.Write(StringToBytes(s))
	w.total += w.n

	return w
}
