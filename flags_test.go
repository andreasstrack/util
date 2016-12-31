package util

import (
	"testing"

	T "github.com/andreasstrack/util/testing"
)

const (
	fOne Flags = 1 << iota
	fTwo
	fThree
	fFour
	fFive
)

func TestHasFlag(t *testing.T) {
	tt := T.NewT(t)
	var fs Flags

	fs = fs | fTwo | fThree
	T.Assert(!fs.HasFlag(fOne), "HasFlag(%#02x): %#02x\n", tt, fs, fOne)
	T.Assert(fs.HasFlag(fTwo), "HasFlag(%#02x): %#02x\n", tt, fs, fTwo)
	T.Assert(fs.HasFlag(fThree), "HasFlag(%#02x): %#02x\n", tt, fs, fThree)
	T.Assert(!fs.HasFlag(fFour), "HasFlag(%#02x): %#02x\n", tt, fs, fFour)
	T.Assert(!fs.HasFlag(fFive), "HasFlag(%#02x): %#02x\n", tt, fs, fFive)
}
