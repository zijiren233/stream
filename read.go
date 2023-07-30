package stream

import "math"

func ReadI8(s []byte) int8 {
	return int8(s[0])
}

func ReadI16BE(s []byte) int16 {
	return int16(s[0])<<8 | int16(s[1])
}

func ReadI16LE(s []byte) int16 {
	return int16(s[1])<<8 | int16(s[0])
}

func ReadI24BE(s []byte) int32 {
	return int32(s[0])<<16 | int32(s[1])<<8 | int32(s[2])
}

func ReadI24LE(s []byte) int32 {
	return int32(s[2])<<16 | int32(s[1])<<8 | int32(s[0])
}

func ReadI32BE(s []byte) int32 {
	return int32(s[0])<<24 | int32(s[1])<<16 | int32(s[2])<<8 | int32(s[3])
}

func ReadI32LE(s []byte) int32 {
	return int32(s[3])<<24 | int32(s[2])<<16 | int32(s[1])<<8 | int32(s[0])
}

func ReadI40BE(s []byte) int64 {
	return int64(s[0])<<32 | int64(s[1])<<24 | int64(s[2])<<16 | int64(s[3])<<8 | int64(s[4])
}

func ReadI40LE(s []byte) int64 {
	return int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadI48BE(s []byte) int64 {
	return int64(s[0])<<40 | int64(s[1])<<32 | int64(s[2])<<24 | int64(s[3])<<16 | int64(s[4])<<8 | int64(s[5])
}

func ReadI48LE(s []byte) int64 {
	return int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadI56BE(s []byte) int64 {
	return int64(s[0])<<48 | int64(s[1])<<40 | int64(s[2])<<32 | int64(s[3])<<24 | int64(s[4])<<16 | int64(s[5])<<8 | int64(s[6])
}

func ReadI56LE(s []byte) int64 {
	return int64(s[6])<<48 | int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadI64BE(s []byte) int64 {
	return int64(s[0])<<56 | int64(s[1])<<48 | int64(s[2])<<40 | int64(s[3])<<32 | int64(s[4])<<24 | int64(s[5])<<16 | int64(s[6])<<8 | int64(s[7])
}

func ReadI64LE(s []byte) int64 {
	return int64(s[7])<<56 | int64(s[6])<<48 | int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func ReadU8(s []byte) uint8 {
	return uint8(s[0])
}

func ReadU16BE(s []byte) uint16 {
	return uint16(s[0])<<8 | uint16(s[1])
}

func ReadU16LE(s []byte) uint16 {
	return uint16(s[1])<<8 | uint16(s[0])
}

func ReadU24BE(s []byte) uint32 {
	return uint32(s[0])<<16 | uint32(s[1])<<8 | uint32(s[2])
}

func ReadU24LE(s []byte) uint32 {
	return uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
}

func ReadU32BE(s []byte) uint32 {
	return uint32(s[0])<<24 | uint32(s[1])<<16 | uint32(s[2])<<8 | uint32(s[3])
}

func ReadU32LE(s []byte) uint32 {
	return uint32(s[3])<<24 | uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
}

func ReadU40BE(s []byte) uint64 {
	return uint64(s[0])<<32 | uint64(s[1])<<24 | uint64(s[2])<<16 | uint64(s[3])<<8 | uint64(s[4])
}

func ReadU40LE(s []byte) uint64 {
	return uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadU48BE(s []byte) uint64 {
	return uint64(s[0])<<40 | uint64(s[1])<<32 | uint64(s[2])<<24 | uint64(s[3])<<16 | uint64(s[4])<<8 | uint64(s[5])
}

func ReadU48LE(s []byte) uint64 {
	return uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadU56BE(s []byte) uint64 {
	return uint64(s[0])<<48 | uint64(s[1])<<40 | uint64(s[2])<<32 | uint64(s[3])<<24 | uint64(s[4])<<16 | uint64(s[5])<<8 | uint64(s[6])
}

func ReadU56LE(s []byte) uint64 {
	return uint64(s[6])<<48 | uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadU64BE(s []byte) uint64 {
	return uint64(s[0])<<56 | uint64(s[1])<<48 | uint64(s[2])<<40 | uint64(s[3])<<32 | uint64(s[4])<<24 | uint64(s[5])<<16 | uint64(s[6])<<8 | uint64(s[7])
}

func ReadU64LE(s []byte) uint64 {
	return uint64(s[7])<<56 | uint64(s[6])<<48 | uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func ReadF32BE(s []byte) float32 {
	return math.Float32frombits(ReadU32BE(s))
}

func ReadF32LE(s []byte) float32 {
	return math.Float32frombits(ReadU32LE(s))
}

func ReadF64BE(s []byte) float64 {
	return math.Float64frombits(ReadU64BE(s))
}

func ReadF64LE(s []byte) float64 {
	return math.Float64frombits(ReadU64LE(s))
}

func ReadBool(s []byte) bool {
	return s[0] != 0
}

func ReadString(s []byte) string {
	return string(s)
}
