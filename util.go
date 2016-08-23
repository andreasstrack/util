// Package util contains helpful utility types and functions.
package util

type SearchStrategy uint

const (
	DepthFirstSearch SearchStrategy = iota
	BreadtFirstSearch
)
