package stream

import (
	"bytes"
	"io"
	"testing"
)

func TestBitReader(t *testing.T) {
	r := NewBitReader(bytes.NewReader([]byte{0b01100011, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99}))
	buf := make([]byte, 1)
	if err := r.ReadBits(buf, 0, 4); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal(buf, []byte{byte(99) & 0xf0}) {
		t.Fatalf("ReadBits returned the wrong value: %08b, need: %08b", buf, byte(99)&0xf0)
	}
	if err := r.ReadBits(buf, 4, 8); err != nil {
		t.Fatal(err)
	} else if !bytes.Equal([]byte{buf[0] & 0x0f}, []byte{byte(99) & 0x0f}) {
		t.Fatalf("ReadBits returned the wrong value: %08b, need: %08b", []byte{buf[0] & 0x0f}, byte(99)&0x0f)
	}
	if !bytes.Equal(buf, []byte{99}) {
		t.Fatal("ReadBits returned the wrong value")
	}
	buf2 := make([]byte, 2)
	var tmp int16 = 0b01100011 << 7
	for {
		if err := r.ReadBits(buf2, 1, 9); err != nil {
			if err == io.EOF {
				break
			} else {
				t.Fatal(err)
			}
		} else if !bytes.Equal(buf2, []byte{99 >> 1, byte(tmp)}) {
			t.Fatalf("ReadBits returned the wrong value: %16b, need: %16b", buf2, []byte{99 >> 1, byte(tmp)})
		}
	}
}
