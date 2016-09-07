package reflect

import (
	"fmt"
	"reflect"

	"github.com/andreasstrack/datastructures/tree"
	"github.com/andreasstrack/util"
	"github.com/andreasstrack/util/patterns"
)

const (
	// FlagIsAddressable - Find only addressable values.
	FlagIsAddressable util.Flags = 1 << iota
	// FlagForceCanInterface - If not all values can be interfaced, return no value at all.
	FlagForceCanInterface
	// FlagIsPointer - Find only pointer values.
	FlagIsPointer
)

type valueIterator struct {
	flags util.Flags
	tree.NodeIterator
}

// NewValueIterator generates an iterator returning values of i
// as specified by the flags.
func NewValueIterator(i interface{}, flags util.Flags) (patterns.Iterator, error) {
	vi := valueIterator{flags: flags}
	vi.NodeIterator = *tree.NewNodeIterator(interfaceToValueNode(i), NewValueBuildingChildIterator, tree.BreadthFirst)
	return &vi, nil
}

type valueBuildingChildIterator struct {
	parent     tree.Node
	next       tree.Node
	fieldIndex int
	v          reflect.Value
}

func interfaceToValueNode(i interface{}) tree.Node {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return tree.NewNode(v)
}

func NewValueBuildingChildIterator(n tree.Node) tree.ChildIterator {
	fmt.Printf("NewValueBuildingChildIterator( %s )\n", tree.String(n))
	ci := &valueBuildingChildIterator{}
	ci.Init(n)
	return ci
}

func (vbci *valueBuildingChildIterator) Init(n tree.Node) {
	v := n.GetValue().Interface().(reflect.Value)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		vbci.parent = nil
		vbci.next = nil
		return
	}
	vbci.parent = n
	vbci.v = v
	vbci.fieldIndex = 0
	vbci.getNext()
}

func (vbci *valueBuildingChildIterator) getNext() {
	if vbci.parent == nil {
		vbci.next = nil
		return
	}
	if vbci.v.Kind() != reflect.Struct {
		vbci.next = nil
		return
	}
	if vbci.fieldIndex >= vbci.v.NumField() {
		vbci.next = nil
		return
	}
	vbci.next = tree.NewNode(vbci.v.Field(vbci.fieldIndex))
}

func (vbci *valueBuildingChildIterator) HasNext() bool {
	return vbci.next != nil
}

func (vbci *valueBuildingChildIterator) Next() interface{} {
	result := vbci.next
	vbci.getNext()
	return result
}
