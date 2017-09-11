package reflect

import (
	"testing"

	"github.com/andreasstrack/data"
	"github.com/andreasstrack/data/tree"
	"github.com/andreasstrack/util"
	"github.com/andreasstrack/util/patterns"
	"github.com/andreasstrack/util/reflect/testData"
	T "github.com/andreasstrack/util/testing"
)

func TestSimpleValueIteration(t *testing.T) {
	tt := T.NewT(t)
	s := testData.NewAb()
	ni, err := NewValueIterator(s, 0, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %v", s)

	for ni.HasNext() {
		n := ni.Next()
		tt.Assert(n != nil, "next node: %s", n)
	}
}

func TestBreadthFirstTraversal(t *testing.T) {
	tt := T.NewT(t)
	s := *testData.NewAbii()
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
	s := *testData.NewAbii()
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

func TestValueTree(t *testing.T) {
	tt := T.NewT(t)
	s := *testData.NewAbbc()
	flags := util.FlagNone
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %s", s)
	tt.Assert(it.HasNext(), "Iterator has at least one value (%s).", it)
	valuesFromIterator := patterns.GetAll(it)
	root := valuesFromIterator[0]
	var valuesFromTree []interface{}
	queue := data.NewFifoQueue()
	queue.Insert(root)
	for !queue.IsEmpty() {
		n := queue.Pop().(tree.Node)
		valuesFromTree = append(valuesFromTree, n.(*ValueNode))
		children := n.GetChildren()
		for i := range children {
			queue.Insert(children[i])
		}
	}
	tt.AssertEquals(len(valuesFromIterator), len(valuesFromTree), "Amount of values from tree equals amount of values from iterator.")
	for i := range valuesFromIterator {
		tt.AssertEquals(valuesFromIterator[i].(*ValueNode), valuesFromTree[i].(*ValueNode), "Value %d from tree is equal to value from iterator.", i)
	}
}

func TestTraversalWithTags(t *testing.T) {
	tt := T.NewT(t)
	s := *testData.NewAbbc()
	flags := FlagHasTag
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %s", s)
	for it.HasNext() {
		next := it.Next()
		tt.Assert(next != nil, "it.Next() != nil (%s)", next)
		nextValue := next.(*ValueNode)
		tt.Assert(nextValue.Tag != "", "nextValue %s has tag: %s", nextValue.GetValue().Interface(), nextValue.Tag)
	}
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}

func TestTraversalOfSimpleData(t *testing.T) {
	tt := T.NewT(t)
	s := *testData.NewAbbc()
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
	s := *testData.NewAbbc()
	flags := FlagIsSimpleData | FlagHasTag
	it, err := NewValueIterator(s, flags, tree.BreadthFirst)
	tt.AssertNoError(err, "NewValueIterator for %s", s)
	for it.HasNext() {
		next := it.Next()
		tt.Assert(next != nil, "it.Next() != nil (%s)", next)
		nextValue := next.(*ValueNode)
		tt.Assert(data.IsSimpleData(nextValue), "nextValue is simple data: %s", nextValue.GetValue())
		tt.Assert(nextValue.Tag != "", "nextValue %s has tag: %s", nextValue.GetValue().Interface(), nextValue.Tag)
	}
	next := it.Next()
	tt.Assert(nil == next, "it.Next() == nil (%s)", next)
}
