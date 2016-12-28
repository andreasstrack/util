package util

// Flags represents a set of boolean flags as a bit mask.
type Flags uint64

const (
	FlagNone Flags = 0
)

// HasFlag return true
func (fs Flags) HasFlag(f Flags) bool {
	return fs&f > 0
}
