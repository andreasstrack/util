package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func GetAllFields(i interface{}) []reflect.Value {
	iv := GetReferenceValue(i)

	result := make([]reflect.Value, 0)
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

func GetAllAddressableFields(i interface{}) []reflect.Value {
	result := make([]reflect.Value, 0)

	iv := GetReferenceValue(i)
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

func GetAddressableFieldsWithTag(i interface{}, tag reflect.StructTag) []reflect.Value {
	result := make([]reflect.Value, 0)

	iv := GetReferenceValue(i)
	it := GetReferenceType(i)

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

func GetFieldsWithTag(i interface{}, tag reflect.StructTag) []reflect.Value {
	result := make([]reflect.Value, 0)

	iv := GetReferenceValue(i)
	it := GetReferenceType(i)

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

func IsPointer(i interface{}) bool {
	return reflect.ValueOf(i).Kind() == reflect.Ptr
}

func GetReferenceValue(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr {
		return v.Elem()
	} else {
		return v
	}
}

func GetReferenceType(i interface{}) reflect.Type {
	return GetReferenceValue(i).Type()
}

func IsStruct(i interface{}) bool {
	var v reflect.Value

	if v = GetReferenceValue(i); v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}

func HasField(i interface{}, name string) bool {
	if !IsStruct(i) {
		return false
	}

	return GetReferenceValue(i).FieldByName(name).Kind() != reflect.Invalid
}

func GetFieldOfType(i interface{}, t interface{}) reflect.Value {
	v := GetReferenceValue(i)
	tv := GetReferenceValue(t)
	for j := 0; j < v.NumField(); j++ {
		f := v.Field(j)
		if f.Type() == tv.Type() {
			return f
		}
	}

	return reflect.Value{}
}

func HasFieldOfType(i interface{}, t interface{}) bool {
	return GetFieldOfType(i, t).Kind() != reflect.Invalid
}

func CopyStruct(from interface{}, to interface{}) error {
	vFrom := reflect.ValueOf(from)
	vTo := reflect.ValueOf(to)

	if vFrom.Kind() == reflect.Ptr {
		vFrom = vFrom.Elem()
	}

	if vTo.Kind() == reflect.Ptr {
		vTo = vTo.Elem()
	}

	fmt.Printf("From: %s\nTo:%s\n", vFrom.Kind().String(), vTo.Kind().String())

	nFieldsFrom := vFrom.NumField()
	nFieldsTo := vTo.NumField()

	if nFieldsFrom != nFieldsTo {
		return errors.New("Types of from and to do not match.")
	}

	for i := 0; i < nFieldsFrom; i++ {
		fFrom := vFrom.Field(i)
		fTo := vTo.Field(i)

		kFrom := fFrom.Kind()
		kTo := fTo.Kind()

		if kFrom != kTo {
			return errors.New("Fields for from and to do not match.")
		}

		switch kFrom {
		case reflect.Int:
			fTo.SetInt(fFrom.Int())
			break
		case reflect.Float32:
		case reflect.Float64:
			fTo.SetFloat(fFrom.Float())
			break
		case reflect.Complex64:
		case reflect.Complex128:
			fTo.SetComplex(fFrom.Complex())
			break
		case reflect.String:
			fTo.SetString(fFrom.String())
			break
		default:
			return errors.New("Invalid kind of field.")
		}
	}

	return nil
}
