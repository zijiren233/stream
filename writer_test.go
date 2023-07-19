package stream

import (
	"bytes"
	"testing"
)

func TestWriter(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	w.Byte('h')
	w.Byte('e')
	w.Byte('l')
	w.Byte('l')
	w.Byte('o')
	w.Byte(' ')
	w.Byte('w')
	w.Byte('o')
	w.Byte('r')
	w.Byte('l')
	w.Byte('d')
	if buf.String() != "hello world" {
		t.Errorf("expected 'hello world', got '%s'", buf.String())
	}
}

func TestWriterI8(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	w.I8(0)
	w.I8(1)
	w.I8(-1)
	w.I8(127)
	w.I8(-128)
	if buf.String() != "\x00\x01\xff\x7f\x80" {
		t.Errorf("expected '\\x00\\x01\\xff\\x7f\\x80', got '%s'", buf.String())
	}
}

func TestWriterU16BE(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	w.U16BE(0)
	w.U16BE(1)
	w.U16BE(65535)
	if buf.String() != "\x00\x00\x00\x01\xff\xff" {
		t.Errorf("expected '\\x00\\x00\\x00\\x01\\xff\\xff', got '%s'", buf.String())
	}
}

func TestWriterI32LE(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf)

	w.I32LE(0)
	w.I32LE(1)
	w.I32LE(-1)
	w.I32LE(2147483647)
	w.I32LE(-2147483648)
	if buf.String() != "\x00\x00\x00\x00\x01\x00\x00\x00\xff\xff\xff\xff\xff\xff\xff\x7f\x00\x00\x00\x80" {
		t.Errorf("expected '\\x00\\x00\\x00\\x00\\x01\\x00\\x00\\x00\\xff\\xff\\xff\\xff\\xff\\xff\\xff\\x7f\\x00\\x00\\x00\\x80', got '%s'", buf.String())
	}
}
