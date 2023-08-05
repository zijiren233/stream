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
