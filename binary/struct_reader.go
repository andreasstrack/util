package binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	R "github.com/andreasstrack/util/reflect"
	"io"
	"reflect"
)

func ReadStructBigEndian(s interface{}, r io.Reader) error {
	return ReadStruct(s, r, binary.BigEndian)
}

func ReadStructLittleEndian(s interface{}, r io.Reader) error {
	return ReadStruct(s, r, binary.LittleEndian)
}

func ReadStruct(s interface{}, r io.Reader, o binary.ByteOrder) error {
	if !R.IsStruct(s) {
		return errors.New(fmt.Sprintf("ReadStruct(): s must be a struct. It is: %s\n", reflect.TypeOf(s)))
	}

	fields := R.GetAllAddressableFields(s)

	for i := range fields {
		v := fields[i]
		if !v.CanInterface() {
			return errors.New(fmt.Sprintf("Value %s cannot be interfaced.", v))
		}

		if err := readValue(v, r, o); err != nil {
			return err
		}
	}

	return nil
}

func readValue(v reflect.Value, r io.Reader, o binary.ByteOrder) error {
	if v.Kind() == reflect.Ptr {
		return errors.New(fmt.Sprintf("Value %s must not be a pointer."))
	}

	switch v.Kind() {
	case reflect.Int8:
	case reflect.Uint8:
	case reflect.Int16:
	case reflect.Uint16:
	case reflect.Int32:
	case reflect.Uint32:
	case reflect.Int64:
	case reflect.Uint64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Array:
	case reflect.Slice:

	default:
		fmt.Printf("Invalid value kind.\n")
		return errors.New(fmt.Sprintf("readValue(): Invalid value kind: %s\n", v.Kind()))
	}

	return binary.Read(r, o, v.Addr().Interface())
}
