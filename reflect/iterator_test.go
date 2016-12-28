package reflect

import (
	"fmt"
	"testing"

	"github.com/andreasstrack/datastructures/tree"
	"github.com/andreasstrack/util"
)

func TestSimpleValueTreeGeneration(t *testing.T) {
	s := &A{AI: 1}
	ni, err := NewValueIterator(s, 0)
	if err != nil {
		t.Error(err.Error())
	}
	for ni.HasNext() {
		n := ni.Next().(tree.Node)
		fmt.Printf("Next: %s\n", tree.String(n))
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

// 	if vn.size() != 11 {
// 		t.Errorf("size of value tree (%d) is not 11.", vn.size())
// )

func TestIterationWithTag(t *testing.T) {
	// tt := T.NewT(t)
	s := *newBc()
	// flags := FlagHasTag
	flags := util.FlagNone
	root := newValueNodeFromInterface(s, nil)
	fmt.Printf("Constructing iterator from root '%s'\n", root)
	it := tree.NewValidatedNodeIterator(
		root,
		func(n tree.Node) tree.ChildIterator {
			nvci := newValueChildIterator(flags)
			nvci.Init(n)
			return nvci
		},
		tree.BreadthFirst,
		NewNodeValidator(flags))

	for it.HasNext() {
		fmt.Printf("About to call Next()...\n")
		nodeInterface := it.Next()
		if nodeInterface == nil {
			t.Errorf("Next: nil\n")
			return
		}
		node := it.Next().(tree.Node)
		fmt.Printf("\n-----\nNext: %s\n-----\n", node)
		// vn := node.(*ValueNode)
		// if vn.structField.Tag == "" {
		// 	t.Errorf("%s does not have a tag.", node)
		// }
	}
}

// func TestTagIteration(t *testing.T) {
// 	// tt := T.NewT(t)
// 	s := newAb()
// 	flags := FlagHasTag
// 	root := newValueNodeFromInterface(s, nil)
// 	fmt.Printf("Constructing iterator from root '%s'\n", root)
// 	it := tree.NewNodeIterator(
// 		root,
// 		func(n tree.Node) tree.ChildIterator {
// 			nvci := newValueChildIterator(flags)
// 			nvci.Init(n)
// 			return nvci
// 		},
// 		tree.BreadthFirst)
// 	for it.HasNext() {
// 		fmt.Printf("Next: %s\n", it.Next())
// 	}
// }
