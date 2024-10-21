package branchandbound

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"math"
)

// BranchAndBound struct that holds data for algorithm
type BranchAndBound struct {
	im    *algo.IncidenceMatrix
	lower int
	upper int
}

// NewBranchAndBound is a constructor for BranchAndBound struct
func NewBranchAndBound(im *algo.IncidenceMatrix) *BranchAndBound {
	return &BranchAndBound{
		im:    im.Copy(),
		lower: 0,
		upper: math.MaxInt64,
	}
}

// Solve function that solves the tsp
func (b *BranchAndBound) Solve() *methods.Res {
	// create priority queue and add first element to it
	var q *PriorityQueue
	{
		tmpIm := b.im.Copy()
		tmpVal := tmpIm.ReduceMatrix()
		b.lower = tmpVal
		tmpArr := []*Node{{
			im:     tmpIm.Copy(),
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
		if t.val <= b.upper {
			if len(t.todo) != 0 {
				for _, v := range t.todo {
					tmp := &Node{
						im:     t.im.Copy(),
						parent: t,
						self:   v,
					}
					tmp.im.DiscardRow(t.self)
					tmp.im.DiscardCol(v)
					tmp.im.SetWeight(v, t.self, -1)
					par := t
					for par.parent != nil {
						par = par.parent
					}
					tmp.im.SetWeight(v, par.self, -1)
					tmp.todo = tmp.im.GetAdj(v)
					tmp.val = t.val + tmp.im.ReduceMatrix() + t.im.GetWeight(t.self, v)
					q.Insert(tmp)
				}
			} else {
				// if found the closest value then return
				minKnown, minKnownInstance := b.calc(t)
				return &methods.Res{
					Value: minKnown,
					Route: minKnownInstance.Reverse(),
				}
			}
		}
	}
	return nil
}

// calc is a function that calculates the current value for an BNB instance
func (b *BranchAndBound) calc(a *Node) (int, algo.Array[int]) {
	count := 0
	in := algo.NewArray[int](0)
	root := a
	for root.parent != nil {
		count += b.im.GetWeight(root.parent.self, root.self)
		root = root.parent
		in = append(in, root.self)
	}
	count += b.im.GetWeight(a.self, root.self)
	in = append([]int{a.self}, in...)

	return count, in
}
