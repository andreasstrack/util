package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andreasstrack/util"
	"github.com/andreasstrack/util/reflect/testData"

	T "github.com/andreasstrack/util/testing"
)

func TestCanAddressAfterValueConversion(t *testing.T) {
	tt := T.NewT(t)
	var d interface{}
	d = testData.NewD()
	v := reflect.ValueOf(d).Elem().FieldByName("DI")
	tt.Assert(v.CanAddr(), "Check: CanAddr value: %#v", v)
	vi := v.Addr().Interface()
	vvi := reflect.ValueOf(vi).Elem()
	tt.Assert(vvi.CanAddr(), "Check: CanAddr value of interface of value of pointer: %#v", vvi)
}

func TestGetAllValues(t *testing.T) {
	tt := T.NewT(t)
	abbc := testData.NewAbbc()
	values := GetAllValues(abbc, util.FlagNone)
	tt.AssertEquals(19, len(values), "Number of values: %v", values)
	var tags []string
	for i := range values {
		tags = append(tags, fmt.Sprintf("%s", values[i].Tag))
	}
	tt.AssertEquals("in", fmt.Sprintf("%s", values[6].Tag), "Tag 6j")
	tt.AssertEquals("out", fmt.Sprintf("%s", values[16].Tag), "Tag 16 (of %#v)", tags)
}

func TestGetAllSimpleDataFields(t *testing.T) {
	tt := T.NewT(t)
	b := testData.NewBifs()

	allFields := GetAllValues(*b, FlagIsSimpleData)

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
	b := testData.NewBifs()

	allFields := GetAllValues(b, FlagIsAddressable|FlagIsSimpleData)

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
	b := testData.NewBifs()

	allValues := GetAllValues(b, FlagHasTag)
	tt.AssertEquals(5, len(allValues), "len(allValues)")
	tt.Assert(allValues != nil, "allValues: %s", allValues)
	for i := range allValues {
		tt.Assert(allValues[i].Tag != "", "Check: tag %d not empty (%s).", i, allValues[i].Tag)
	}
}

func TestFlagIncludeCannotInterface(t *testing.T) {
	tt := T.NewT(t)
	s := testData.NewCannotInterface()
	{
		values := GetAllValues(s, FlagIsSimpleData)
		tt.AssertEquals(1, len(values), "Amount of values that can interface %s: %s\n", s, values)
		for i := range values {
			tt.Assert(values[i].CanInterface(), "Check if value %d can interface: %s\n", i, values[i])
		}
	}

	{
		values := GetAllValues(s, FlagIsSimpleData|FlagIncludeCannotInterface)
		tt.AssertEquals(2, len(values), "Amount of values that can or cannot interface in %s: %s\n", s, values)
	}
}
