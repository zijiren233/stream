package stream

import "math"

type Order interface {
	ReadI8(s byte) int8
	// A maximum of 2 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI16(s []byte) int16
	// A maximum of 3 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI24(s []byte) int32
	// A maximum of 4 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI32(s []byte) int32
	// A maximum of 5 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI40(s []byte) int64
	// A maximum of 6 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI48(s []byte) int64
	// A maximum of 7 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI56(s []byte) int64
	// A maximum of 8 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadI64(s []byte) int64
	ReadU8(s byte) uint8
	// A maximum of 2 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU16(s []byte) uint16
	// A maximum of 3 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU24(s []byte) uint32
	// A maximum of 4 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU32(s []byte) uint32
	// A maximum of 5 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU40(s []byte) uint64
	// A maximum of 6 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU48(s []byte) uint64
	// A maximum of 7 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU56(s []byte) uint64
	// A maximum of 8 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadU64(s []byte) uint64
	// A maximum of 4 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadF32(s []byte) float32
	// A maximum of 8 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	ReadF64(s []byte) float64
	ReadBool(s byte) bool

	WriteI8(buf *byte, v int8)
	// A maximum of 2 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI16(buf []byte, v int16)
	// A maximum of 3 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI24(buf []byte, v int32)
	// A maximum of 4 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI32(buf []byte, v int32)
	// A maximum of 5 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI40(buf []byte, v int64)
	// A maximum of 6 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI48(buf []byte, v int64)
	// A maximum of 7 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI56(buf []byte, v int64)
	// A maximum of 8 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteI64(buf []byte, v int64)
	WriteU8(buf *byte, v uint8)
	// A maximum of 2 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU16(buf []byte, v uint16)
	// A maximum of 3 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU24(buf []byte, v uint32)
	// A maximum of 4 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU32(buf []byte, v uint32)
	// A maximum of 5 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU40(buf []byte, v uint64)
	// A maximum of 6 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU48(buf []byte, v uint64)
	// A maximum of 7 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU56(buf []byte, v uint64)
	// A maximum of 8 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteU64(buf []byte, v uint64)
	// A maximum of 4 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteF32(buf []byte, v float32)
	// A maximum of 8 bytes is required. If the size of the incoming slice is not enough, the precision may be lost
	WriteF64(buf []byte, v float64)
	WriteBool(buf *byte, v bool)
}

var (
	BigEndian    Order = (*bigEndian)(nil)
	LittleEndian Order = (*littleEndian)(nil)
)

type bigEndian struct{}

func (b *bigEndian) ReadI8(s byte) int8 {
	return int8(s)
}

func (b *bigEndian) ReadI16(s []byte) int16 {
	var d int16
	for i, v := range s {
		if i > 1 {
			break
		}
		d |= int16(v) << uint(8*(1-i))
	}
	return d
}

func (b *bigEndian) ReadI24(s []byte) int32 {
	var d int32
	for i, v := range s {
		if i > 2 {
			break
		}
		d |= int32(v) << uint(8*(2-i))
	}
	return d
}

func (b *bigEndian) ReadI32(s []byte) int32 {
	var d int32
	for i, v := range s {
		if i > 3 {
			break
		}
		d |= int32(v) << uint(8*(3-i))
	}
	return d
}

func (b *bigEndian) ReadI40(s []byte) int64 {
	var d int64
	for i, v := range s {
		if i > 4 {
			break
		}
		d |= int64(v) << uint(8*(4-i))
	}
	return d
}

func (b *bigEndian) ReadI48(s []byte) int64 {
	var d int64
	for i, v := range s {
		if i > 5 {
			break
		}
		d |= int64(v) << uint(8*(5-i))
	}
	return d
}

func (b *bigEndian) ReadI56(s []byte) int64 {
	var d int64
	for i, v := range s {
		if i > 6 {
			break
		}
		d |= int64(v) << uint(8*(6-i))
	}
	return d
}

func (b *bigEndian) ReadI64(s []byte) int64 {
	var d int64
	for i, v := range s {
		if i > 7 {
			break
		}
		d |= int64(v) << uint(8*(7-i))
	}
	return d
}

func (b *bigEndian) ReadU8(s byte) uint8 {
	return s
}

func (b *bigEndian) ReadU16(s []byte) uint16 {
	var d uint16
	for i, v := range s {
		if i > 1 {
			break
		}
		d |= uint16(v) << uint(8*(1-i))
	}
	return d
}

func (b *bigEndian) ReadU24(s []byte) uint32 {
	var d uint32
	for i, v := range s {
		if i > 2 {
			break
		}
		d |= uint32(v) << uint(8*(2-i))
	}
	return d
}

