package binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	R "github.com/andreasstrack/util/reflect"
	"io"
	"reflect"
)

func WriteStructBigEndian(s interface{}, w io.Writer) error {
	return WriteStruct(s, w, binary.BigEndian)
}

func WriteStructLittleEndian(s interface{}, w io.Writer) error {
	return WriteStruct(s, w, binary.LittleEndian)
}

func WriteStruct(s interface{}, w io.Writer, o binary.ByteOrder) error {
	if !R.IsStruct(s) {
		return errors.New("WriteStruct(): s must be a struct.")
	}

	fields := R.GetAllAddressableFields(s)

	for i := range fields {
		v := fields[i]
		if !v.CanInterface() {
			return errors.New(fmt.Sprintf("Value %s cannot be interfaced.", v))
		}

		if err := writeValue(v, w, o); err != nil {
			return err
		}
	}

	return nil
}

func writeValue(v reflect.Value, w io.Writer, o binary.ByteOrder) error {
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
		fmt.Printf("Invalid value kind: %s\n", v.Kind())
		return errors.New(fmt.Sprintf("writeValue(): Invalid value kind: %s\n", v.Kind()))
	}

	return binary.Write(w, o, v.Addr().Interface())
}
