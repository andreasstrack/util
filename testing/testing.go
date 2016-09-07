// Package testing provides some convenience functionality
// for testing.
package testing

import (
	"fmt"
	"testing"
)

type T struct {
	t *testing.T
}

func NewT(t *testing.T) *T {
	return &T{t: t}
}

func (t *T) Assert(condition bool, message string, arg ...interface{}) {
	Assert(condition, message, t.t, arg)
}

func (t *T) AssertEquals(expected, actual interface{}, description string, arg ...interface{}) {
	if expected != actual {
		t.t.Errorf("%s: expected != actual ('%s' != '%s')", fmt.Sprintf(description, arg...), expected, actual)
		t.t.FailNow()
	}

	fmt.Printf("%s: expected == actual ('%s' == '%s')\n", fmt.Sprintf(description, arg...), expected, actual)
}

func (t *T) AssertError(err error, description string, arg ...interface{}) {
	if err == nil {
		t.t.Errorf("%s: expected an error", fmt.Sprintf(description, arg...))
		t.t.FailNow()
	}

	fmt.Printf("%s: correctly received error '%s'\n", fmt.Sprintf(description, arg...), err.Error())
}

func (t *T) AssertNoError(err error, description string, arg ...interface{}) {
	if err != nil {
		t.t.Errorf("%s: unexpected error '%s'", fmt.Sprintf(description, arg...), err.Error())
		t.t.FailNow()
	}

	fmt.Printf("%s: correctly received no error\n", fmt.Sprintf(description, arg...))
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
