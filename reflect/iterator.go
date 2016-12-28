package reflect

import (
	"reflect"

	"fmt"
	"reflect"

	"github.com/andreasstrack/datastructures"
	"github.com/andreasstrack/datastructures/tree"
	"github.com/andreasstrack/util"
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
=======
	// FlagHasTag - Find only values with a tag.
	FlagHasTag util.Flags = 1 << iota
)

// TODO: Optimize to store only parent and child index
// for iterator performance? Or have both?
type ValueNode struct {
	parent *ValueNode
	reflect.Value
	structField reflect.StructField
}

func (vn *ValueNode) String() string {
	return fmt.Sprintf("%s", vn.Value)
}

func newValueNodeFromInterface(i interface{}, parent *ValueNode) *ValueNode {
	return newValueNode(parent, reflect.ValueOf(i), nil)
}

func newValueNode(parent *ValueNode, v reflect.Value, sf *reflect.StructField) *ValueNode {
	vn := &ValueNode{}
	vn.parent = parent
	vn.Value = v
	if sf != nil {
		vn.structField = *sf
	}
	return vn
}

func (vn *ValueNode) ReflectValue() *reflect.Value {
	return &vn.Value
}

func (vn *ValueNode) GetValue() datastructures.Value {
	return vn
}

func (vn *ValueNode) GetChildren() []tree.Node {
	children := make([]tree.Node, 0)
	ev := vn.Value
	if ev.Kind() == reflect.Ptr {
		ev = ev.Elem()
	}
	if ev.Kind() != reflect.Struct {
		return children
	}
	for i := 0; i < ev.NumField(); i++ {
		fv := ev.Field(i)
		sf := ev.Type().Field(i)
		children = append(children, newValueNode(vn, fv, &sf))
	}
	return children
}

func (vn *ValueNode) Add(child tree.Node) error {
	return fmt.Errorf("cannot add child node to ValueNode")
}

func (vn *ValueNode) Insert(child tree.Node, index int) error {
	return fmt.Errorf("cannot insert child node to ValueNode")
}

func (vn *ValueNode) Remove(index int) error {
	return fmt.Errorf("cannot remove child node from ValueNode")
}

func (vn *ValueNode) GetParent() tree.Node {
	return vn.parent
}

func (vn *ValueNode) SetParent(n tree.Node) error {
	return fmt.Errorf("cannot set parent for ValueNode")
}

type valueNodeValidator struct {
	flags util.Flags
}

func NewNodeValidator(flags util.Flags) tree.NodeValidator {
	return &valueNodeValidator{flags}
}

func (vnv valueNodeValidator) IsValid(n tree.Node) bool {
	vn := n.(*ValueNode)
	v := vn.Value
	sf := vn.structField
	valid := true
	valid = valid && (!vnv.flags.HasFlag(FlagHasTag) || sf.Tag != "")
	if vnv.flags.HasFlag(FlagHasTag) {
		fmt.Printf("%s does ", v)
		if !valid {
			fmt.Printf("NOT ")
		}
		fmt.Printf("have a tag.\n")
	}
	return valid
}

// TODO: Store parent, num children and current child index
//       for not opening children unnecessarily?
type valueChildIterator struct {
	children  []tree.Node
	nextIndex int
	next      tree.Node
	flags     util.Flags
}

func newValueChildIterator(flags util.Flags) *valueChildIterator {
	return &valueChildIterator{flags: flags}
}

func (vci *valueChildIterator) Init(n tree.Node) {
	fmt.Printf("vci.Init(%s)\n", n)
	vci.children = n.GetChildren()
	vci.nextIndex = -1
	vci.next = vci.getNext()
}

func (vci *valueChildIterator) getNext() tree.Node {
	vci.nextIndex++
	l := len(vci.children)
	if vci.nextIndex >= l {
		return nil
	}
	n := vci.children[vci.nextIndex].(*ValueNode)
	fmt.Printf("vci.getNext(): %s\n", n)
	return n
}

func (vci *valueChildIterator) Next() interface{} {
	result := vci.next
	vci.next = vci.getNext()
	fmt.Printf("vci.Next(): %s (next: %s)\n", result, vci.next)
	return result
}

func (vci *valueChildIterator) HasNext() bool {
	return vci.next != nil
}
