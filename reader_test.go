package stream

import (
	"bytes"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf)

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
	r := NewReader(buf)

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

func TestReaderI16BE(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf)

	var i int16
	r.I16BE(&i)
	if i != 0x6865 {
		t.Errorf("expected 0x6865, got 0x%x", i)
	}
	r.I16BE(&i)
	if i != 0x6c6c {
		t.Errorf("expected 0x6c6c, got 0x%x", i)
	}
	r.I16BE(&i)
	if i != 0x6f20 {
		t.Errorf("expected 0x6f20, got 0x%x", i)
	}
	r.I16BE(&i)
	if i != 0x776f {
		t.Errorf("expected 0x776f, got 0x%x", i)
	}
	r.I16BE(&i)
	if i != 0x726c {
		t.Errorf("expected 0x726c, got 0x%x", i)
	}
	r.I16BE(&i)
	if r.Error() != io.ErrUnexpectedEOF {
		t.Errorf("expected EOF, got %v", r.Error())
	}
}

func TestReaderI16LE(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf)

	var i int16
	r.I16LE(&i)
	if i != 0x6568 {
		t.Errorf("expected 0x6568, got 0x%x", i)
	}
	r.I16LE(&i)
	if i != 0x6c6c {
		t.Errorf("expected 0x6c6c, got 0x%x", i)
	}
	r.I16LE(&i)
	if i != 0x206f {
		t.Errorf("expected 0x206f, got 0x%x", i)
	}
	r.I16LE(&i)
	if i != 0x6f77 {
		t.Errorf("expected 0x6f77, got 0x%x", i)
	}
	r.I16LE(&i)
	if i != 0x6c72 {
		t.Errorf("expected 0x6c72, got 0x%x", i)
	}
	r.I16LE(&i)
	if r.Error() != io.ErrUnexpectedEOF {
		t.Errorf("expected EOF, got %v", r.Error())
	}
}

func TestReaderI32BE(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf)

	var i int32
	r.I32BE(&i)
	if i != 0x68656c6c {
		t.Errorf("expected 0x68656c6c, got 0x%x", i)
	}
	r.I32BE(&i)
	if i != 0x6f20776f {
		t.Errorf("expected 0x6f20776f, got 0x%x", i)
	}
	r.I32BE(&i)
	if r.Error() != io.ErrUnexpectedEOF {
		t.Errorf("expected EOF, got %v", r.Error())
	}
}

func TestReaderI32LE(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	r := NewReader(buf)

	var i int32
	r.I32LE(&i)
	if i != 0x6c6c6568 {
		t.Errorf("expected 0x6c6c6568, got 0x%x", i)
	}
	r.I32LE(&i)
	if i != 0x6f77206f {
		t.Errorf("expected 0x6f77206f, got 0x%x", i)
	}
	r.I32LE(&i)
	if r.Error() != io.ErrUnexpectedEOF {
		t.Errorf("expected EOF, got %v", r.Error())
	}
}

func TestReadSeeker(t *testing.T) {
	buf := bytes.NewReader([]byte("hello world"))
	r := NewReadSeeker(buf)

	var i int32
	r.I32LE(&i)
	if i != 0x6c6c6568 {
		t.Errorf("expected 0x6c6c6568, got 0x%x", i)
	}

	// Seek back to the beginning
	r.Seek(0, 0).I32LE(&i)

	if i != 0x6c6c6568 {
		t.Errorf("expected 0x6c6c6568, got 0x%x", i)
	}
}
