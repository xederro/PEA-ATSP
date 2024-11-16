package memoization

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"math"
)

type Memoization struct {
	im     *algo.IncidenceMatrix
	memo   []algo.Array[int]
	parent []algo.Array[int]
}

func NewMemoization(im *algo.IncidenceMatrix) *Memoization {
	memo := make([]algo.Array[int], 1<<im.Len())
	parent := make([]algo.Array[int], 1<<im.Len())
	for i := range memo {
		memo[i] = make([]int, im.Len())
		parent[i] = make([]int, im.Len())
		for j := range memo[i] {
			memo[i][j] = math.MaxInt32
			parent[i][j] = -1
		}
	}
	memo[1][0] = 0

	return &Memoization{
		im:     im.Copy(),
		memo:   memo,
		parent: parent,
	}
}

func (m *Memoization) Solve() *methods.Res {
	n := m.im.Len()
	subsetCount := 1 << n

	for mask := 1; mask < subsetCount; mask++ {
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}

			for j := 0; j < n; j++ {
				if j == i || mask&(1<<j) == 0 {
					continue
				}

				prevMask := mask ^ (1 << i)
				if m.memo[prevMask][j] != math.MaxInt32 && m.im.GetWeight(j, i) != math.MaxInt32 {
					tmp := m.memo[prevMask][j] + m.im.GetWeight(j, i)
					if tmp < m.memo[mask][i] {
						m.memo[mask][i] = tmp
						m.parent[mask][i] = j
					}
				}
			}
		}
	}

	minCount, minInstance := m.calc()
	return &methods.Res{
		Value: minCount,
		Route: minInstance.Reverse(),
	}
}

// calc is a function that calculates the current value for an BF instance
func (m *Memoization) calc() (int, algo.Array[int]) {
	// Calculate minimum cost to complete the tour by returning to the starting node
	minCost := math.MaxInt32
	mask := (1 << m.im.Len()) - 1
	end := 0
	var path algo.Array[int]
	for i := 1; i < m.im.Len(); i++ {
		if m.memo[mask][i] != math.MaxInt32 && m.im.GetWeight(i, 0) != math.MaxInt32 {
			tmp := m.memo[mask][i] + m.im.GetWeight(i, 0)
			if tmp < minCost {
				minCost = tmp
				end = i
			}
		}
	}

	last := end
	for mask != 0 {
		path = append(path, last)
		tmp := m.parent[mask][last]
		mask ^= 1 << last
		last = tmp
	}
	path[len(path)-1] = 0

	return minCost, path
}
