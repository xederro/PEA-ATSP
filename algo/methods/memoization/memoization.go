package memoization

import (
	"github.com/xederro/PEA-ATSP/algo"
	"github.com/xederro/PEA-ATSP/algo/methods"
	"math"
)

// Memoization is a struct that represents the Memoization method
type Memoization struct {
	im     *algo.IncidenceMatrix
	memo   []algo.Array[int]
	parent []algo.Array[int]
}

// NewMemoization is a function that creates a new Memoization instance
func NewMemoization(im *algo.IncidenceMatrix) *Memoization {
	// Create memoization table and parent table
	memo := make([]algo.Array[int], 1<<im.Len())
	parent := make([]algo.Array[int], 1<<im.Len())
	for i := range memo {
		// Create a new array of size im.Len() and populate it with math.MaxInt32
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

// Solve is a method that solves the ATSP problem using the Memoization method
func (m *Memoization) Solve() *methods.Res {
	// Calculate the minimum cost to complete the tour by returning to the starting node
	n := m.im.Len()
	subsetCount := 1 << n

	for mask := 1; mask < subsetCount; mask++ {
		// Iterate over all possible subsets of the nodes
		for i := 0; i < n; i++ {
			if mask&(1<<i) == 0 {
				continue
			}
			// Iterate over all possible nodes in the subset
			for j := 0; j < n; j++ {
				if j == i || mask&(1<<j) == 0 {
					continue
				}

				prevMask := mask ^ (1 << i)
				// try to find the minimum cost to reach node i from the previous subset
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

	// Calculate the minimum cost to complete the tour by returning to the starting node
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