func (b *bigEndian) ReadU32(s []byte) uint32 {
	var d uint32
	for i, v := range s {
		if i > 3 {
			break
		}
		d |= uint32(v) << uint(8*(3-i))
	}
	return d
}

func (b *bigEndian) ReadU40(s []byte) uint64 {
	var d uint64
	for i, v := range s {
		if i > 4 {
			break
		}
		d |= uint64(v) << uint(8*(4-i))
	}
	return d
}

func (b *bigEndian) ReadU48(s []byte) uint64 {
	var d uint64
	for i, v := range s {
		if i > 5 {
			break
		}
		d |= uint64(v) << uint(8*(5-i))
	}
	return d
}

func (b *bigEndian) ReadU56(s []byte) uint64 {
	var d uint64
	for i, v := range s {
		if i > 6 {
			break
		}
		d |= uint64(v) << uint(8*(6-i))
	}
	return d
}

func (b *bigEndian) ReadU64(s []byte) uint64 {
	var d uint64
	for i, v := range s {
		if i > 7 {
			break
		}
		d |= uint64(v) << uint(8*(7-i))
	}
	return d
}

func (b *bigEndian) ReadF32(s []byte) float32 {
	return math.Float32frombits(b.ReadU32(s))
}

func (b *bigEndian) ReadF64(s []byte) float64 {
	return math.Float64frombits(b.ReadU64(s))
}

func (b *bigEndian) ReadBool(s byte) bool {
	return s != 0
}

func (b *bigEndian) WriteI8(buf *byte, v int8) {
	*buf = byte(v)
}

func (b *bigEndian) WriteI16(buf []byte, v int16) {
	for i := range buf {
		if i > 1 {
			break
		}
		buf[i] = byte(v >> uint(8*(1-i)))
	}
}

func (b *bigEndian) WriteI24(buf []byte, v int32) {
	for i := range buf {
		if i > 2 {
			break
		}
		buf[i] = byte(v >> uint(8*(2-i)))
	}
}

func (b *bigEndian) WriteI32(buf []byte, v int32) {
	for i := range buf {
		if i > 3 {
			break
		}
		buf[i] = byte(v >> uint(8*(3-i)))
	}
}

func (b *bigEndian) WriteI40(buf []byte, v int64) {
	for i := range buf {
		if i > 4 {
			break
		}
		buf[i] = byte(v >> uint(8*(4-i)))
	}
}

func (b *bigEndian) WriteI48(buf []byte, v int64) {
	for i := range buf {
		if i > 5 {
			break
		}
		buf[i] = byte(v >> uint(8*(5-i)))
	}
}

func (b *bigEndian) WriteI56(buf []byte, v int64) {
	for i := range buf {
		if i > 6 {
			break
		}
		buf[i] = byte(v >> uint(8*(6-i)))
	}
}

func (b *bigEndian) WriteI64(buf []byte, v int64) {
	for i := range buf {
		if i > 7 {
			break
		}
		buf[i] = byte(v >> uint(8*(7-i)))
	}
}

func (b *bigEndian) WriteU8(buf *byte, v uint8) {
	*buf = byte(v)
}

func (b *bigEndian) WriteU16(buf []byte, v uint16) {
	for i := range buf {
		if i > 1 {
			break
		}
		buf[i] = byte(v >> uint(8*(1-i)))
	}
}

func (b *bigEndian) WriteU24(buf []byte, v uint32) {
	for i := range buf {
		if i > 2 {
			break
		}
		buf[i] = byte(v >> uint(8*(2-i)))
	}
}

func (b *bigEndian) WriteU32(buf []byte, v uint32) {
	for i := range buf {
		if i > 3 {
			break
		}
		buf[i] = byte(v >> uint(8*(3-i)))
	}
}

func (b *bigEndian) WriteU40(buf []byte, v uint64) {
	for i := range buf {
		if i > 4 {
			break
		}
		buf[i] = byte(v >> uint(8*(4-i)))
	}
}

func (b *bigEndian) WriteU48(buf []byte, v uint64) {
	for i := range buf {
		if i > 5 {
			break
		}
		buf[i] = byte(v >> uint(8*(5-i)))
	}
}

func (b *bigEndian) WriteU56(buf []byte, v uint64) {
	for i := range buf {
		if i > 6 {
			break
		}
		buf[i] = byte(v >> uint(8*(6-i)))
	}
}

func (b *bigEndian) WriteU64(buf []byte, v uint64) {
	for i := range buf {
		if i > 7 {
			break
		}
		buf[i] = byte(v >> uint(8*(7-i)))
	}
}

func (b *bigEndian) WriteF32(buf []byte, v float32) {
	b.WriteU32(buf, math.Float32bits(v))
}

