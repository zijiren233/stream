package stream

func I8(s byte) int8 {
	return int8(s)
}

func I16BE(s []byte) int16 {
	return int16(s[0])<<8 | int16(s[1])
}

func I16LE(s []byte) int16 {
	return int16(s[1])<<8 | int16(s[0])
}

func I24BE(s []byte) int32 {
	return int32(s[0])<<16 | int32(s[1])<<8 | int32(s[2])
}

func I24LE(s []byte) int32 {
	return int32(s[2])<<16 | int32(s[1])<<8 | int32(s[0])
}

func I32BE(s []byte) int32 {
	return int32(s[0])<<24 | int32(s[1])<<16 | int32(s[2])<<8 | int32(s[3])
}

func I32LE(s []byte) int32 {
	return int32(s[3])<<24 | int32(s[2])<<16 | int32(s[1])<<8 | int32(s[0])
}

func I40BE(s []byte) int64 {
	return int64(s[0])<<32 | int64(s[1])<<24 | int64(s[2])<<16 | int64(s[3])<<8 | int64(s[4])
}

func I40LE(s []byte) int64 {
	return int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func I48BE(s []byte) int64 {
	return int64(s[0])<<40 | int64(s[1])<<32 | int64(s[2])<<24 | int64(s[3])<<16 | int64(s[4])<<8 | int64(s[5])
}

func I48LE(s []byte) int64 {
	return int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func I56BE(s []byte) int64 {
	return int64(s[0])<<48 | int64(s[1])<<40 | int64(s[2])<<32 | int64(s[3])<<24 | int64(s[4])<<16 | int64(s[5])<<8 | int64(s[6])
}

func I56LE(s []byte) int64 {
	return int64(s[6])<<48 | int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}

func I64BE(s []byte) int64 {
	return int64(s[0])<<56 | int64(s[1])<<48 | int64(s[2])<<40 | int64(s[3])<<32 | int64(s[4])<<24 | int64(s[5])<<16 | int64(s[6])<<8 | int64(s[7])
}

func I64LE(s []byte) int64 {
	return int64(s[7])<<56 | int64(s[6])<<48 | int64(s[5])<<40 | int64(s[4])<<32 | int64(s[3])<<24 | int64(s[2])<<16 | int64(s[1])<<8 | int64(s[0])
}
