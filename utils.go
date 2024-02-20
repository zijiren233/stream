package stream

import (
	"io"
	"unsafe"
)

// The change of bytes will cause the change of string synchronously
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// If string is readonly, modifying bytes will cause panic
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BitToU8(0b1001_0110, 0, 3) == 0b0000_0100
func BitToU8(b byte, start int8, end int8) uint8 {
	b <<= start
	b >>= 8 - end
	return b
}

// BitsCut([]byte{0b0110_1101}, 1, 4, BigEndian) -> []byte{0b0000_0110}
// BitsCut([]byte{0b0110_1101, 0b1001_0010}, 1, 9, BigEndian) -> []byte{0b1101_1011}
// BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 2, 14, BigEndian) -> []byte{0b1011_0110, 0b0000_0100}
// BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 2, 18, BigEndian) -> []byte{0b1011_0110, 0b0100_1010}
// BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 9, BigEndian) -> []byte{0b0000_0011}
// BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 15, BigEndian) -> []byte{0b1100_1001}
// BitsCut([]byte{0b0110_1101, 0b1001_0010, 0b1011_0111}, 7, 20, LittleEndian) -> []byte{0b0000_1011, 0b1100_1001}
func BitsCut(data []byte, start, end int, o Order) ([]byte, error) {
	if start < 0 || end > len(data)*8 {
		return nil, io.ErrShortBuffer
	}
	// compute result size, ceiling of (end - start) / 8
	resultSize := (end - start + 7) / 8
	result := make([]byte, resultSize)

	switch o {
	case BigEndian:
		for i := start; i < end; i++ {
			resultByteIndex := (i - start) / 8

			// shift result byte left by 1, and set the last bit to the bit at data current bit index
			result[resultByteIndex] = result[resultByteIndex] << 1
			if data[i/8]&(1<<(7-i%8)) != 0 {
				result[resultByteIndex] |= 1
			}
		}
	case LittleEndian:
		for i := start; i < end; i++ {
			resultByteIndex := resultSize - 1 - (i-start)/8

			// shift result byte left by 1, and set the last bit to the bit at data current bit index
			result[resultByteIndex] = result[resultByteIndex] << 1
			if data[i/8]&(1<<(7-i%8)) != 0 {
				result[resultByteIndex] |= 1
			}
		}
	}

	return result, nil
}

// func BitsTo[T uint8|uint16|uint32|uint64](data []byte, start uint, end uint,o Order) T {
// }

// CopyBits copies bits from src to dst, starting at offset.
// offset is the bit index of dst.
// len(dst)*8 must be greater than or equal to len(src)*8 + offset.
//
// src := []byte{0b10110101, 0b10100011}
// dst := []byte{0b00001111, 0b11110000, 0b10100101}
// CopyBits(src, dst, 2) // dst == []byte{0b00101101, 0b01101000, 0b11100101}
func CopyBits(src, dst []byte, offset int) error {
	if len(src)*8+offset >= len(dst)*8 {
		return io.ErrShortBuffer
	}
	for _, b := range src {
		for i := 0; i < 8; i++ {
			bit := (b >> (7 - i)) & 1
			dstIdx := offset / 8
			dstBitIdx := 7 - (offset % 8)
			dst[dstIdx] = (dst[dstIdx] & ^(1 << dstBitIdx)) | (bit << dstBitIdx)
			offset++
		}
	}
	return nil
}
