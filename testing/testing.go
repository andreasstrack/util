// Package testing provides some convenience functionality
// for testing.
package testing

import (
	"fmt"
	"testing"
)

// Assert tests a condition occuring during a test t for being true.
// If it is true, it will format-print message with the arguments given
// in arg. Otherwise it will add an error to t containing the same format-
// printed message.
func Assert(condition bool, message string, t *testing.T, arg ...interface{}) {
	if condition {
		if len(arg) > 0 {
			fmt.Printf("Correct: "+message+"\n", arg)
		} else {
			fmt.Printf("Correct: " + message + "\n")
		}
	} else {
		t.Errorf(message, arg)
	}
}
