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

	switch dv := data.(type) {
	case *bool:
		w.Bool(*dv)
	case bool:
		w.Bool(dv)
	case []bool:
		for i := range dv {
			w.Bool(dv[i])
		}
	case *int8:
		w.I8(*dv)
	case int8:
		w.I8(dv)
	case []int8:
		for i := range dv {
			w.I8(dv[i])
		}
	case *uint8:
		w.Byte(*dv)
	case uint8:
		w.Byte(dv)
	case []uint8:
		w.Bytes(dv)
	case *int16:
		w.I16(*dv)
	case int16:
		w.I16(dv)
	case []int16:
		for i := range dv {
			w.I16(dv[i])
		}
	case *uint16:
		w.U16(*dv)
	case uint16:
		w.U16(dv)
	case []uint16:
		for i := range dv {
			w.U16(dv[i])
		}
	case *int32:
		w.I32(*dv)
	case int32:
		w.I32(dv)
	case []int32:
		for i := range dv {
			w.I32(dv[i])
		}
	case *uint32:
		w.U32(*dv)
	case uint32:
		w.U32(dv)
	case []uint32:
		for i := range dv {
			w.U32(dv[i])
		}
	case *int64:
		w.I64(*dv)
	case int64:
		w.I64(dv)
	case []int64:
		for i := range dv {
			w.I64(dv[i])
		}
	case *uint64:
		w.U64(*dv)
	case uint64:
		w.U64(dv)
	case []uint64:
		for i := range dv {
			w.U64(dv[i])
		}
	case *float32:
		w.F32(*dv)
	case float32:
		w.F32(dv)
	case []float32:
		for i := range dv {
			w.F32(dv[i])
		}
	case *float64:
		w.F64(*dv)
	case float64:
		w.F64(dv)
	case []float64:
		for i := range dv {
			w.F64(dv[i])
		}
	default:
		if v, ok := data.([]byte); ok {
			w.Bytes(v)
			break
		}
		v := reflect.Indirect(reflect.ValueOf(dv))
		switch v.Kind() {
		case reflect.Bool:
			w.Bool(v.Bool())
		}
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
		} else if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
			for i := 0; i < v.Len(); i++ {
				elem := reflect.Indirect(v.Index(i))
				if elem.CanInterface() {
					if w.Write(elem.Interface()).err != nil {
						break
					}
				} else {
					w.err = FormatUnsupportedTypeError(elem.Type().String())
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
	w.o.WriteI8(&w.buf[0], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) I16(s int16) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI16(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) I24(s int32) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI24(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) I32(s int32) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI32(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) I40(s int64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI40(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) I48(s int64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI48(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) I56(s int64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI56(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) I64(s int64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteI64(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) U8(s uint8) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU8(&w.buf[0], s)
	w.n, w.err = w.w.Write(w.buf[:1])
	w.total += w.n

	return w
}

func (w *Writer) U16(s uint16) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU16(w.buf[:2], s)
	w.n, w.err = w.w.Write(w.buf[:2])
	w.total += w.n

	return w
}

func (w *Writer) U24(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU24(w.buf[:3], s)
	w.n, w.err = w.w.Write(w.buf[:3])
	w.total += w.n

	return w
}

func (w *Writer) U32(s uint32) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU32(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) U40(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU40(w.buf[:5], s)
	w.n, w.err = w.w.Write(w.buf[:5])
	w.total += w.n

	return w
}

func (w *Writer) U48(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU48(w.buf[:6], s)
	w.n, w.err = w.w.Write(w.buf[:6])
	w.total += w.n

	return w
}

func (w *Writer) U56(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU56(w.buf[:7], s)
	w.n, w.err = w.w.Write(w.buf[:7])
	w.total += w.n

	return w
}

func (w *Writer) U64(s uint64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteU64(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf)
	w.total += w.n

	return w
}

func (w *Writer) F32(s float32) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteF32(w.buf[:4], s)
	w.n, w.err = w.w.Write(w.buf[:4])
	w.total += w.n

	return w
}

func (w *Writer) F64(s float64) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteF64(w.buf[:8], s)
	w.n, w.err = w.w.Write(w.buf[:8])
	w.total += w.n

	return w
}

func (w *Writer) Bool(s bool) *Writer {
	if w.err != nil {
		return w
	}
	w.o.WriteBool(&w.buf[0], s)
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
