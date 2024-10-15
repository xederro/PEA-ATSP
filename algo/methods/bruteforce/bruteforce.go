package bruteforce

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"math"
)

type Bruteforce struct {
	im *algo.IncidenceMatrix
}

func NewBruteforce(im *algo.IncidenceMatrix) *Bruteforce {
	return &Bruteforce{
		im: im.Copy(),
	}
}

func (b *Bruteforce) Solve() *methods.Res {
	//https://www.quickperm.org/
	a := algo.Array[int](b.im.GetNodes())
	p := algo.NewArray[int](b.im.Len() + 1).PopulateWithCounting()
	minKnown := b.calc(a)
	minKnownInstance := algo.Array[int](b.im.GetNodes())
	for i := 1; i < b.im.Len(); {
		p[i]--
		j := 0
		if i%2 == 1 {
			j = p[i]
		}
		a.Swap(i, j)
		i = 1
		for p[i] == 0 {
			p[i] = i
			i++
		}

		c := b.calc(a)
		if minKnown > c {
			minKnown = c
			copy(minKnownInstance, a)
		}
	}

	return &methods.Res{
		Value: minKnown,
		Route: minKnownInstance,
	}
}

func (b *Bruteforce) calc(a algo.Array[int]) int {
	count := b.im.GetWeight(a[0], a[len(a)-1])
	for i := 1; i < len(a); i++ {
		curr := b.im.GetWeight(a[i], a[i-1])
		if curr == -1 {
			return math.MaxInt64
		}
		count += curr
	}
	return count
}
