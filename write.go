package stream

import "math"

func WriteI8(buf *byte, v int8) {
	*buf = byte(v)
}

func WriteI16BE(buf []byte, v int16) error {
	if len(buf) != 2 {
		return FormatLengthError(2)
	}
	writeI16BE(buf, v)
	return nil
}

func writeI16BE(buf []byte, v int16) {
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func WriteI16LE(buf []byte, v int16) error {
	if len(buf) != 2 {
		return FormatLengthError(2)
	}
	writeI16LE(buf, v)
	return nil
}

func writeI16LE(buf []byte, v int16) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
}

func WriteI24BE(buf []byte, v int32) error {
	if len(buf) != 3 {
		return FormatLengthError(3)
	}
	writeI24BE(buf, v)
	return nil
}

func writeI24BE(buf []byte, v int32) {
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
}

func WriteI24LE(buf []byte, v int32) error {
	if len(buf) != 3 {
		return FormatLengthError(3)
	}
	writeI24LE(buf, v)
	return nil
}

func writeI24LE(buf []byte, v int32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
}

func WriteI32BE(buf []byte, v int32) error {
	if len(buf) != 4 {
		return FormatLengthError(4)
	}
	writeI32BE(buf, v)
	return nil
}

func writeI32BE(buf []byte, v int32) {
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func WriteI32LE(buf []byte, v int32) error {
	if len(buf) != 4 {
		return FormatLengthError(4)
	}
	writeI32LE(buf, v)
	return nil
}

func writeI32LE(buf []byte, v int32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

func WriteI40BE(buf []byte, v int64) error {
	if len(buf) != 5 {
		return FormatLengthError(5)
	}
	writeI40BE(buf, v)
	return nil
}

func writeI40BE(buf []byte, v int64) {
	buf[0] = byte(v >> 32)
	buf[1] = byte(v >> 24)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 8)
	buf[4] = byte(v)
}

func WriteI40LE(buf []byte, v int64) error {
	if len(buf) != 5 {
		return FormatLengthError(5)
	}
	writeI40LE(buf, v)
	return nil
}

func writeI40LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
}

func WriteI48BE(buf []byte, v int64) error {
	if len(buf) != 6 {
		return FormatLengthError(6)
	}
	writeI48BE(buf, v)
	return nil
}

func writeI48BE(buf []byte, v int64) {
	buf[0] = byte(v >> 40)
	buf[1] = byte(v >> 32)
	buf[2] = byte(v >> 24)
	buf[3] = byte(v >> 16)
	buf[4] = byte(v >> 8)
	buf[5] = byte(v)
}

func WriteI48LE(buf []byte, v int64) error {
	if len(buf) != 6 {
		return FormatLengthError(6)
	}
	writeI48LE(buf, v)
	return nil
}

func writeI48LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
}

func WriteI56BE(buf []byte, v int64) error {
	if len(buf) != 7 {
		return FormatLengthError(7)
	}
	writeI56BE(buf, v)
	return nil
}

func writeI56BE(buf []byte, v int64) {
	buf[0] = byte(v >> 48)
	buf[1] = byte(v >> 40)
	buf[2] = byte(v >> 32)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 16)
	buf[5] = byte(v >> 8)
	buf[6] = byte(v)
}

func WriteI56LE(buf []byte, v int64) error {
	if len(buf) != 7 {
		return FormatLengthError(7)
	}
	writeI56LE(buf, v)
	return nil
}

func writeI56LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
}

func WriteI64BE(buf []byte, v int64) error {
	if len(buf) != 8 {
		return FormatLengthError(8)
	}
	writeI64BE(buf, v)
	return nil
}

func writeI64BE(buf []byte, v int64) {
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
}

func WriteI64LE(buf []byte, v int64) error {
	if len(buf) != 8 {
		return FormatLengthError(8)
	}
	writeI64LE(buf, v)
	return nil
}

func writeI64LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
}

func WriteU8(buf *byte, v uint8) {
	*buf = byte(v)
}

func WriteU16BE(buf []byte, v uint16) error {
	if len(buf) != 2 {
		return FormatLengthError(2)
	}
	writeU16BE(buf, v)
	return nil
}

func writeU16BE(buf []byte, v uint16) {
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func WriteU16LE(buf []byte, v uint16) error {
	if len(buf) != 2 {
		return FormatLengthError(2)
	}
	writeU16LE(buf, v)
	return nil
}

func writeU16LE(buf []byte, v uint16) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
}

