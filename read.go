package stream

import "math"

func ReadI8(s byte) int8 {
	return int8(s)
}

func ReadI16BE(s []byte) (int16, error) {
	if len(s) != 2 {
		return 0, FormatLengthError(2)
	}
	return readI16BE(s), nil
}

func readI16BE(s []byte) int16 {
	return int16(s[0])<<8 | int16(s[1])
}

func ReadI16LE(s []byte) (int16, error) {
	if len(s) != 2 {
		return 0, FormatLengthError(2)
	}
	return readI16LE(s), nil
}

func readI16LE(s []byte) int16 {
	return int16(s[1])<<8 | int16(s[0])
}

func ReadI24BE(s []byte) (int32, error) {
	if len(s) != 3 {
		return 0, FormatLengthError(3)
	}
	return readI24BE(s), nil
}

func readI24BE(s []byte) int32 {
	return int32(s[0])<<16 | int32(s[1])<<8 | int32(s[2])
}

func ReadI24LE(s []byte) (int32, error) {
	if len(s) != 3 {
		return 0, FormatLengthError(3)
	}
	return readI24LE(s), nil
}

func readI24LE(s []byte) int32 {
	return int32(s[2])<<16 | int32(s[1])<<8 | int32(s[0])
}

func ReadI32BE(s []byte) (int32, error) {
	if len(s) != 4 {
		return 0, FormatLengthError(4)
	}
	return readI32BE(s), nil
}

func readI32BE(s []byte) int32 {
	return int32(s[0])<<24 | int32(s[1])<<16 | int32(s[2])<<8 | int32(s[3])
}

func ReadI32LE(s []byte) (int32, error) {
	if len(s) != 4 {
		return 0, FormatLengthError(4)
	}
	return readI32LE(s), nil
}

func readI32LE(s []byte) int32 {
	return int32(s[3])<<24 | int32(s[2])<<16 | int32(s[1])<<8 | int32(s[0])
}

func ReadI40BE(s []byte) (int64, error) {
	if len(s) != 5 {
		return 0, FormatLengthError(5)
	}
	return readI40BE(s), nil
}

func readI40BE(s []byte) int64 {
	return int64(s[0])<<32 | int64(s[1])<<24 | int64(s[2])<<16 | int64(s[3])<<8 | int64(s[4])
}

func ReadI40LE(s []byte) (int64, error) {
	if len(s) != 5 {
		return 0, FormatLengthError(5)
	}
	return readI40LE(s), nil
}

