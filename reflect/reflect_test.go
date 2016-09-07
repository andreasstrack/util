package reflect

// import (
// 	"reflect"
// 	"testing"

// 	T "github.com/andreasstrack/util/testing"
// )

// func TestGetAllFields(t *testing.T) {
// 	b := RB{W: true, RA: RA{I: 2, F: 3.14, S: "hello"}}

// 	allFields := GetAllFields(b)

// 	T.Assert(len(allFields) == 4, "len(allFields) == 4: %v", t, allFields)

// 	T.Assert(allFields[0].Kind() == reflect.Bool, "allFields[0].Kind() == reflect.Bool: %v", t, allFields[0].Interface())
// 	T.Assert(allFields[0].Interface().(bool) == b.W, "allFields[0].Interface().(bool) == B.W", t)

// 	T.Assert(allFields[1].Kind() == reflect.Int, "allFields[1].Kind() == reflect.Int: %v", t, allFields[1].Interface())
// 	T.Assert(allFields[1].Interface().(int) == b.I, "allFields[1].Interface().(int) == B.I", t)

// 	T.Assert(allFields[2].Kind() == reflect.Float32, "allFields[2].Kind() == reflect.Float32: %v", t, allFields[2].Interface())
// 	T.Assert(allFields[2].Interface().(float32) == b.F, "allFields[2].Interface().(float32) == B.F", t)

// 	T.Assert(allFields[3].Kind() == reflect.String, "allFields[3].Kind() == reflect.String: %v", t, allFields[3].Interface())
// 	T.Assert(allFields[3].Interface().(string) == b.S, "allFields[3].Interface().(string) == B.S", t)
// }

// func TestGetAllAddressableFields(t *testing.T) {
// 	b := RB{W: true, RA: RA{I: 2, F: 3.14, S: "hello"}}

// 	allFields := GetAllAddressableFields(&b)

// 	T.Assert(len(allFields) == 4, "len(allFields) == 4: %v", t, allFields)

// 	T.Assert(allFields[0].Kind() == reflect.Bool, "allFields[0].Kind() == reflect.Bool: %v", t, allFields[0].Interface())
// 	T.Assert(allFields[0].Interface().(bool) == b.W, "allFields[0].Interface().(bool) == B.W", t)

// 	T.Assert(allFields[1].Kind() == reflect.Int, "allFields[1].Kind() == reflect.Int: %v", t, allFields[1].Interface())
// 	T.Assert(allFields[1].Interface().(int) == b.I, "allFields[1].Interface().(int) == B.I", t)

// 	T.Assert(allFields[2].Kind() == reflect.Float32, "allFields[2].Kind() == reflect.Float32: %v", t, allFields[2].Interface())
// 	T.Assert(allFields[2].Interface().(float32) == b.F, "allFields[2].Interface().(float32) == B.F", t)

// 	T.Assert(allFields[3].Kind() == reflect.String, "allFields[3].Kind() == reflect.String: %v", t, allFields[3].Interface())
// 	T.Assert(allFields[3].Interface().(string) == b.S, "allFields[3].Interface().(string) == B.S", t)

// 	for i := range allFields {
// 		T.Assert(allFields[i].CanAddr(), "allFields[i].CanAddr(): %v", t, allFields[i])
// 	}
// }

// func TestGetAddressableFieldsWithTag(t *testing.T) {
// 	b := RB{W: true, RA: RA{I: 2, F: 3.14, S: "hello"}}

// 	allFields := GetAddressableFieldsWithTag(b, "foo")
// 	T.Assert(len(allFields) == 0, "len(allFields) == 0: %v", t, allFields)

// 	allFields = GetAddressableFieldsWithTag(&b, "bar")
// 	T.Assert(len(allFields) == 1, "len(allFields) == 1: %v", t, allFields)
// }
