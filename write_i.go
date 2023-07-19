package stream

func PutI8(buf []byte, v int8) {
	buf[0] = byte(v)
}

func PutI16BE(buf []byte, v int16) {
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func PutI16LE(buf []byte, v int16) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
}

func PutI24BE(buf []byte, v int32) {
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
}

func PutI24LE(buf []byte, v int32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
}

func PutI32BE(buf []byte, v int32) {
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func PutI32LE(buf []byte, v int32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

func PutI40BE(buf []byte, v int64) {
	buf[0] = byte(v >> 32)
	buf[1] = byte(v >> 24)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 8)
	buf[4] = byte(v)
}

func PutI40LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
}

func PutI48BE(buf []byte, v int64) {
	buf[0] = byte(v >> 40)
	buf[1] = byte(v >> 32)
	buf[2] = byte(v >> 24)
	buf[3] = byte(v >> 16)
	buf[4] = byte(v >> 8)
	buf[5] = byte(v)
}

func PutI48LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
}

func PutI56BE(buf []byte, v int64) {
	buf[0] = byte(v >> 48)
	buf[1] = byte(v >> 40)
	buf[2] = byte(v >> 32)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 16)
	buf[5] = byte(v >> 8)
	buf[6] = byte(v)
}

func PutI56LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
}

func PutI64BE(buf []byte, v int64) {
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
}

func PutI64LE(buf []byte, v int64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
}
