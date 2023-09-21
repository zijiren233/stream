package stream

import (
	"bytes"
	"testing"
)

func TestWriterU16BE(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf, BigEndian)

	w.U16(0)
	w.U16(1)
	w.U16(65535)
	if buf.String() != "\x00\x00\x00\x01\xff\xff" {
		t.Errorf("expected '\\x00\\x00\\x00\\x01\\xff\\xff', got '%s'", buf.String())
	}
}

func TestWriterI32LE(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf, LittleEndian)

	w.I32(0)
	w.I32(1)
	w.I32(-1)
	w.I32(2147483647)
	w.I32(-2147483648)
	if buf.String() != "\x00\x00\x00\x00\x01\x00\x00\x00\xff\xff\xff\xff\xff\xff\xff\x7f\x00\x00\x00\x80" {
		t.Errorf("expected '\\x00\\x00\\x00\\x00\\x01\\x00\\x00\\x00\\xff\\xff\\xff\\xff\\xff\\xff\\xff\\x7f\\x00\\x00\\x00\\x80', got '%s'", buf.String())
	}
}

func TestWriteBE(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf, BigEndian)

	type test struct {
		A int8
		B int16
		C int32
	}

	w.Write(test{111, 9999, 88888})

	if w.Error() != nil {
		t.Errorf("expected no error, got %v", w.Error())
		return
	}
	if w.Total() != 7 || w.N() != 7 {
		t.Errorf("expected 7 bytes, got %d", w.Total())
		return
	}
	if bytes.Equal(buf.Bytes(), []byte("\x6f\x27\x0f\x00\x01\x5b\x38")) == false {
		t.Errorf("expected \x6f\x27\x0f\x00\x01\x5b\x38, got %x", buf.Bytes())
		return
	}

	tmp := new(test)
	r := NewReader(buf, BigEndian)
	if r.Read(tmp).Error() != nil {
		t.Errorf("expected no error, got %v", r.Error())
		return
	}
	if r.Total() != 7 || r.N() != 7 {
		t.Errorf("expected 7 bytes, got %d", r.Total())
		return
	}
}

func TestWriteBE2(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf, BigEndian)

	type test struct {
		A int8
		B int16
		C int32
		D []byte
		E [16]byte
		F []*test
	}

	w.Write(&test{111, 9999, 88888, make([]byte, 8), [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, []*test{{111, 9999, 88888, make([]byte, 8), [16]byte{}, nil}}})

	if w.Error() != nil {
		t.Errorf("expected no error, got %v", w.Error())
		return
	}
	if w.Total() != 62 || w.N() != 62 {
		t.Errorf("expected 31 bytes, got %d", w.Total())
		return
	}

	tmp := &test{
		D: make([]byte, 8),
		E: [16]byte{},
		F: []*test{{}},
	}
	r := NewReader(buf, BigEndian)
	if r.Read(tmp).Error() != nil {
		t.Errorf("expected no error, got %v", r.Error())
		return
	}
}

func TestWriteLE(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf, LittleEndian)

	type test struct {
		A int8
		B int16
		C int32
	}

	w.Write(test{111, 9999, 88888})

	if w.Error() != nil {
		t.Errorf("expected no error, got %v", w.Error())
		return
	}
	if w.Total() != 7 || w.N() != 7 {
		t.Errorf("expected 7 bytes, got %d", w.Total())
		return
	}
	if bytes.Equal(buf.Bytes(), []byte("\x6f\x0f\x27\x38\x5b\x01\x00")) == false {
		t.Errorf("expected \x6f\x0f\x27\x38\x5b\x01\x00, got %x", buf.Bytes())
		return
	}

	tmp := new(test)
	r := NewReader(buf, LittleEndian)
	if r.Read(tmp).Error() != nil {
		t.Errorf("expected no error, got %v", r.Error())
		return
	}
	if r.Total() != 7 || r.N() != 7 {
		t.Errorf("expected 7 bytes, got %d", r.Total())
		return
	}
}

func TestWriteBytes(t *testing.T) {
	buf := &bytes.Buffer{}
	w := NewWriter(buf, BigEndian)

	w.Bytes([]byte{1, 2, 3, 4, 5, 6, 7, 8})

	if w.Error() != nil {
		t.Errorf("expected no error, got %v", w.Error())
		return
	}
	if w.Total() != 8 || w.N() != 8 {
		t.Errorf("expected 8 bytes, got %d", w.Total())
		return
	}
	if bytes.Equal(buf.Bytes(), []byte{1, 2, 3, 4, 5, 6, 7, 8}) == false {
		t.Errorf("expected [1 2 3 4 5 6 7 8], got %x", buf.Bytes())
		return
	}

	tmp := make([]byte, 8)
	r := NewReader(buf, BigEndian)
	if r.Read(tmp).Error() != nil {
		t.Errorf("expected no error, got %v", r.Error())
		return
	}
	if r.Total() != 8 || r.N() != 8 {
		t.Errorf("expected 8 bytes, got %d", r.Total())
		return
	}
}
