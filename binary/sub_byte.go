package binary

import (
	"fmt"
)

func GetBitsFromByte(b byte, begin, numBits uint8) uint8 {
	if begin > 7 || numBits+begin > 8 {
		panic(fmt.Sprintf("Indizes must remain within 8 bits (begin: %d, numBits: %d)\n", begin, numBits))
	}
	result := uint8(b) >> begin
	result = result & (uint8(0xFF) >> (uint8(8) - numBits))

	return result
}

func SetBitsToByte(b *byte, bits byte, begin, numBits uint8) {
	if begin > 7 || numBits+begin > 8 {
		panic(fmt.Sprintf("Indizes must remain within 8 bits (begin: %d, numBits: %d)\n", begin, numBits))
	}

	deleteMask := (uint8(0xFF) >> (uint8(8) - begin)) | (uint8(0xFF) << (begin + numBits))
	setMask := GetBitsFromByte(bits, 0, numBits) << begin

	*b = *b & byte(deleteMask)
	*b = *b | byte(setMask)
}
