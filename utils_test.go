package stream

import "testing"

func TestBitToIU8(t *testing.T) {
	if BitToU8(0b10010110, 0, 3) != 0b00000110 {
		t.Error("BitToIU8 failed")
	}
}
