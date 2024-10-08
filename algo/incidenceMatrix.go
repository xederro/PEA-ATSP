package algo

import (
	"fmt"
	"math/rand"
)

// IncidenceMatrix is a representation of a graph using an incidence matrix
type IncidenceMatrix []Array[int]

// NewIncidenceMatrix creates a new IncidenceMatrix
func NewIncidenceMatrix(n int) *IncidenceMatrix {
	am := make([]Array[int], n)
	for i := 0; i < len(am); i++ {
		am[i] = NewArray[int](n).Populate()
	}

	return (*IncidenceMatrix)(&am)
}

// BuildIncidenceMatrixFromEdges creates a new IncidenceMatrix from a map of edges
func BuildIncidenceMatrixFromEdges(n int, e map[[2]int]*int) *IncidenceMatrix {
	graph := NewIncidenceMatrix(n)

	for k, v := range e {
		graph.AddEdge(k[0], k[1], *v)
	}

	return graph
}

// Generate generates a graph with a given density
func (m *IncidenceMatrix) Generate() *IncidenceMatrix {
	// calculate how many edges are needed
	n := m.Len()
	needed := n * (n - 1)
	countEdges := 0
	//  add edges
	for countEdges < needed {
		u := rand.Intn(n)
		v := rand.Intn(n)
		if u != v && !m.Exist(u, v) {
			m.AddEdge(u, v, rand.Intn(1<<12-1)+1)
			countEdges++
		}
	}

	return m
}

// AddEdge adds an edge to the graph
func (m *IncidenceMatrix) AddEdge(u, v, w int) *IncidenceMatrix {
	if u >= m.Len() || v >= m.Len() {
		panic("values too big")
	}

	(*m)[u][v] = w

	return m
}

// Exist checks if an edge exists
func (m *IncidenceMatrix) Exist(u, v int) bool {
	if u == v {
		return false
	}
	return (*m)[u][v] >= 0
}

// GetEdges returns a map of edges
func (m *IncidenceMatrix) GetEdges() map[[2]int]*int {
	e := map[[2]int]*int{}
	for i := range m.Len() {
		for j := range m.Len() {
			if i != j && (*m)[i][j] >= 0 {
				e[[2]int{i, j}] = &(*m)[i][j]
			}
		}
	}
	return e
}

// GetNodes returns a list of nodes
func (m *IncidenceMatrix) GetNodes() []int {
	return NewArray[int](m.Len())
}

// GetAdj returns a list of adjacent nodes
func (m *IncidenceMatrix) GetAdj(u int) []int {
	adj := []int{}

	for i := range m.Len() {
		if u != i && (*m)[u][i] >= 0 {
			adj = append(adj, i)
		}
	}

	return adj
}

// GetWeight returns the weight of an edge
func (m *IncidenceMatrix) GetWeight(u int, v int) int {
	return (*m)[u][v]
}

// Len returns the length of the graph
func (m *IncidenceMatrix) Len() int {
	return len(*m)
}

// Stringify returns a string representation of the graph
func (m *IncidenceMatrix) Stringify() string {
	str := ""

	for i := range m.Len() {
		str += fmt.Sprintf("%5d[", i)
		for j := range m.Len() {
			str += fmt.Sprintf("%5d ", (*m)[i][j])
		}
		str += "]\n"
	}
	return str
}

// Copy returns a copy of the graph
func (m *IncidenceMatrix) Copy() *IncidenceMatrix {
	al := make([]Array[int], m.Len())
	for i := 0; i < len(al); i++ {
		al[i] = NewArray[int](m.Len())
		copy(al[i], (*m)[i])
	}

	return (*IncidenceMatrix)(&al)
}
