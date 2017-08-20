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

func GetAllValues(i interface{}) ([]reflect.Value, [][]reflect.StructTag) {
	return GetAllValuesWithFlags(i, util.FlagNone)
}

// GetAllValuesWithFlags traverses the value tree of i and returns the values
// filtered by the given flags.
func GetAllValuesWithFlags(i interface{}, flags util.Flags) ([]reflect.Value, [][]reflect.StructTag) {
	var resultValues []reflect.Value
	var resultTags [][]reflect.StructTag
	if it, err := NewValueIterator(i, flags, tree.BreadthFirst); err == nil {
		allAsInterface := patterns.GetAll(it)
		for i := range allAsInterface {
			resultValues = append(resultValues, *allAsInterface[i].(tree.Node).GetValue().ReflectValue())
			resultTags = append(resultTags, allAsInterface[i].(*ValueNode).tags)
		}
	}
	return resultValues, resultTags
}

// GetAllAddressableFields returns all values of the value tree fitting FlagIsSimpleData
// and FlagIsAddressable.
func GetAllAddressableFields(i interface{}) ([]reflect.Value, [][]reflect.StructTag) {
	return GetAllValuesWithFlags(i, FlagIsSimpleData|FlagIsAddressable)
}

// GetAllAddressableFieldsWithTag returns all values of the value tree fitting FlagIsSimpleData,
// FlagIsAddressable, and FlagHasTag.
func GetAllAddressableFieldsWithTag(i interface{}, tag reflect.StructTag) ([]reflect.Value, [][]reflect.StructTag) {
	return GetAllValuesWithFlags(i, FlagIsSimpleData|FlagIsAddressable|FlagHasTag)
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