func WriteU24BE(buf []byte, v uint32) error {
	if len(buf) != 3 {
		return FormatLengthError(3)
	}
	writeU24BE(buf, v)
	return nil
}

func writeU24BE(buf []byte, v uint32) {
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
}

func WriteU24LE(buf []byte, v uint32) error {
	if len(buf) != 3 {
		return FormatLengthError(3)
	}
	writeU24LE(buf, v)
	return nil
}

func writeU24LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
}

func WriteU32BE(buf []byte, v uint32) error {
	if len(buf) != 4 {
		return FormatLengthError(4)
	}
	writeU32BE(buf, v)
	return nil
}

func writeU32BE(buf []byte, v uint32) {
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func WriteU32LE(buf []byte, v uint32) error {
	if len(buf) != 4 {
		return FormatLengthError(4)
	}
	writeU32LE(buf, v)
	return nil
}

func writeU32LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

func WriteU40BE(buf []byte, v uint64) error {
	if len(buf) != 5 {
		return FormatLengthError(5)
	}
	writeU40BE(buf, v)
	return nil
}

func writeU40BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 32)
	buf[1] = byte(v >> 24)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 8)
	buf[4] = byte(v)
}

func WriteU40LE(buf []byte, v uint64) error {
	if len(buf) != 5 {
		return FormatLengthError(5)
	}
	writeU40LE(buf, v)
	return nil
}

func writeU40LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
}

func WriteU48BE(buf []byte, v uint64) error {
	if len(buf) != 6 {
		return FormatLengthError(6)
	}
	writeU48BE(buf, v)
	return nil
}

func writeU48BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 40)
	buf[1] = byte(v >> 32)
	buf[2] = byte(v >> 24)
	buf[3] = byte(v >> 16)
	buf[4] = byte(v >> 8)
	buf[5] = byte(v)
}

func WriteU48LE(buf []byte, v uint64) error {
	if len(buf) != 6 {
		return FormatLengthError(6)
	}
	writeU48LE(buf, v)
	return nil
}

func writeU48LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
}

func WriteU56BE(buf []byte, v uint64) error {
	if len(buf) != 7 {
		return FormatLengthError(7)
	}
	writeU56BE(buf, v)
	return nil
}

func writeU56BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 48)
	buf[1] = byte(v >> 40)
	buf[2] = byte(v >> 32)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 16)
	buf[5] = byte(v >> 8)
	buf[6] = byte(v)
}

func WriteU56LE(buf []byte, v uint64) error {
	if len(buf) != 7 {
		return FormatLengthError(7)
	}
	writeU56LE(buf, v)
	return nil
}

func writeU56LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
}

func WriteU64BE(buf []byte, v uint64) error {
	if len(buf) != 8 {
		return FormatLengthError(8)
	}
	writeU64BE(buf, v)
	return nil
}

func writeU64BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
}

func WriteU64LE(buf []byte, v uint64) error {
	if len(buf) != 8 {
		return FormatLengthError(8)
	}
	writeU64LE(buf, v)
	return nil
}

func writeU64LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
}

func WriteF32Be(buf []byte, v float32) error {
	if len(buf) != 4 {
		return FormatLengthError(4)
	}
	writeF32BE(buf, v)
	return nil
}

func writeF32BE(buf []byte, v float32) {
	writeU32BE(buf, math.Float32bits(v))
}

func WriteF32LE(buf []byte, v float32) error {
	if len(buf) != 4 {
		return FormatLengthError(4)
	}
	writeF32LE(buf, v)
	return nil
}

func writeF32LE(buf []byte, v float32) {
	writeU32LE(buf, math.Float32bits(v))
}

func WriteF64BE(buf []byte, v float64) error {
	if len(buf) != 8 {
		return FormatLengthError(8)
	}
	writeF64BE(buf, v)
	return nil
}

func writeF64BE(buf []byte, v float64) {
	writeU64BE(buf, math.Float64bits(v))
}

func WriteF64LE(buf []byte, v float64) error {
	if len(buf) != 8 {
		return FormatLengthError(8)
	}
	writeF64LE(buf, v)
	return nil
}

func writeF64LE(buf []byte, v float64) {
	writeU64LE(buf, math.Float64bits(v))
}

func WriteBool(buf *byte, v bool) {
	if v {
		*buf = 1
	} else {
		*buf = 0
	}
}
