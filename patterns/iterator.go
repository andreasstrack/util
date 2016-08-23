package patterns

// Iterator is the interface for the iterate package.
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// GetAll returns all the values that the Iterator i
// (still) has.
func GetAll(i Iterator) []interface{} {
	var result []interface{}
	for i.HasNext() {
		result = append(result, i.Next())
	}
	return result
}
