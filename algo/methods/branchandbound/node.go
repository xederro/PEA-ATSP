package branchandbound

import "github.com/xederro/PEA-ATSP/algo"

type Node struct {
	val    int
	self   int
	im     *algo.IncidenceMatrix
	todo   algo.Array[int]
	parent *Node
}
