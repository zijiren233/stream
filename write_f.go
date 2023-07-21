package stream

import "math"

func PutF32BE(buf []byte, v float32) {
	PutU32BE(buf, math.Float32bits(v))
}

func PutF32LE(buf []byte, v float32) {
	PutU32LE(buf, math.Float32bits(v))
}

func PutF64BE(buf []byte, v float64) {
	PutU64BE(buf, math.Float64bits(v))
}

func PutF64LE(buf []byte, v float64) {
	PutU64LE(buf, math.Float64bits(v))
}
