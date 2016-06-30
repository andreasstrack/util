package binary

import (
	"bytes"
	"fmt"
	"testing"
)

type TestStruct struct {
	Ui16  uint16
	Ui8_1 uint8
	B20   [20]byte
}

func TestStructRead(t *testing.T) {
	buf := GetBytesBufferToFillTestStruct()
	s := TestStruct{Ui16: uint16('A'), Ui8_1: uint8(0)}

	if err := ReadStructBigEndian(&s, bytes.NewReader(buf.Bytes())); err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Printf("Read struct:\n%s\nFrom:\n%q\n", s, buf.Bytes())
}

func TestReadAndWriteStruct(t *testing.T) {
	readBuf := GetBytesBufferToFillTestStruct()
	s := TestStruct{Ui16: uint16('A'), Ui8_1: uint8(1)}

	if err := ReadStructBigEndian(&s, bytes.NewReader(readBuf.Bytes())); err != nil {
		t.Error(err.Error())
		return
	}

	writeBuf := new(bytes.Buffer)

	if err := WriteStructBigEndian(&s, writeBuf); err != nil {
		t.Error(err.Error())
		return
	}

	fmt.Printf("Wrote:\n%s\nRead:\n%s\n", readBuf.Bytes(), writeBuf.Bytes())

	if readBuf.Len() != writeBuf.Len() || readBuf.String() != writeBuf.String() {
		t.Error("Read and write buffer not equal.\n")
	}
}

func GetBytesBufferToFillTestStruct() *bytes.Buffer {
	buf := new(bytes.Buffer)

	for i := 0; i < 3; i++ {
		buf.WriteByte(byte(255))
	}

	for i := 3; i < 23; i++ {
		buf.WriteByte(byte('A') + byte(i))
	}

	return buf
}
