// Package reflect provides utility functions for reflecting
// on go types that add further functionality to the built-in
// reflect package.
package reflect

import (
	"fmt"
	"reflect"
)

// GetAllFields returns:
// - all fields if i represents a struct
// - the value of i if i does not represent a struct
func GetAllFields(i interface{}) []reflect.Value {
	iv := GetElementValue(i)

	var result []reflect.Value
	if iv.Kind() != reflect.Struct {
		result = append(result, iv)
		return result
	}

	for i := 0; i < iv.NumField(); i++ {
		allFields := GetAllFields(iv.Field(i).Interface())
		result = append(result, allFields...)
	}

	return result
}

// GetAllAddressableFields returns:
// - all addressable fields if i represents a struct
// - the value of i if i does not represent a struct and
//   its value is addressable
// TODO: Implement using GetAllFields
func GetAllAddressableFields(i interface{}) []reflect.Value {
	var result []reflect.Value

	iv := GetElementValue(i)
	if !iv.CanAddr() {
		return result
	}

	if iv.Kind() != reflect.Struct {
		result = append(result, iv)
		return result
	}

	for i := 0; i < iv.NumField(); i++ {
		if iv.Field(i).CanAddr() {
			allFields := GetAllAddressableFields(iv.Field(i).Addr().Interface())
			result = append(result, allFields...)
		}
	}

	return result
}

// GetFieldsWithTag returns all fields of i (cf. GetAllFields),
// which are tagged with tag.
func GetFieldsWithTag(i interface{}, tag reflect.StructTag) []reflect.Value {
	var result []reflect.Value

	iv := GetElementValue(i)
	it := GetElementType(i)

	if iv.Kind() != reflect.Struct {
		return result
	}

	for j := 0; j < iv.NumField(); j++ {
		f := it.Field(j)

		if f.Tag == tag {
			result = append(result, iv.Field(j))
		}
	}

	return result
}

// GetAddressableFieldsWithTag returns all addressable fields of i (cf. GetAllAddressableFields),
// which are tagged with tag.
func GetAddressableFieldsWithTag(i interface{}, tag reflect.StructTag) []reflect.Value {
	var result []reflect.Value

	iv := GetElementValue(i)
	it := GetElementType(i)

	if iv.Kind() != reflect.Struct || !iv.CanAddr() {
		fmt.Printf("No result for: %v (kind: %s)\n", iv, iv.Kind())
		return result
	}

	for j := 0; j < iv.NumField(); j++ {
		f := it.Field(j)

		if f.Tag == tag && iv.Field(j).CanAddr() {
			result = append(result, iv.Field(j))
		}
	}

	return result
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

// GetFieldOfType looks up a field in i of the same type as t.
// It returns:
// - The first field in i of type t if one is found.
// - An invalid value otherwise.
func GetFieldOfType(i interface{}, t interface{}) reflect.Value {
	v := GetElementValue(i)
	tv := GetElementValue(t)
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		if f.Type() == tv.Type() {
			return f
		}
	}

	return reflect.Value{}
}

// HasFieldOfType returns true if GetFieldOfType returns a valid
// valid and false otherwise.
func HasFieldOfType(i interface{}, t interface{}) bool {
	return GetFieldOfType(i, t).Kind() != reflect.Invalid
}

// CopyStruct performs a deep copy of the struct represented by
// from to the struct represented by to.
// Note that to must refer to a pointer to a struct for it to be addressable and thus setable.
func CopyStruct(from interface{}, to interface{}) error {
	fromValue := GetElementValue(from)
	toValue := GetElementValue(to)

	if fromValue.Kind() != reflect.Struct ||
		toValue.Kind() != reflect.Struct {
		return fmt.Errorf("Need to structs, got %v (from) and %v (to)", from, to)
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
