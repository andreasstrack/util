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

// type valueBuildingChildIterator struct {
// 	parent     tree.Node
// 	next       tree.Node
// 	fieldIndex int
// 	v          reflect.Value
// }

func interfaceToValueNode(i interface{}) tree.Node {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	result := newValueNode(nil, v, nil, util.FlagNone)
	return result
}

// func NewValueBuildingChildIterator(n tree.Node) tree.ChildIterator {
// 	fmt.Printf("NewValueBuildingChildIterator( %s )\n", tree.String(n))
// 	ci := &valueBuildingChildIterator{}
// 	ci.Init(n)
// 	return ci
// }

// func (vbci *valueBuildingChildIterator) Init(n tree.Node) {
// 	fmt.Printf("vbci.Init(%s, value: %s, interface: %s)\n", n, n.GetValue(), n.GetValue().Interface())
// 	v := *n.GetValue().ReflectValue()
// 	if v.Kind() == reflect.Ptr {
// 		v = v.Elem()
// 	}
// 	if v.Kind() != reflect.Struct {
// 		vbci.parent = nil
// 		vbci.next = nil
// 		return
// 	}
// 	vbci.parent = n
// 	vbci.v = v
// 	vbci.fieldIndex = 0
// 	vbci.getNext()
// }

// func (vbci *valueBuildingChildIterator) getNext() {
// 	if vbci.parent == nil {
// 		vbci.next = nil
// 		return
// 	}
// 	if vbci.v.Kind() != reflect.Struct {
// 		vbci.next = nil
// 		return
// 	}
// 	if vbci.fieldIndex >= vbci.v.NumField() {
// 		vbci.next = nil
// 		return
// 	}
// 	v := vbci.v.Field(vbci.fieldIndex)
// 	sf := vbci.v.Type().Field(vbci.fieldIndex)
// 	vbci.next = newValueNode(vbci.parent, v, &sf)
// }

// func (vbci *valueBuildingChildIterator) HasNext() bool {
// 	return vbci.next != nil
// }

// func (vbci *valueBuildingChildIterator) Next() interface{} {
// 	result := vbci.next
// 	vbci.getNext()
// 	return result
// }

// TODO: Optimize to store only parent and child index
// for iterator performance? Or have both?
type ValueNode struct {
	tree.ValueNode
	structField         reflect.StructField
	tags                []reflect.StructTag
	childrenInitialized bool
}

func (vn *ValueNode) String() string {
	return fmt.Sprintf("%s", vn.ValueNode.String())
}

func newValueNode(parent tree.Node, v reflect.Value, tag *reflect.StructTag, flags util.Flags) *ValueNode {
	vn := &ValueNode{ValueNode: *tree.NewValueNode(v), childrenInitialized: false, tags: make([]reflect.StructTag, 0)}

	if flags.HasFlag(FlagInheritTags) && parent != nil {
		pnv := parent.(*ValueNode)
		vn.tags = append(vn.tags, pnv.tags...)

	}
	if tag != nil && *tag != "" {
		vn.tags = append(vn.tags, *tag)
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
	valid = valid && (!vnv.flags.HasFlag(FlagHasTag) || len(vn.tags) > 0)
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
	child := newValueNode(vci.parent, fv, &sf.Tag, vci.flags)
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
