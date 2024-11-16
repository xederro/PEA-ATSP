package branchandbound

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"math"
)

// BranchAndBound struct that holds data for algorithm
type BranchAndBound struct {
	im *algo.IncidenceMatrix
}

// NewBranchAndBound is a constructor for BranchAndBound struct
func NewBranchAndBound(im *algo.IncidenceMatrix) *BranchAndBound {
	return &BranchAndBound{
		im: im.Copy(),
	}
}

// Solve function that solves the tsp
func (b *BranchAndBound) Solve() *methods.Res {
	// create priority queue and add first element to it
	minKnown := math.MaxInt64
	var minKnownInstance algo.Array[int]
	var q *PriorityQueue
	{
		tmpVal := b.im.ReduceMatrix()
		tmpArr := []*Node{{
			im:     b.im.Copy(),
			todo:   b.im.GetAdj(0),
			parent: nil,
			self:   0,
			val:    tmpVal,
		}}
		q = NewPriorityQueue(tmpArr)
	}

	// if not empty then go through all states with the lowest current value
	for !q.IsEmpty() {
		t, err := q.GetRoot()
		if err != nil {
			panic(err)
		}
		for _, v := range t.todo {
			tmp := &Node{
				im:     t.im.Copy(),
				parent: t,
				self:   v,
			}
			tmp.im.DiscardRow(t.self)
			tmp.im.DiscardCol(v)
			tmp.im.SetWeight(v, t.self, -1)
			tmp.im.SetWeight(v, 0, -1)
			tmp.todo = tmp.im.GetAdj(v)
			tmp.val = t.val + tmp.im.ReduceMatrix() + t.im.GetWeight(t.self, v)
			if len(tmp.todo) != 0 {
				if tmp.val < minKnown {
					q.Insert(tmp)
				}
			} else {
				tmpKnownInstance := b.calc(tmp)
				if t.val < minKnown {
					minKnown = t.val
					minKnownInstance = tmpKnownInstance
					q.Remove(minKnown)
				}
			}
		}
	}

	return &methods.Res{
		Value: minKnown,
		Route: minKnownInstance.Reverse(),
	}
}

// calc is a function that calculates the current value for an BNB instance
func (b *BranchAndBound) calc(a *Node) algo.Array[int] {
	in := algo.NewArray[int](0)
	root := a
	for root.parent != nil {
		root = root.parent
		in = append(in, root.self)
	}
	in = append([]int{a.self}, in...)

	return in
}
