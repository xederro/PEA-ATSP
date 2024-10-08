package algo

// Sets is a data structure that allows to keep track of disjoint sets
type Sets struct {
	parent Array[int]
	rank   Array[int]
}

// NewSets creates a new Sets data structure
func NewSets(v int) *Sets {
	a := NewArray[int](v).PopulateWithCounting()
	b := NewArray[int](v)
	return &Sets{
		parent: a,
		rank:   b,
	}
}

// FindSet finds the set that contains the element v
func (s *Sets) FindSet(v int) int {
	a := v
	for a != s.parent[a] {
		a = s.parent[a]
	}
	s.collapseSet(v, a)
	return a
}

// collapseSet collapses the set of the element v to the set p
func (s *Sets) collapseSet(v, p int) {
	a := v
	for a != s.parent[a] {
		tmp := a
		a = s.parent[a]
		s.parent[tmp] = p
	}
}

// Union merges the sets that contain the elements x and y
func (s *Sets) Union(x, y int) {
	a := s.FindSet(x)
	b := s.FindSet(y)
	if s.rank[a] < s.rank[b] {
		s.parent[a] = b
	} else {
		s.parent[b] = a
	}

	if s.rank[a] == s.rank[b] {
		s.rank[a] = s.rank[a] + 1
	}
}

// IsSameSet checks if the elements v and u are in the same set
func (s *Sets) IsSameSet(v, u int) bool {
	return s.FindSet(v) == s.FindSet(u)
}
