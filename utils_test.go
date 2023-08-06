package stream

import (
	"bytes"
	"testing"
)

func TestBitToIU8(t *testing.T) {
	if BitToU8(0b1001_0110, 0, 3) != 0b0000_0100 {
		t.Error("BitToIU8 failed")
	}
}

func TestBitsCut(t *testing.T) {
	if b, _ := BitsCut([]byte{0b0110_1101}, 1, 4, BigEndian); !bytes.Equal(b, []byte{0b0000_0110}) {
		t.Error("BitsCut failed")
	}
	if b, _ := BitsCut([]byte{0b0110_1101, 0b1001_0010}, 1, 9, BigEndian); !bytes.Equal(b, []byte{0b1101_1011}) {
		t.Error("BitsCut failed")
	}
	if b, _ := BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 2, 14, BigEndian); !bytes.Equal(b, []byte{0b1011_0110, 0b0000_0100}) {
		t.Error("BitsCut failed")
	}
	if b, _ := BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 2, 18, BigEndian); !bytes.Equal(b, []byte{0b1011_0110, 0b0100_1010}) {
		t.Error("BitsCut failed")
	}
	if b, _ := BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 9, BigEndian); !bytes.Equal(b, []byte{0b0000_0011}) {
		t.Error("BitsCut failed")
	}
	if b, _ := BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 15, BigEndian); !bytes.Equal(b, []byte{0b1100_1001}) {
		t.Error("BitsCut failed")
	}
	// out of range
	if _, err := BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 100, BigEndian); err == nil {
		t.Error("BitsCut failed")
	}
	if b, _ := BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 20, LittleEndian); !bytes.Equal(b, []byte{0b0000_1011, 0b1100_1001}) {
		t.Errorf("BitsCut failed, get: %16b, need: %16b", b, []byte{0b0000_1011, 0b1100_1001})
	}
}

func TestCopyBit(t *testing.T) {
	src := []byte{0b10110101, 0b10100011}
	dst := []byte{0b00001111, 0b11110000, 0b10100101}

	CopyBits(src, dst, 2)

	if !bytes.Equal(dst, []byte{0b00101101, 0b01101000, 0b11100101}) {
		t.Error("CopyBit failed")
	}
}
