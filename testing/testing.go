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
	t       *testing.T
	verbose bool
}

// NewT creates a new instance of T, which wraps t.
func NewT(t *testing.T) *T {
	return &T{t: t, verbose: false}
}

// NewVerboseT creates a new verbose instance of T, which wraps t.
// A verbose T will also print messages when assertions hold.
func NewVerboseT(t *testing.T) *T {
	return &T{t: t, verbose: true}
}

// Assert performs Assert(...) on the testing.T wrapped by t.
func (t *T) Assert(condition bool, message string, arg ...interface{}) {
	Assert(condition, message, t, arg)
}

// AssertEquals compares v1 and v2 and performs Assert(...) on the testing.T wrapped by t with the
// result of this comparison.
func (t *T) AssertEquals(v1 interface{}, v2 interface{}, message string, arg ...interface{}) {
	Assert(v1 == v2, fmt.Sprintf("Expected: %v, Actual: %v (%s)", v1, v2, message), t, arg)
}

// AssertError performs Assert(...) on the testing.T wrapped by t,
// testing err for not being nil.
func (t *T) AssertError(err error, message string, arg ...interface{}) {
	Assert(nil != err, message, t, arg)
}

// AssertNoError performs Assert(...) on the testing.T wrapped by t,
// testing err for being nil.
func (t *T) AssertNoError(err error, message string, arg ...interface{}) {
	Assert(nil == err, message, t, arg)
}

// Assert tests a condition occuring during a test t for being true.
// If it is true, it will format-print message with the arguments given
// in arg. Otherwise it will add an error to t containing the same format-
// printed message.
func Assert(condition bool, message string, t *T, arg ...interface{}) {
	if !condition {
		t.t.Errorf(message, arg...)
		t.t.FailNow()
	}

	if t.verbose {
		fmt.Printf("Correct: "+message+"\n", arg...)
	}
}
