package reflect

import (
	"testing"

	"github.com/andreasstrack/data"
	"github.com/andreasstrack/data/tree"
	"github.com/andreasstrack/util"
	T "github.com/andreasstrack/util/testing"
)

func TestSimpleValueIteration(t *testing.T) {
	tt := T.NewT(t)
	s := newAb()
	ni, err := NewValueIterator(s, 0, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %v", s)

	for ni.HasNext() {
		n := ni.Next()
		tt.Assert(n != nil, "next node: %s", n)
	}
}

// func TestValueTreeGeneration(t *testing.T) {
// 	s := newAbbc()
// 	vn, err := buildValueTree(s, 0)

// 	if err != nil {
// 		t.Error(err.Error())
// 		return
// 	} else if vn == nil {
// 		t.Errorf("value tree is nil")
// 		return
// 	}

// 	fmt.Printf("%s\n", vn.String())

//	if vn.size() != 11 {
// 		t.Errorf("size of value tree (%d) is not 11.", vn.size())
// )

func TestBreadthFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	s := *newAbii()
	flags := util.FlagNone
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator(%s,%s)", s, flags)
	var expectedNodeValues = [...]int64{3, 4, 1, 2, 6, 5}
	i := 0
	for it.HasNext() {
		nextValue := it.Next().(tree.Node).GetValue()
		if nextValue.IsInt() {
			tt.AssertEquals(expectedNodeValues[i], nextValue.Int(), "next Int value")
			i++
		}
	}
	tt.AssertEquals(len(expectedNodeValues), i, "number of Ints")
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}

func TestDepthFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	s := *newAbii()
	flags := util.FlagNone
	it, err := NewValueIterator(s, flags, tree.DepthFirst)
	tt.AssertNoError(err, "NewValueIterator(%s,%s)", s, flags)
	var expectedNodeValues = [...]int64{1, 5, 6, 2, 3, 4}
	i := 0
	for i < len(expectedNodeValues) && it.HasNext() {
		nextValue := it.Next().(tree.Node).GetValue()
		if nextValue.IsInt() {
			tt.AssertEquals(expectedNodeValues[i], nextValue.Int(), "next Int value")
			i++
		}
	}
	tt.AssertEquals(len(expectedNodeValues), i, "number of Ints")
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}

func TestTraversalWithTags(t *testing.T) {
	tt := T.NewT(t)
	s := *newAbbc()
	flags := FlagHasTag
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %s", s)
	for it.HasNext() {
		next := it.Next()
		tt.Assert(next != nil, "it.Next() != nil (%s)", next)
		nextValue := next.(*ValueNode)
		tags := nextValue.tags
		tt.Assert(len(tags) > 0, "nextValue %s has tag: %s", nextValue.GetValue().Interface(), tags)
	}
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}

func TestTraversalOfSimpleData(t *testing.T) {
	tt := T.NewT(t)
	s := *newAbbc()
	flags := FlagIsSimpleData
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %s", s)
	for it.HasNext() {
		next := it.Next()
		tt.Assert(next != nil, "it.Next() != nil (%s)", next)
		nextValue := next.(*ValueNode)
		tt.Assert(data.IsSimpleData(nextValue), "nextValue is simple data: %s", nextValue.GetValue())
	}
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}

func TestTraversalOfSimpleDataWithTags(t *testing.T) {
	tt := T.NewT(t)
	s := *newAbbc()
	flags := FlagIsSimpleData | FlagHasTag
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %s", s)
	for it.HasNext() {
		next := it.Next()
		tt.Assert(next != nil, "it.Next() != nil (%s)", next)
		nextValue := next.(*ValueNode)
		tt.Assert(data.IsSimpleData(nextValue), "nextValue is simple data: %s", nextValue.GetValue())
		tags := nextValue.tags
		tt.Assert(len(tags) > 0, "nextValue %s has tag: %s", nextValue.GetValue().Interface(), tags)
	}
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}
