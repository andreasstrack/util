// Package testing provides some convenience functionality
// for testing.
package testing

import (
	"fmt"
	"testing"
)

// T is a wrapper for testing.T, providing
// assertion functions.
type T struct {
	t *testing.T
}

// NewT creates a new instance of T, which wraps t.
func NewT(t *testing.T) *T {
	return &T{t}
}

// Assert performs Assert(...) on the testing.T wrapped by t.
func (t *T) Assert(condition bool, message string, arg ...interface{}) {
	Assert(condition, message, t.t, arg)
}

// AssertEquals compares v1 and v2 and performs Assert(...) on the testing.T wrapped by t with the
// result of this comparison.
func (t *T) AssertEquals(v1 interface{}, v2 interface{}, message string, arg ...interface{}) {
	Assert(v1 == v2, fmt.Sprintf("Expected: %v, Actual: %v (%s)", v1, v2, message), t.t, arg)
}

// AssertError performs Assert(...) on the testing.T wrapped by t,
// testing err for not being nil.
func (t *T) AssertError(err error, message string, arg ...interface{}) {
	Assert(nil != err, message, t.t, arg)
}

// AssertNoError performs Assert(...) on the testing.T wrapped by t,
// testing err for being nil.
func (t *T) AssertNoError(err error, message string, arg ...interface{}) {
	Assert(nil == err, message, t.t, arg)
}

// Assert tests a condition occuring during a test t for being true.
// If it is true, it will format-print message with the arguments given
// in arg. Otherwise it will add an error to t containing the same format-
// printed message.
func Assert(condition bool, message string, t *testing.T, arg ...interface{}) {
	if !condition {
		t.Errorf(message, arg...)
		t.FailNow()
	}

	fmt.Printf("Correct: "+message+"\n", arg...)
}
