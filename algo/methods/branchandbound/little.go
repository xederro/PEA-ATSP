package branchandbound

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"math"
)

type Little struct {
	im    *algo.IncidenceMatrix
	lower int
	upper int
}

func NewLittle(im *algo.IncidenceMatrix) *Little {
	return &Little{
		im:    im.Copy(),
		lower: 0,
		upper: math.MaxInt64,
	}
}

func (b *Little) Solve() *methods.Res {
	var q *PriorityQueue
	{
		var tmpArr []*Node
		tmpIm := b.im.Copy()
		tmpVal := tmpIm.ReduceMatrix()
		for i := range b.im.Len() {
			tmpArr = append(tmpArr, &Node{
				im:     tmpIm.Copy(),
				todo:   b.im.GetAdj(i),
				parent: nil,
				self:   i,
				val:    tmpVal,
			})
		}
		q = NewPriorityQueue(tmpArr)
	}

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

func (b *Little) calc(a *Node) (int, algo.Array[int]) {
	count := 0
	in := algo.NewArray[int](0)
	root := a
	for root.parent != nil {
		count += b.im.GetWeight(root.parent.self, root.self)
		root = root.parent
		in = append(in, root.self)
	}
	count += b.im.GetWeight(a.self, root.self)
	in = append(in, a.self)

	return count, in
}
