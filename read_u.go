package stream

func U8(s byte) uint8 {
	return uint8(s)
}

func U16BE(s []byte) uint16 {
	return uint16(s[0])<<8 | uint16(s[1])
}

func U16LE(s []byte) uint16 {
	return uint16(s[1])<<8 | uint16(s[0])
}

func U24BE(s []byte) uint32 {
	return uint32(s[0])<<16 | uint32(s[1])<<8 | uint32(s[2])
}

func U24LE(s []byte) uint32 {
	return uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
}

func U32BE(s []byte) uint32 {
	return uint32(s[0])<<24 | uint32(s[1])<<16 | uint32(s[2])<<8 | uint32(s[3])
}

func U32LE(s []byte) uint32 {
	return uint32(s[3])<<24 | uint32(s[2])<<16 | uint32(s[1])<<8 | uint32(s[0])
}

func U40BE(s []byte) uint64 {
	return uint64(s[0])<<32 | uint64(s[1])<<24 | uint64(s[2])<<16 | uint64(s[3])<<8 | uint64(s[4])
}

func U40LE(s []byte) uint64 {
	return uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func U48BE(s []byte) uint64 {
	return uint64(s[0])<<40 | uint64(s[1])<<32 | uint64(s[2])<<24 | uint64(s[3])<<16 | uint64(s[4])<<8 | uint64(s[5])
}

func U48LE(s []byte) uint64 {
	return uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func U56BE(s []byte) uint64 {
	return uint64(s[0])<<48 | uint64(s[1])<<40 | uint64(s[2])<<32 | uint64(s[3])<<24 | uint64(s[4])<<16 | uint64(s[5])<<8 | uint64(s[6])
}

func U56LE(s []byte) uint64 {
	return uint64(s[6])<<48 | uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}

func U64BE(s []byte) uint64 {
	return uint64(s[0])<<56 | uint64(s[1])<<48 | uint64(s[2])<<40 | uint64(s[3])<<32 | uint64(s[4])<<24 | uint64(s[5])<<16 | uint64(s[6])<<8 | uint64(s[7])
}

func U64LE(s []byte) uint64 {
	return uint64(s[7])<<56 | uint64(s[6])<<48 | uint64(s[5])<<40 | uint64(s[4])<<32 | uint64(s[3])<<24 | uint64(s[2])<<16 | uint64(s[1])<<8 | uint64(s[0])
}
