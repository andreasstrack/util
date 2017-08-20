package reflect

import (
	"fmt"
	"reflect"
	"testing"

	T "github.com/andreasstrack/util/testing"
)

func TestGetAllValues(t *testing.T) {
	tt := T.NewT(t)
	abbc := newAbbc()
	values, tags := GetAllValues(abbc)
	tt.AssertEquals(19, len(values), "Number of values: %v", values)
	tt.AssertEquals(19, len(values), "Number of tag lists: %v", tags)
	tt.AssertEquals("in", fmt.Sprintf("%s", tags[6][0]), "Tag 6.0")
	tt.AssertEquals("out", fmt.Sprintf("%s", tags[18][1]), "Tag 18.1")
}

func TestGetAllFields(t *testing.T) {
	tt := T.NewT(t)
	b := newBifs()

	allFields, _ := GetAllValuesWithFlags(*b, FlagIsSimpleData)

	tt.AssertEquals(4, len(allFields), "len(allFields) for %s", b)

	tt.Assert(allFields[0].Kind() == reflect.Bool, "allFields[0].Kind() == reflect.Bool")
	tt.Assert(allFields[0].Interface().(bool) == b.B, "allFields[0].Interface().(bool) == b.B")

	tt.Assert(allFields[1].Kind() == reflect.Int, "allFields[1].Kind() == reflect.Int: %v", allFields[1].Interface())
	tt.Assert(allFields[1].Interface().(int) == b.I, "allFields[1].Interface().(int) == B.I")

	tt.Assert(allFields[2].Kind() == reflect.Float32, "allFields[2].Kind() == reflect.Float32: %v", allFields[2].Interface())
	tt.Assert(allFields[2].Interface().(float32) == b.F, "allFields[2].Interface().(float32) == B.F")

	tt.Assert(allFields[3].Kind() == reflect.String, "allFields[3].Kind() == reflect.String: %v", allFields[3].Interface())
	tt.Assert(allFields[3].Interface().(string) == b.S, "allFields[3].Interface().(string) == B.S")
}

func TestGetAllAddressableFields(t *testing.T) {
	tt := T.NewT(t)
	b := newBifs()

	allFields, _ := GetAllValuesWithFlags(b, FlagIsAddressable|FlagIsSimpleData)

	tt.AssertEquals(4, len(allFields), "len(allFields) (%s)", allFields)

	tt.AssertEquals(reflect.Bool, allFields[0].Kind(), "allFields[0].Kind()")
	tt.Assert(allFields[0].Interface().(bool) == b.B, "allFields[0].Interface().(bool) == b.B")

	tt.Assert(allFields[1].Kind() == reflect.Int, "allFields[1].Kind() == reflect.Int: %v", allFields[1].Interface())
	tt.Assert(allFields[1].Interface().(int) == b.I, "allFields[1].Interface().(int) == B.I")

	tt.Assert(allFields[2].Kind() == reflect.Float32, "allFields[2].Kind() == reflect.Float32: %v", allFields[2].Interface())
	tt.Assert(allFields[2].Interface().(float32) == b.F, "allFields[2].Interface().(float32) == B.F")

	tt.Assert(allFields[3].Kind() == reflect.String, "allFields[3].Kind() == reflect.String: %v", allFields[3].Interface())
	tt.Assert(allFields[3].Interface().(string) == b.S, "allFields[3].Interface().(string) == B.S")

	for i := range allFields {
		tt.Assert(allFields[i].CanAddr(), "allFields[i].CanAddr(): %v", allFields[i])
	}
}

func TestGetAddressableFieldsWithTag(t *testing.T) {
	tt := T.NewT(t)
	b := newBifs()

	allFields, allTags := GetAllValuesWithFlags(b, FlagIsSimpleData|FlagIsAddressable|FlagHasTag)
	tt.AssertEquals(4, len(allFields), "len(allFields)")
	tt.AssertEquals(len(allFields), len(allTags), "len(fields) = len(tags)")
	tt.Assert(allTags != nil, "allFields: %s", allFields)
	tt.Assert(allTags != nil, "allTags: %s", allTags)

	// allFields = GetAddressableFieldsWithTag(b, "bar")
	// tt.Assert(len(allFields) == 1, "len(allFields) == 1: %v", allFields)
}
