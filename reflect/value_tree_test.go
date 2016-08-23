package reflect

import (
	"fmt"
	"testing"
)

func TestSimpleValueTreeGeneration(t *testing.T) {
	s := &A{AI: 1}

	vn, err := buildValueTree(s, 0)

	if err != nil {
		t.Error(err.Error())
		return
	} else if vn == nil {
		t.Errorf("value tree is nil")
		return
	}

	fmt.Printf("%s\n", vn.String())

	if vn.size() != 2 {
		t.Errorf("size of value tree (%d) is not 2.", vn.size())
	}
}

func TestValueTreeGeneration(t *testing.T) {
	s := newAbbc()
	vn, err := buildValueTree(s, 0)

	if err != nil {
		t.Error(err.Error())
		return
	} else if vn == nil {
		t.Errorf("value tree is nil")
		return
	}

	fmt.Printf("%s\n", vn.String())

	if vn.size() != 11 {
		t.Errorf("size of value tree (%d) is not 11.", vn.size())
	}
}
