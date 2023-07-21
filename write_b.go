package stream

func PutBool(buf []byte, v bool) {
	if v {
		buf[0] = 1
	} else {
		buf[0] = 0
	}
}
