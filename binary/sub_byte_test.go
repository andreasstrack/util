package binary

import (
	"testing"
)

func TestGetBitsFromByte(t *testing.T) {
	b := byte(0xF5)

	if GetBitsFromByte(b, 0, 4) != 5 {
		t.Errorf("GetBitsFromByte() does not work correctly (b = %#02x).\n", b)
	}
	if GetBitsFromByte(b, 4, 4) != 15 {
		t.Errorf("GetBitsFromByte() does not work correctly (b = %#02x).\n", b)
	}
	if GetBitsFromByte(b, 2, 4) != 13 {
		t.Errorf("GetBitsFromByte() does not work correctly (b = %#02x).\n", b)
	}
	if GetBitsFromByte(b, 2, 3) != 5 {
		t.Errorf("GetBitsFromByte() does not work correctly (b = %#02x).\n", b)
	}
	if GetBitsFromByte(b, 5, 1) != 1 {
		t.Errorf("GetBitsFromByte() does not work correctly (b = %#02x).\n", b)
	}
	if GetBitsFromByte(b, 5, 3) != 7 {
		t.Errorf("GetBitsFromByte() does not work correctly (b = %#02x).\n", b)
	}
}

func TestSetBitsToByte(t *testing.T) {
	b := byte(0x00)

	if SetBitsToByte(&b, 0xF0, 4, 4); b != 0xF0 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02x).\n", b)
	}

	if SetBitsToByte(&b, 0x0F, 1, 2); b != 0xF6 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}

	if SetBitsToByte(&b, 0x5F, 4, 4); b != 0x56 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}

	if SetBitsToByte(&b, 0x5F, 4, 4); b != 0x56 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}

	if SetBitsToByte(&b, 0x50, 4, 4); b != 0x56 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}

	if SetBitsToByte(&b, 0x80, 7, 1); b != 0xD6 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}

	// Sets a '0' byte (index 6 of 0x80 is '0')
	if SetBitsToByte(&b, 0x80, 6, 1); b != 0x96 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}

	if SetBitsToByte(&b, 0x34, 0, 8); b != 0x34 {
		t.Errorf("SetBitsToByte() does not work correctly (b = %#02X).\n", b)
	}
}