func (b *bigEndian) WriteF64(buf []byte, v float64) {
	b.WriteU64(buf, math.Float64bits(v))
}

func (b *bigEndian) WriteBool(buf *byte, v bool) {
	if v {
		*buf = 1
	} else {
		*buf = 0
	}
}

type littleEndian struct{}

func (l *littleEndian) ReadI8(s byte) int8 {
	return int8(s)
}

func (l *littleEndian) ReadI16(s []byte) int16 {
	var v int16
	for i := range s {
		if i > 1 {
			break
		}
		v |= int16(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadI24(s []byte) int32 {
	var v int32
	for i := range s {
		if i > 2 {
			break
		}
		v |= int32(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadI32(s []byte) int32 {
	var v int32
	for i := range s {
		if i > 3 {
			break
		}
		v |= int32(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadI40(s []byte) int64 {
	var v int64
	for i := range s {
		if i > 4 {
			break
		}
		v |= int64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadI48(s []byte) int64 {
	var v int64
	for i := range s {
		if i > 5 {
			break
		}
		v |= int64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadI56(s []byte) int64 {
	var v int64
	for i := range s {
		if i > 6 {
			break
		}
		v |= int64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadI64(s []byte) int64 {
	var v int64
	for i := range s {
		if i > 7 {
			break
		}
		v |= int64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU8(s byte) uint8 {
	return uint8(s)
}

func (l *littleEndian) ReadU16(s []byte) uint16 {
	var v uint16
	for i := range s {
		if i > 1 {
			break
		}
		v |= uint16(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU24(s []byte) uint32 {
	var v uint32
	for i := range s {
		if i > 2 {
			break
		}
		v |= uint32(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU32(s []byte) uint32 {
	var v uint32
	for i := range s {
		if i > 3 {
			break
		}
		v |= uint32(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU40(s []byte) uint64 {
	var v uint64
	for i := range s {
		if i > 4 {
			break
		}
		v |= uint64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU48(s []byte) uint64 {
	var v uint64
	for i := range s {
		if i > 5 {
			break
		}
		v |= uint64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU56(s []byte) uint64 {
	var v uint64
	for i := range s {
		if i > 6 {
			break
		}
		v |= uint64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadU64(s []byte) uint64 {
	var v uint64
	for i := range s {
		if i > 7 {
			break
		}
		v |= uint64(s[i]) << uint(8*i)
	}
	return v
}

func (l *littleEndian) ReadF32(s []byte) float32 {
	return math.Float32frombits(l.ReadU32(s))
}

func (l *littleEndian) ReadF64(s []byte) float64 {
	return math.Float64frombits(l.ReadU64(s))
}

func (l *littleEndian) ReadBool(s byte) bool {
	return s != 0
}

func (l *littleEndian) WriteI8(buf *byte, v int8) {
	*buf = byte(v)
}

func (l *littleEndian) WriteI16(buf []byte, v int16) {
	for i := range buf {
		if i > 1 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteI24(buf []byte, v int32) {
	for i := range buf {
		if i > 2 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteI32(buf []byte, v int32) {
	for i := range buf {
		if i > 3 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteI40(buf []byte, v int64) {
	for i := range buf {
		if i > 4 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteI48(buf []byte, v int64) {
	for i := range buf {
		if i > 5 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteI56(buf []byte, v int64) {
	for i := range buf {
		if i > 6 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteI64(buf []byte, v int64) {
	for i := range buf {
		if i > 7 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU8(buf *byte, v uint8) {
	*buf = v
}

func (l *littleEndian) WriteU16(buf []byte, v uint16) {
	for i := range buf {
		if i > 1 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU24(buf []byte, v uint32) {
	for i := range buf {
		if i > 2 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU32(buf []byte, v uint32) {
	for i := range buf {
		if i > 3 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU40(buf []byte, v uint64) {
	for i := range buf {
		if i > 4 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU48(buf []byte, v uint64) {
	for i := range buf {
		if i > 5 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU56(buf []byte, v uint64) {
	for i := range buf {
		if i > 6 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteU64(buf []byte, v uint64) {
	for i := range buf {
		if i > 7 {
			break
		}
		buf[i] = byte(v >> uint(8*i))
	}
}

func (l *littleEndian) WriteF32(buf []byte, v float32) {
	l.WriteU32(buf, math.Float32bits(v))
}

func (l *littleEndian) WriteF64(buf []byte, v float64) {
	l.WriteU64(buf, math.Float64bits(v))
}

func (l *littleEndian) WriteBool(buf *byte, v bool) {
	if v {
		*buf = 1
	} else {
		*buf = 0
	}
}
