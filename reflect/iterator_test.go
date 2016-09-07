package reflect

import (
	"fmt"
	"testing"

	"github.com/andreasstrack/datastructures/tree"
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
// 	}
// }
