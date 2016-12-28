package reflect

// import (
// 	"bytes"
// 	"fmt"
// 	"reflect"

// 	"github.com/andreasstrack/util"
// )

// // TODO: Implement a tree iterator and use it for the value iterator.
// //       At each point in time, the tree iterator would only hold those
// //       parts of the tree in memory which are currently needed for iterating.

// type nextNodeStrategy interface {
// 	nextNode(vn *valueNode) *valueNode
// }

// type valueNode struct {
// 	parent       *valueNode
// 	children     []*valueNode
// 	siblingIndex int
// 	traversed    bool
// 	value        reflect.Value
// }

// func newValueNode(parent *valueNode, value reflect.Value) *valueNode {
// 	return &valueNode{parent: nil, siblingIndex: -1, traversed: false, value: value}
// }

// func (vn *valueNode) nextChild() *valueNode {
// 	for i := range vn.children {
// 		if !vn.children[i].traversed {
// 			return vn.children[i]
// 		}
// 	}

// 	return nil
// }

// func (vn *valueNode) String() string {
// 	var buffer bytes.Buffer
// 	buffer.WriteString(vn.value.String())
// 	if vn.value.Kind() != reflect.Struct && vn.value.CanInterface() {
// 		buffer.WriteString(fmt.Sprintf("(%s)", vn.value.Interface()))
// 	}
// 	if len(vn.children) == 0 {
// 		return buffer.String()
// 	}

// 	buffer.WriteString(" -> [\n")
// 	for i := range vn.children {
// 		buffer.WriteString(vn.children[i].String())
// 		buffer.WriteString("\n")
// 	}
// 	buffer.WriteString("]")
// 	return buffer.String()
// }

// func (vn *valueNode) size() int {
// 	if vn == nil {
// 		return 0
// 	}

// 	size := 1

// 	for i := range vn.children {
// 		size += vn.children[i].size()
// 	}

// 	return size
// }

// func (vn *valueNode) nextSibling() *valueNode {
// 	if vn.parent == nil {
// 		return nil
// 	}

// 	if vn.siblingIndex < 0 {
// 		return nil
// 	}

// 	if len(vn.parent.children) <= vn.siblingIndex+1 {
// 		return nil
// 	}

// 	if c := vn.parent.nextChild(); c != nil {
// 		return c
// 	}

// 	return vn.parent.nextSibling()
// }

// func buildValueTree(root interface{}, flags util.Flags) (*valueNode, error) {
// 	v := reflect.ValueOf(root)
// 	ev := GetElementValue(root)
// 	et := GetElementType(root)
// 	vn := newValueNode(nil, v)

// 	if flags.HasFlag(FlagIsAddressable) && !v.CanAddr() {
// 		return nil, nil
// 	}
// 	if flags.HasFlag(FlagIsPointer) && v.Kind() == reflect.Ptr {
// 		return nil, nil
// 	}

// 	switch {
// 	case ev.Kind() == reflect.Invalid:
// 		return nil, nil
// 	case ev.Kind() != reflect.Struct:
// 		return vn, nil
// 	}

// 	for i := 0; i < ev.NumField(); i++ {
// 		var cn *valueNode
// 		var err error
// 		if flags.HasFlag(FlagHasTag) {
// 			ft := et.Field(i)
// 			if len(ft.Tag) == 0 {
// 				continue
// 			}
// 		}
// 		if fv := ev.Field(i); fv.CanInterface() {
// 			cn, err = buildValueTree(fv.Interface(), flags)
// 			if err != nil {
// 				return nil, err
// 			} else if cn == nil {
// 				return nil, fmt.Errorf("could not generate node for value %s", fv)
// 			}
// 		} else if flags.HasFlag(FlagForceCanInterface) {
// 			return nil, fmt.Errorf("cannot interface value %s", fv)
// 		}

// 		cn.siblingIndex = i
// 		cn.parent = vn
// 		vn.children = append(vn.children, cn)
// 	}

// 	return vn, nil
// }

// type nextNodeDepthFirstStrategy struct {
// }

// func (s nextNodeDepthFirstStrategy) nextNode(vn *valueNode) *valueNode {
// 	vn.traversed = true

// 	if c := vn.nextChild(); c != nil {
// 		return c
// 	}

// 	if s := vn.nextSibling(); s != nil {
// 		return s
// 	}

// 	if vn.parent == nil {
// 		return nil
// 	}

// 	return nil
// }

// type nextNodeBreadthFirstStrategy struct {
// }

// func (s nextNodeBreadthFirstStrategy) nextNode(vn *valueNode) *valueNode {
// 	return nil
// }
