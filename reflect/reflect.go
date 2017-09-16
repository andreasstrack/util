// Package reflect provides utility functions for reflecting
// on go types that add further functionality to the built-in
// reflect package.
package reflect

import (
	"fmt"
	"reflect"

	"github.com/andreasstrack/data/tree"
	"github.com/andreasstrack/util"
	"github.com/andreasstrack/util/patterns"
)

// GetAllValues traverses the value tree of i and returns the values
// filtered by the given flags.
func GetAllValues(i interface{}, flags util.Flags) []*ValueNode {
	var nodes []*ValueNode
	if it, err := NewValueIterator(i, flags, tree.BreadthFirst); err == nil {
		allAsInterface := patterns.GetAll(it)
		for i := range allAsInterface {
			nodes = append(nodes, allAsInterface[i].(*ValueNode))
		}
	}
	return nodes
}

func GetSlice(slice interface{}) (*[]interface{}, error) {
	v := GetElementValue(slice)
	if v.Kind() != reflect.Slice || !v.CanAddr() {
		return nil, fmt.Errorf("Need an addressable slice to initialize set, given value was %#v", slice)
	}

	var s []interface{}
	for i := 0; i < v.Len(); i++ {
		vi := v.Index(i)
		if !vi.CanInterface() {
			return nil, fmt.Errorf("Need a value that can interface at index %d of slice", i)
		}
		s = append(s, vi.Interface())
	}
	return &s, nil
}

// IsPointer returns whether i represents a pointer value or not.
func IsPointer(i interface{}) bool {
	return reflect.ValueOf(i).Kind() == reflect.Ptr
}

// GetElementValue returns:
// - The value of i if i does not represent a pointer.
// - The element value of i if i represents a pointer.
func GetElementValue(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr {
		return v.Elem()
	}

	return v
}

// GetElementType returns the type of the element value
// provided by GetElementValue
func GetElementType(i interface{}) reflect.Type {
	return GetElementValue(i).Type()
}

// IsStruct returns:
// - True if i represents a struct or a pointer to a struct
// - False otherwise
func IsStruct(i interface{}) bool {
	var v reflect.Value

	if v = GetElementValue(i); v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}

// HasField returns:
// - True if IsStruct(i) is true and the struct referred
//   to by i has a field named name.
// - False otherwise
func HasField(i interface{}, name string) bool {
	if !IsStruct(i) {
		return false
	}

	return GetElementValue(i).FieldByName(name).Kind() != reflect.Invalid
}

// CopyStruct performs a deep copy of the struct represented by
// from to the struct represented by to.
// Note that to must refer to a pointer to a struct for it to be addressable and thus setable.
func CopyStruct(from interface{}, to interface{}) error {
	fromValue := GetElementValue(from)
	toValue := GetElementValue(to)

	if fromValue.Kind() != reflect.Struct ||
		toValue.Kind() != reflect.Struct {
		return fmt.Errorf("Need two structs, got %v (from) and %v (to)", from, to)
	}

	nFieldsFrom := fromValue.NumField()
	nFieldsTo := toValue.NumField()

	if nFieldsFrom != nFieldsTo {
		return fmt.Errorf("Number of fields in from (%d) does not match number of fields in to (%d)", nFieldsFrom, nFieldsTo)
	}

	for i := 0; i < nFieldsFrom; i++ {
		fromField := fromValue.Field(i)
		toField := toValue.Field(i)

		fromFieldKind := fromField.Kind()
		toFieldKind := toField.Kind()

		if fromFieldKind != toFieldKind {
			return fmt.Errorf("Field for from (%v) and to (%v) do not match", fromField, toField)
		}

		switch fromFieldKind {
		case reflect.Int:
			toField.SetInt(fromField.Int())
			break
		case reflect.Float32:
		case reflect.Float64:
			toField.SetFloat(fromField.Float())
			break
		case reflect.Complex64:
		case reflect.Complex128:
			toField.SetComplex(fromField.Complex())
			break
		case reflect.String:
			toField.SetString(fromField.String())
			break
		default:
			return fmt.Errorf("Invalid kind of field: %s", fromFieldKind)
		}
	}

	return nil
}
