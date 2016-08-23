package reflect

import (
	"fmt"

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
	// FlagHasTag - Find only values with a tag.
	FlagHasTag
)

type valueIterator struct {
	flags   util.Flags
	tree    *valueNode
	current *valueNode
	nextNodeStrategy
}

// NewValueIterator generates an iterator returning values of i
// as specified by the flags.
func NewValueIterator(i interface{}, flags util.Flags, searchStrategy util.SearchStrategy) (patterns.Iterator, error) {
	var err error
	vi := valueIterator{flags: flags, tree: nil, current: nil}
	if vi.tree, err = buildValueTree(i, flags); err != nil {
		return nil, fmt.Errorf("could not generate value tree for iterator: %s", err.Error())
	}
	vi.current = vi.tree
	switch searchStrategy {
	case util.DepthFirstSearch:
		vi.nextNodeStrategy = nextNodeDepthFirstStrategy{}
	case util.BreadtFirstSearch:
		vi.nextNodeStrategy = nextNodeBreadthFirstStrategy{}
	default:
		return nil, fmt.Errorf("invalid search strategy: %d", searchStrategy)
	}
	return &vi, nil
}

func (vi *valueIterator) Next() interface{} {
	if !vi.HasNext() {
		return nil
	}

	result := vi.current.value.Interface()
	vi.current = vi.nextNodeStrategy.nextNode(vi.current)
	return result
}

func (vi *valueIterator) HasNext() bool {
	return vi.current != nil
}
