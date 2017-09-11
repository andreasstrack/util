package binary

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"

	R "github.com/andreasstrack/util/reflect"
)

// WriteStructBigEndian does a binary serizaling from the struct represented
// by s to w. It will write the data in BigEndian byte order.
func WriteStructBigEndian(s interface{}, w io.Writer) error {
	return WriteStruct(s, w, binary.BigEndian)
}

// WriteStructLittleEndian does a binary serizaling from the struct represented
// by s to w. It will write the data in LittleEndian byte order.
func WriteStructLittleEndian(s interface{}, w io.Writer) error {
	return WriteStruct(s, w, binary.LittleEndian)
}

// WriteStruct does a binary serizaling from the struct represented
// by s to w. It will write the data the format specified by o.
func WriteStruct(s interface{}, w io.Writer, o binary.ByteOrder) error {
	if !R.IsStruct(s) {
		return errors.New("WriteStruct(): s must be a struct.")
	}

	fields, _ := R.GetAllValues(s, R.FlagIsAddressable)

	for i := range fields {
		v := fields[i]
		if !v.CanInterface() {
			return fmt.Errorf("Value %s cannot be interfaced.", v)
		}

		if err := writeValue(v, w, o); err != nil {
			return err
		}
	}

	return nil
}

func writeValue(v reflect.Value, w io.Writer, o binary.ByteOrder) error {
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
		return fmt.Errorf("writeValue(): Invalid value kind: %s\n", v.Kind())
	}

	return binary.Write(w, o, v.Addr().Interface())
}
