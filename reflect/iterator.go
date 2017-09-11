package reflect

import (
	"fmt"
	"reflect"

	"github.com/andreasstrack/data"
	"github.com/andreasstrack/data/tree"
	"github.com/andreasstrack/util"
	"github.com/andreasstrack/util/patterns"
)

const (
	// FlagHasTag - Find only struct fields with a tag.
	FlagHasTag util.Flags = 1 << iota
	// FlagInheritTags - Passes the tag of a struct field to its (potential)
	// children
	FlagInheritTags
	// FlagIsSimpleData - Find only values representing non-aggregate
	// (non-struct) data, e.g. int, float, string etc.
	FlagIsSimpleData
	// FlagIsAddressable - Find only addressable values.
	FlagIsAddressable
	// FlagIncludeCannotInterface - Find also values that cannot interface.
	FlagIncludeCannotInterface
)

// NewValueIterator generates an iterator returning values of i
// as specified by the flags.
func NewValueIterator(i interface{}, flags util.Flags, traversalStrategy tree.TraversalStrategy) (patterns.Iterator, error) {
	vi := *tree.NewValidatedNodeIterator(interfaceToValueNode(i),
		func(n tree.Node) tree.ChildIterator {
			ci := newValueChildIterator(flags)
			ci.Init(n)
			return ci
		},
		traversalStrategy,
		NewNodeValidator(flags))
	return &vi, nil
}

func interfaceToValueNode(i interface{}) tree.Node {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	result := newValueNode(nil, v, nil, util.FlagNone)
	return result
}

type ValueNode struct {
	tree.ValueNode
	reflect.StructField
}

func (vn *ValueNode) String() string {
	return fmt.Sprintf("%s", vn.ValueNode.String())
}

func newValueNode(parent tree.Node, v reflect.Value, structField *reflect.StructField, flags util.Flags) *ValueNode {
	vn := &ValueNode{ValueNode: *tree.NewValueNode(v)}
	if structField != nil {
		vn.StructField = *structField
	}
	if parent != nil {
		parent.Add(vn)
	}
	return vn
}

func (vn *ValueNode) ReflectValue() *reflect.Value {
	return &vn.Value
}

func (vn *ValueNode) GetValue() data.Value {
	return vn
}

type valueNodeValidator struct {
	flags util.Flags
}

func NewNodeValidator(flags util.Flags) tree.NodeValidator {
	return &valueNodeValidator{flags}
}

func (vnv valueNodeValidator) IsValid(n tree.Node) bool {
	vn := n.(*ValueNode)
	valid := true
	if !vn.CanInterface() {
		valid = vnv.flags.HasFlag(FlagIncludeCannotInterface)
	}
	valid = valid && (!vnv.flags.HasFlag(FlagHasTag) || vn.Tag != "")
	valid = valid && (!vnv.flags.HasFlag(FlagIsSimpleData) || data.IsSimpleData(n.(data.Value)))
	valid = valid && (!vnv.flags.HasFlag(FlagIsAddressable) || vn.CanAddr())
	return valid
}

type valueChildIterator struct {
	parent       tree.Node
	elementValue reflect.Value
	nextIndex    int
	next         tree.Node
	flags        util.Flags
}

func newValueChildIterator(flags util.Flags) *valueChildIterator {
	return &valueChildIterator{flags: flags}
}

func (vci *valueChildIterator) Init(n tree.Node) {
	vci.parent = n
	vci.elementValue = n.(*ValueNode).Value
	if vci.elementValue.Kind() == reflect.Ptr {
		vci.elementValue = vci.elementValue.Elem()
	}
	vci.nextIndex = 0
	vci.next = vci.getNext()
}

func (vci *valueChildIterator) getNext() tree.Node {
	if vci.elementValue.Kind() != reflect.Struct || vci.nextIndex >= vci.elementValue.NumField() {
		return nil
	}
	fv := vci.elementValue.Field(vci.nextIndex)
	sf := vci.elementValue.Type().Field(vci.nextIndex)
	child := newValueNode(vci.parent, fv, &sf, vci.flags)
	vci.nextIndex++
	return child
}

func (vci *valueChildIterator) Next() interface{} {
	result := vci.next
	vci.next = vci.getNext()
	return result
}

func (vci *valueChildIterator) HasNext() bool {
	return vci.next != nil
}
