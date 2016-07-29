package binary

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"

	R "github.com/andreasstrack/util/reflect"
)

// ReadStructBigEndian does a binary deserializing from r to the
// struct represented by s. It assumes the data to be in BigEndian byte order.
// Note that s must refer to a pointer to a struct for it to be addressable and thus setable.
func ReadStructBigEndian(s interface{}, r io.Reader) error {
	return ReadStruct(s, r, binary.BigEndian)
}

// ReadStructLittleEndian does a binary deserializing from r to the
// struct represented by s. It assumes the data to be in LittleEndian byte order.
// Note that s must refer to a pointer to a struct for it to be addressable and thus setable.
func ReadStructLittleEndian(s interface{}, r io.Reader) error {
	return ReadStruct(s, r, binary.LittleEndian)
}

// ReadStruct does a binary deserializing from r to the
// struct represented by s, where the binary byte order of the data is specified by o.
// Note that s must refer to a pointer to a struct for it to be addressable and thus setable.
func ReadStruct(s interface{}, r io.Reader, o binary.ByteOrder) error {
	if !R.IsStruct(s) {
		return fmt.Errorf("ReadStruct(): s must be a struct. It is: %s\n", reflect.TypeOf(s))
	}

	fields := R.GetAllAddressableFields(s)

	for i := range fields {
		v := fields[i]
		if !v.CanInterface() {
			return fmt.Errorf("Value %s cannot be interfaced.", v)
		}

		if err := readValue(v, r, o); err != nil {
			return err
		}
	}

	return nil
}

func readValue(v reflect.Value, r io.Reader, o binary.ByteOrder) error {
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
		return fmt.Errorf("readValue(): Invalid value kind: %s\n", v.Kind())
	}

	return binary.Read(r, o, v.Addr().Interface())
}
