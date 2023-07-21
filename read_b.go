package stream

func Bool(s []byte) bool {
	return s[0] != 0
}
