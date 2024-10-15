package methods

import "github.com/xederro/PEA-ATSP/algo"

type Method interface {
	Solve() *Res
}

type Res struct {
	Value int
	Route algo.Array[int]
}
