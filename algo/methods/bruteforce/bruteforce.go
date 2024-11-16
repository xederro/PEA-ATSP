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
	// get every permutation using quick perm algorithm https://www.quickperm.org/
	a := algo.Array[int](b.im.GetNodes())[1:]
	p := algo.NewArray[int](b.im.Len()).PopulateWithCounting()
	minKnown := b.calc(a)
	minKnownInstance := algo.NewArray[int](b.im.Len() - 1)
	copy(minKnownInstance, a)
	for i := 1; i < b.im.Len()-1; {
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

		// calc value for current permutation
		c := b.calc(a)
		if minKnown > c {
			minKnown = c
			copy(minKnownInstance, a)
		}
	}

	minKnownInstance = append(minKnownInstance, 0)
	return &methods.Res{
		Value: minKnown,
		Route: minKnownInstance.Reverse(),
	}
}

// calc is a function that calculates the current value for an BF instance
func (b *Bruteforce) calc(a algo.Array[int]) int {
	count := b.im.GetWeight(a[0], 0) + b.im.GetWeight(0, a[len(a)-1])
	for i := 1; i < len(a); i++ {
		curr := b.im.GetWeight(a[i], a[i-1])
		if curr == -1 {
			return math.MaxInt64
		}
		count += curr
	}
	return count
}
