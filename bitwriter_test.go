package stream

import (
	"bytes"
	"testing"
)

func TestBitWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewBitWriter(buf)
	w.WriteBits([]byte{0x00, 0x01, 0x02, 0x03}, 0, 32)
	if buf.String() != "\x00\x01\x02\x03" {
		t.Errorf("expected '\\x00\\x01\\x02\\x03', got '%v'", buf.Bytes())
	}
}
