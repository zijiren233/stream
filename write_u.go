package stream

func PutU8(buf []byte, v uint8) {
	buf[0] = byte(v)
}

func PutU16BE(buf []byte, v uint16) {
	buf[0] = byte(v >> 8)
	buf[1] = byte(v)
}

func PutU16LE(buf []byte, v uint16) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
}

func PutU24BE(buf []byte, v uint32) {
	buf[0] = byte(v >> 16)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v)
}

func PutU24LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
}

func PutU32BE(buf []byte, v uint32) {
	buf[0] = byte(v >> 24)
	buf[1] = byte(v >> 16)
	buf[2] = byte(v >> 8)
	buf[3] = byte(v)
}

func PutU32LE(buf []byte, v uint32) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
}

func PutU40BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 32)
	buf[1] = byte(v >> 24)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 8)
	buf[4] = byte(v)
}

func PutU40LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
}

func PutU48BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 40)
	buf[1] = byte(v >> 32)
	buf[2] = byte(v >> 24)
	buf[3] = byte(v >> 16)
	buf[4] = byte(v >> 8)
	buf[5] = byte(v)
}

func PutU48LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
}

func PutU56BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 48)
	buf[1] = byte(v >> 40)
	buf[2] = byte(v >> 32)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 16)
	buf[5] = byte(v >> 8)
	buf[6] = byte(v)
}

func PutU56LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
}

func PutU64BE(buf []byte, v uint64) {
	buf[0] = byte(v >> 56)
	buf[1] = byte(v >> 48)
	buf[2] = byte(v >> 40)
	buf[3] = byte(v >> 32)
	buf[4] = byte(v >> 24)
	buf[5] = byte(v >> 16)
	buf[6] = byte(v >> 8)
	buf[7] = byte(v)
}

func PutU64LE(buf []byte, v uint64) {
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	buf[4] = byte(v >> 32)
	buf[5] = byte(v >> 40)
	buf[6] = byte(v >> 48)
	buf[7] = byte(v >> 56)
}
