package stream

import (
	"bytes"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf, BigEndian)

	var b byte
	r.Byte(&b)
	if b != 'h' {
		t.Errorf("expected 'h', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'e' {
		t.Errorf("expected 'e', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'l' {
		t.Errorf("expected 'l', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'l' {
		t.Errorf("expected 'l', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'o' {
		t.Errorf("expected 'o', got '%c'", b)
	}
	r.Byte(&b)
	if b != ' ' {
		t.Errorf("expected ' ', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'w' {
		t.Errorf("expected 'w', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'o' {
		t.Errorf("expected 'o', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'r' {
		t.Errorf("expected 'r', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'l' {
		t.Errorf("expected 'l', got '%c'", b)
	}
	r.Byte(&b)
	if b != 'd' {
		t.Errorf("expected 'd', got '%c'", b)
	}
	r.Byte(&b)
	if r.Error() != io.EOF {
		t.Errorf("expected EOF, got %v", r.Error())
	}
}

func TestReaderBytes(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf, BigEndian)

	b := make([]byte, 5)
	r.Bytes(b)
	if string(b) != "hello" {
		t.Errorf("expected 'hello', got '%s'", string(b))
	}
	r.Bytes(b)
	if string(b) != " worl" {
		t.Errorf("expected ' worl', got '%s'", string(b))
	}
	r.Bytes(b)
	if r.Error() != io.ErrUnexpectedEOF {
		t.Errorf("expected EOF, got %v", r.Error())
	}
}

func TestBits(t *testing.T) {
	buf := bytes.NewReader([]byte{0b1101_1010})
	rb := NewReader(buf, BigEndian)

	var b bool
	for k, v := range []bool{true, true, false, true, true, false, true, false} {
		if rb.Bits(1, &b).Error() != nil {
			t.Errorf("expected no error, got %v", rb.Error())
		}
		if b != v {
			t.Errorf("expected %v, got %v at bit %d", v, b, k)
		}
	}

	ru8 := NewReader(bytes.NewReader([]byte{12, 42, 13, 45, 34, 76, 84}), BigEndian)
	var u8 uint8
	for k, v := range []uint8{12, 42, 13, 45, 34, 76, 84} {
		if ru8.Bits(8, &u8).Error() != nil {
			t.Errorf("expected no error, got %v", ru8.Error())
		}
		if u8 != v {
			t.Errorf("expected %v, got %v at bit %d", v, u8, k*8)
		}
	}

}

func TestBits2(t *testing.T) {
	ru8 := NewReader(bytes.NewReader([]byte{12, 42, 13, 45, 34, 76, 84}), BigEndian)
	for i, v := range []byte{12, 42, 13, 45, 34, 76, 84} {
		for j := 0; j < 2; j++ {
			var u8 uint8
			if ru8.Bits(4, &u8).Error() != nil {
				t.Errorf("expected no error, got %v", ru8.Error())
				return
			}
			if j%2 == 0 {
				if !bytes.Equal([]byte{v >> 4}, []byte{u8 & 0x0f}) {
					t.Errorf("expected %08b, got %08b at bit %d", v>>4, u8&0x0f, i*4)
					return
				}
			} else {
				if !bytes.Equal([]byte{v & 0x0f}, []byte{u8}) {
					t.Errorf("expected %08b, got %08b at bit %d", v&0x0f, u8, i*4)
					return
				}
			}
		}
	}
}
