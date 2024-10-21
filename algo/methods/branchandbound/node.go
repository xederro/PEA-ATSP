package branchandbound

import "github.com/xederro/PEA-ATSP/algo"

// Node struct that holds data for current iteration of the BNB
type Node struct {
	val    int
	self   int
	im     *algo.IncidenceMatrix
	todo   algo.Array[int]
	parent *Node
}