func readI40LE(s []byte) int64 {
	return int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadI48BE(s []byte) (int64, error) {
	if len(s) != 6 {
		return 0, FormatLengthError(6)
	}
	return readI48BE(s), nil
}

func readI48BE(s []byte) int64 {
	return int64(s[0])<<40 | int64(s[1])<<32 | int64(s[2])<<24 | int64(s[3])<<16 | int64(s[4])<<8 | int64(s[5])
}

func ReadI48LE(s []byte) (int64, error) {
	if len(s) != 6 {
		return 0, FormatLengthError(6)
	}
	return readI48LE(s), nil
}

func readI48LE(s []byte) int64 {
	return int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadI56BE(s []byte) (int64, error) {
	if len(s) != 7 {
		return 0, FormatLengthError(7)
	}
	return readI56BE(s), nil
}

func readI56BE(s []byte) int64 {
	return int64(s[0])<<48 | int64(s[1])<<40 | int64(s[2])<<32 | int64(s[3])<<24 | int64(s[4])<<16 | int64(s[5])<<8 | int64(s[6])
}

func ReadI56LE(s []byte) (int64, error) {
	if len(s) != 7 {
		return 0, FormatLengthError(7)
	}
	return readI56LE(s), nil
}

func readI56LE(s []byte) int64 {
	return int64(s[6])<<48 | int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadI64BE(s []byte) (int64, error) {
	if len(s) != 8 {
		return 0, FormatLengthError(8)
	}
	return readI64BE(s), nil
}

func readI64BE(s []byte) int64 {
	return int64(s[0])<<56 | int64(s[1])<<48 | int64(s[2])<<40 | int64(s[3])<<32 | int64(s[4])<<24 | int64(s[5])<<16 | int64(s[6])<<8 | int64(s[7])
}

func ReadI64LE(s []byte) (int64, error) {
	if len(s) != 8 {
		return 0, FormatLengthError(8)
	}
	return readI64LE(s), nil
}

func readI64LE(s []byte) int64 {
	return int64(s[7])<<56 | int64(s[6])<<48 | int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadU8(s byte) uint8 {
	return uint8(s)
}

func ReadU16BE(s []byte) (uint16, error) {
	if len(s) != 2 {
		return 0, FormatLengthError(2)
	}
	return readU16BE(s), nil
}

func readU16BE(s []byte) uint16 {
	return uint16(s[0])<<8 | uint16(s[1])
}

func ReadU16LE(s []byte) (uint16, error) {
	if len(s) != 2 {
		return 0, FormatLengthError(2)
	}
	return readU16LE(s), nil
}

func readU16LE(s []byte) uint16 {
	return uint16(s[1])<<8 | uint16(s[0])
}

func ReadU24BE(s []byte) (uint32, error) {
	if len(s) != 3 {
		return 0, FormatLengthError(3)
	}
	return readU24BE(s), nil
}

func readU24BE(s []byte) uint32 {
	return uint32(s[0])<<16 | uint32(s[1])<<8 | uint32(s[2])
}

func ReadU24LE(s []byte) (uint32, error) {
	if len(s) != 3 {
		return 0, FormatLengthError(3)
	}
	return readU24LE(s), nil
}

func readU24LE(s []byte) uint32 {
	return uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
}

func ReadU32BE(s []byte) (uint32, error) {
	if len(s) != 4 {
		return 0, FormatLengthError(4)
	}
	return readU32BE(s), nil
}

func readU32BE(s []byte) uint32 {
	return uint32(s[0])<<24 | uint32(s[1])<<16 | uint32(s[2])<<8 | uint32(s[3])
}

func ReadU32LE(s []byte) (uint32, error) {
	if len(s) != 4 {
		return 0, FormatLengthError(4)
	}
	return readU32LE(s), nil
}

func readU32LE(s []byte) uint32 {
	return uint32(s[3])<<24 | uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
}

func ReadU40BE(s []byte) (uint64, error) {
	if len(s) != 5 {
		return 0, FormatLengthError(5)
	}
	return readU40BE(s), nil
}

func readU40BE(s []byte) uint64 {
	return uint64(s[0])<<32 | uint64(s[1])<<24 | uint64(s[2])<<16 | uint64(s[3])<<8 | uint64(s[4])
}

func ReadU40LE(s []byte) (uint64, error) {
	if len(s) != 5 {
		return 0, FormatLengthError(5)
	}
	return readU40LE(s), nil
}

func readU40LE(s []byte) uint64 {
	return uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadU48BE(s []byte) (uint64, error) {
	if len(s) != 6 {
		return 0, FormatLengthError(6)
	}
	return readU48BE(s), nil
}

func readU48BE(s []byte) uint64 {
	return uint64(s[0])<<40 | uint64(s[1])<<32 | uint64(s[2])<<24 | uint64(s[3])<<16 | uint64(s[4])<<8 | uint64(s[5])
}

func ReadU48LE(s []byte) (uint64, error) {
	if len(s) != 6 {
		return 0, FormatLengthError(6)
	}
	return readU48LE(s), nil
}

func readU48LE(s []byte) uint64 {
	return uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadU56BE(s []byte) (uint64, error) {
	if len(s) != 7 {
		return 0, FormatLengthError(7)
	}
	return readU56BE(s), nil
}

func readU56BE(s []byte) uint64 {
	return uint64(s[0])<<48 | uint64(s[1])<<40 | uint64(s[2])<<32 | uint64(s[3])<<24 | uint64(s[4])<<16 | uint64(s[5])<<8 | uint64(s[6])
}

func ReadU56LE(s []byte) (uint64, error) {
	if len(s) != 7 {
		return 0, FormatLengthError(7)
	}
	return readU56LE(s), nil
}

func readU56LE(s []byte) uint64 {
	return uint64(s[6])<<48 | uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadU64BE(s []byte) (uint64, error) {
	if len(s) != 8 {
		return 0, FormatLengthError(8)
	}
	return readU64BE(s), nil
}

func readU64BE(s []byte) uint64 {
	return uint64(s[0])<<56 | uint64(s[1])<<48 | uint64(s[2])<<40 | uint64(s[3])<<32 | uint64(s[4])<<24 | uint64(s[5])<<16 | uint64(s[6])<<8 | uint64(s[7])
}

func ReadU64LE(s []byte) (uint64, error) {
	if len(s) != 8 {
		return 0, FormatLengthError(8)
	}
	return readU64LE(s), nil
}

func readU64LE(s []byte) uint64 {
	return uint64(s[7])<<56 | uint64(s[6])<<48 | uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadF32BE(s []byte) (float32, error) {
	if len(s) != 4 {
		return 0, FormatLengthError(4)
	}
	return readF32BE(s), nil
}

func readF32BE(s []byte) float32 {
	return math.Float32frombits(readU32BE(s))
}

func ReadF32LE(s []byte) (float32, error) {
	if len(s) != 4 {
		return 0, FormatLengthError(4)
	}
	return readF32LE(s), nil
}

func readF32LE(s []byte) float32 {
	return math.Float32frombits(readU32LE(s))
}

func ReadF64BE(s []byte) (float64, error) {
	if len(s) != 8 {
		return 0, FormatLengthError(8)
	}
	return readF64BE(s), nil
}

func readF64BE(s []byte) float64 {
	return math.Float64frombits(readU64BE(s))
}

func ReadF64LE(s []byte) (float64, error) {
	if len(s) != 8 {
		return 0, FormatLengthError(8)
	}
	return readF64LE(s), nil
}

func readF64LE(s []byte) float64 {
	return math.Float64frombits(readU64LE(s))
}

func ReadBool(s byte) bool {
	return s != 0
}

func ReadString(s []byte) string {
	return string(s)
}
