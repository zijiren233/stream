package stream

import (
	"math"
)

func F32BE(s []byte) float32 {
	return math.Float32frombits(U32BE(s))
}

func F32LE(s []byte) float32 {
	return math.Float32frombits(U32LE(s))
}

func F64BE(s []byte) float64 {
	return math.Float64frombits(U64BE(s))
}

func F64LE(s []byte) float64 {
	return math.Float64frombits(U64LE(s))
}
