package methods

import "github.com/xederro/PEA-ATSP/algo"

// Method is an interface that represents a method of solving the ATSP problem.
type Method interface {
	Solve() *Res
}

// Res is a struct that represents the result of a method.
type Res struct {
	Value int
	Route algo.Array[int]
}
