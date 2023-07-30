package stream

import "math"

func WriteI8(buf []byte, v int8) {
	buf[0] = byte(v)
}

func WriteI16BE(buf []byte, v int16) {
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func WriteI16LE(buf []byte, v int16) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
}

func WriteI24BE(buf []byte, v int32) {
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
}

func WriteI24LE(buf []byte, v int32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
}

func WriteI32BE(buf []byte, v int32) {
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func WriteI32LE(buf []byte, v int32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

func WriteI40BE(buf []byte, v int64) {
	buf[0] = byte(v >> 32)
	buf[1] = byte(v >> 24)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 8)
	buf[4] = byte(v)
}

func WriteI40LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
}

func WriteI48BE(buf []byte, v int64) {
	buf[0] = byte(v >> 40)
	buf[1] = byte(v >> 32)
	buf[2] = byte(v >> 24)
	buf[3] = byte(v >> 16)
	buf[4] = byte(v >> 8)
	buf[5] = byte(v)
}

func WriteI48LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
}

func WriteI56BE(buf []byte, v int64) {
	buf[0] = byte(v >> 48)
	buf[1] = byte(v >> 40)
	buf[2] = byte(v >> 32)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 16)
	buf[5] = byte(v >> 8)
	buf[6] = byte(v)
}

func WriteI56LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
}

func WriteI64BE(buf []byte, v int64) {
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
}

func WriteI64LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
}

func WriteU8(buf []byte, v uint8) {
	buf[0] = byte(v)
}

func WriteU16BE(buf []byte, v uint16) {
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func WriteU16LE(buf []byte, v uint16) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
}

func WriteU24BE(buf []byte, v uint32) {
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
}

func WriteU24LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
}

func WriteU32BE(buf []byte, v uint32) {
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func WriteU32LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

func WriteU40BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 32)
	buf[1] = byte(v >> 24)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 8)
	buf[4] = byte(v)
}

func WriteU40LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
}

func WriteU48BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 40)
	buf[1] = byte(v >> 32)
	buf[2] = byte(v >> 24)
	buf[3] = byte(v >> 16)
	buf[4] = byte(v >> 8)
	buf[5] = byte(v)
}

func WriteU48LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
}

func WriteU56BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 48)
	buf[1] = byte(v >> 40)
	buf[2] = byte(v >> 32)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 16)
	buf[5] = byte(v >> 8)
	buf[6] = byte(v)
}

func WriteU56LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
}

func WriteU64BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
}

func WriteU64LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
}

func WriteF32BE(buf []byte, v float32) {
	WriteU32BE(buf, math.Float32bits(v))
}

func WriteF32LE(buf []byte, v float32) {
	WriteU32LE(buf, math.Float32bits(v))
}

func WriteF64BE(buf []byte, v float64) {
	WriteU64BE(buf, math.Float64bits(v))
}

func WriteF64LE(buf []byte, v float64) {
	WriteU64LE(buf, math.Float64bits(v))
}

func WriteBool(buf []byte, v bool) {
	if v {
		buf[0] = 1
	} else {
		buf[0] = 0
	}
}
