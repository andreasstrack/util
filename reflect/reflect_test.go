package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andreasstrack/util"
	"github.com/andreasstrack/util/reflect/testData"

	T "github.com/andreasstrack/util/testing"
)

func TestGetAllValues(t *testing.T) {
	tt := T.NewT(t)
	abbc := testData.NewAbbc()
	values, tags := GetAllValuesWithFlags(abbc, FlagInheritTags)
	tt.AssertEquals(19, len(values), "Number of values: %v", values)
	tt.AssertEquals(19, len(values), "Number of tag lists: %v", tags)
	tt.Assert(len(tags) >= 18, "Length of tags result.")
	tt.Assert(len(tags[6]) >= 1, "Length of 6th tag result.")
	tt.Assert(len(tags[18]) >= 2, "Length of 18th tag result.")
	tt.AssertEquals("in", fmt.Sprintf("%s", tags[6][0]), "Tag 6.0")
	tt.AssertEquals("out", fmt.Sprintf("%s", tags[18][1]), "Tag 18.1")
}

func TestGetAllSimpleDataFields(t *testing.T) {
	tt := T.NewT(t)
	b := testData.NewBifs()

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
	b := testData.NewBifs()

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
	b := testData.NewBifs()

	allFields, allTags := GetAllValuesWithFlags(b, FlagHasTag)
	tt.AssertEquals(5, len(allFields), "len(allFields)")
	tt.AssertEquals(len(allFields), len(allTags), "len(fields) = len(tags)")
	tt.Assert(allTags != nil, "allFields: %s", allFields)
	tt.Assert(allTags != nil, "allTags: %s", allTags)
	tt.AssertEquals(len(allTags), len(allFields), "Every field shall have a tag.")
	for i := range allTags {
		tt.Assert(len(allTags[0]) > 0, "Check: tag %d not empty.", i)
	}
}

func TestFlagInheritTags(t *testing.T) {
	tt := T.NewT(t)
	s := testData.NewAbbc()
	{
		values, tags := GetAllValuesWithFlags(s, util.FlagNone)
		tt.AssertEquals(len(values), len(tags), "Check: Have all tags from abbc:\n%s\n", tags)
		tt.Assert(len(tags) > 10 && len(tags[10]) == 0, "Check: Have no tags in index 10 abbc:\n%s\n", tags)
		tt.Assert(len(tags) > 18 && len(tags[18]) == 0, "Check: Have no tags in index 18 abbc:\n%s\n", tags)
	}

	{
		values, tags := GetAllValuesWithFlags(s, FlagInheritTags)
		tt.AssertEquals(len(values), len(tags), "Check: Have all tags from abbc:\n%s\n", tags)
		tt.Assert(len(tags) > 10 && len(tags[10]) > 0, "Check: Have tags in index 10 abbc:\n%s\n", tags)
		tt.Assert(len(tags) > 18 && len(tags[18]) > 1, "Check: Have tags in index 18 abbc:\n%s\n", tags)
		tt.Assert(tags[10][0] == "out", "Check: Have 'out' tag at index 10.0.")
		tt.Assert(tags[18][1] == "out", "Check: Have 'out' tag at index 18.1.")
	}
}

func TestFlagIncludeCannotInterface(t *testing.T) {
	tt := T.NewT(t)
	s := testData.NewCannotInterface()
	{
		values, _ := GetAllValuesWithFlags(s, FlagIsSimpleData)
		tt.AssertEquals(1, len(values), "Amount of values that can interface %s: %s\n", s, values)
		for i := range values {
			tt.Assert(values[i].CanInterface(), "Check if value %d can interface: %s\n", i, values[i])
		}
	}

	{
		values, _ := GetAllValuesWithFlags(s, FlagIsSimpleData|FlagIncludeCannotInterface)
		tt.AssertEquals(2, len(values), "Amount of values that can or cannot interface in %s: %s\n", s, values)
	}
}
