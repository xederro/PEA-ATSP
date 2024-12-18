package algo

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// IncidenceMatrix is a representation of a graph using an incidence matrix
type IncidenceMatrix []Array[int]

// NewIncidenceMatrixFromFile creates a new IncidenceMatrix
func NewIncidenceMatrixFromFile(path string) *IncidenceMatrix {
	// read file
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// split file into lines
	f := strings.ReplaceAll(string(file), "\r", "")
	ret := strings.Split(f, "\n")
	atoi, err := strconv.Atoi(ret[0])
	if err != nil {
		fmt.Println(err)
		return nil
	}

	r, err := regexp.Compile(`\d+|-1`)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// create new IncidenceMatrix of size atoi
	im := NewIncidenceMatrix(atoi)
	// populate IncidenceMatrix with values from file
	for y, v := range ret[1 : im.Len()+1] {
		lines := r.FindAllString(v, -1)
		for x, val := range lines {
			weight, err := strconv.Atoi(val)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			(*im)[y][x] = weight
		}
	}

	return im
}

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
	return NewArray[int](m.Len()).PopulateWithCounting()
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

// SetWeight sets weight
func (m *IncidenceMatrix) SetWeight(u int, v int, w int) {
	(*m)[u][v] = w
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

// GetMinCol returns the minimum value in a column
func (b *IncidenceMatrix) GetMinCol(col int) int {
	minVal := math.MaxInt32
	for i := range b.Len() {
		if (*b)[i][col] != -1 && (*b)[i][col] < minVal {
			minVal = (*b)[i][col]
		}
	}
	if minVal == math.MaxInt32 {
		return 0
	}
	return minVal
}

// GetMinRow returns the minimum value in a row
func (b *IncidenceMatrix) GetMinRow(row int) int {
	minVal := math.MaxInt32
	for i := range b.Len() {
		if (*b)[row][i] != -1 && (*b)[row][i] < minVal {
			minVal = (*b)[row][i]
		}
	}
	if minVal == math.MaxInt32 {
		return 0
	}
	return minVal
}

// ReduceCol reduces a column in the matrix by the minimum value in the column
func (b *IncidenceMatrix) ReduceCol(col int) int {
	m := b.GetMinCol(col)
	if m == 0 {
		return m
	}
	for i := range b.Len() {
		if (*b)[i][col] != -1 {
			(*b)[i][col] -= m
		}
	}
	return m
}

// ReduceRow reduces a row in the matrix by the minimum value in the row
func (b *IncidenceMatrix) ReduceRow(row int) int {
	m := b.GetMinRow(row)
	if m == 0 {
		return m
	}
	for i := range b.Len() {
		if (*b)[row][i] != -1 {
			(*b)[row][i] -= m
		}
	}
	return m
}

// DiscardCol sets all values in a column to -1
func (b *IncidenceMatrix) DiscardCol(col int) {
	for i := range b.Len() {
		(*b)[i][col] = -1
	}
}

// DiscardRow sets all values in a row to -1
func (b *IncidenceMatrix) DiscardRow(row int) {
	for i := range b.Len() {
		(*b)[row][i] = -1
	}
}

// ReduceMatrix reduces the matrix
func (b *IncidenceMatrix) ReduceMatrix() int {
	reduced := 0
	for i := range b.Len() {
		reduced += b.ReduceRow(i)
	}
	for i := range b.Len() {
		reduced += b.ReduceCol(i)
	}
	return reduced
}
