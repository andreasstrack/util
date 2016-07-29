package binary

import "fmt"

// GetBitsFromByte extracts a number of bits from a byte.
// It returns numBits starting at begin in the byte b as
// a uint8 value.
// If any of the bits it is told to return are above the 8bit
// border, it will panic.
func GetBitsFromByte(b byte, begin, numBits uint8) uint8 {
	if begin > 7 || numBits+begin > 8 {
		panic(fmt.Sprintf("Indizes must remain within 8 bits (begin: %d, numBits: %d)\n", begin, numBits))
	}

	result := uint8(b) >> begin
	result = result & (uint8(0xFF) >> (uint8(8) - numBits))

	return result
}

// SetBitsToByte sets a number of bits to a byte.
// It sets numBits starting at begin from the byte bits
// to the byte b.
// If any of the bits it is told to set are above the 8bit
// border, it will panic.
func SetBitsToByte(b *byte, bits byte, begin, numBits uint8) {
	if begin > 7 || numBits+begin > 8 {
		panic(fmt.Sprintf("Indizes must remain within 8 bits (begin: %d, numBits: %d)\n", begin, numBits))
	}

	deleteMask := (uint8(0xFF) >> (uint8(8) - begin)) | (uint8(0xFF) << (begin + numBits))
	setMask := GetBitsFromByte(bits, begin, numBits) << begin

	*b = *b & byte(deleteMask)
	*b = *b | byte(setMask)
}
