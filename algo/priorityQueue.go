package algo

import (
	"errors"
	"log"
)

// member is a struct that represents an element in the priority queue
type member struct {
	self     any
	key      *int
	qp       int
	contains bool
}

// PriorityQueue is a data structure that allows to keep track of elements in a priority queue
type PriorityQueue struct {
	arr  []*member
	pos  []*member
	sort func(pObj, cObj *member) bool
}

// NewPriorityQueueFromEdges is a constructor for the PriorityQueue type
func NewPriorityQueueFromEdges(m *map[[2]int]*int) *PriorityQueue {
	a := []*member{}
	b := make([]bool, len(*m))
	c := make([]*member, len(*m))
	for ints := range *m {
		a = append(a, &member{
			key:      (*m)[ints],
			self:     ints,
			contains: true,
		})
	}
	for i := range b {
		b[i] = true
		a[i].qp = i
		c[i] = a[i]
	}

	pq := PriorityQueue{
		arr:  a,
		pos:  c,
		sort: Min,
	}

	return &pq
}

// NewPriorityQueueFromNodes is a constructor for the PriorityQueue type
func NewPriorityQueueFromNodes(m *[]int) *PriorityQueue {
	a := []*member{}
	b := make([]bool, len(*m))
	c := make([]*member, len(*m))
	for k := range *m {
		a = append(a, &member{
			key:      &(*m)[k],
			self:     k,
			qp:       k,
			contains: true,
		})
	}
	for i := range b {
		b[i] = true
		c[i] = a[i]
	}

	pq := PriorityQueue{
		arr:  a,
		pos:  c,
		sort: Min,
	}

	return &pq
}

// NewPriorityQueue is a constructor for the PriorityQueue type
func NewPriorityQueue(n int) *PriorityQueue {
	if n < 0 {
		log.Fatalln("Provided invalid size")
	}
	pq := PriorityQueue{
		arr: make([]*member, n),
	}
	return &pq
}

// shiftDown is a method that shifts the element down the queue
func (h *PriorityQueue) shiftDown(n, i int) *PriorityQueue {
	k := i
	for {
		j := k

		if 2*j+1 < n && (h.sort(h.arr[2*j+1], h.arr[k])) {
			k = 2*j + 1
		}
		if 2*j+2 < n && (h.sort(h.arr[2*j+2], h.arr[k])) {
			k = 2*j + 2
		}

		h.Swap(j, k)

		if j == k {
			break
		}
	}
	return h
}

// ShiftUp is a method that shifts the element up the queue
func (h *PriorityQueue) ShiftUp(i int) *PriorityQueue {
	return h.shiftUp(h.pos[i].qp)
}

// shiftUp is a method that shifts the element up the queue
func (h *PriorityQueue) shiftUp(i int) *PriorityQueue {
	for i > 0 && h.sort((h.arr)[i], h.arr[(i-1)/2]) {
		h.Swap((i-1)/2, i)
		i = (i - 1) / 2
	}
	return h
}

// BuildQueue is a method that builds the heap
func (h *PriorityQueue) BuildQueue() *PriorityQueue {
	for i := (len(h.arr) - 1) / 2; i >= 0; i-- {
		h.shiftDown(len(h.arr), i)
	}
	return h
}

// Swap is a method that swaps two elements in the heap
func (h *PriorityQueue) Swap(p1, p2 int) *PriorityQueue {
	if p1 != p2 {
		h.arr[p1].qp, h.arr[p2].qp = h.arr[p2].qp, h.arr[p1].qp
		h.arr[p1], h.arr[p2] = h.arr[p2], h.arr[p1]
	}
	return h
}

// IsEmpty is a method that checks if the heap is empty
func (h *PriorityQueue) IsEmpty() bool {
	return len(h.arr) == 0
}

// GetRoot is a method that returns the root of the heap
func (h *PriorityQueue) GetRoot() (any, error) {
	if h.IsEmpty() {
		return nil, errors.New("tried to get root from empty priority queue")
	}
	tmp := h.arr[0]
	n := len(h.arr) - 1
	h.Swap(0, n)
	h.arr = h.arr[:n]
	h.shiftDown(n, 0)
	tmp.contains = false
	return tmp.self, nil
}

// SetSort is a method that sets the sorting function for the heap
func (h *PriorityQueue) SetSort(s func(pIdx, cIdx *member) bool) *PriorityQueue {
	h.sort = s
	return h
}

// Contains is a method that checks if the heap contains a certain element
func (h *PriorityQueue) Contains(c int) bool {
	return h.pos[c].contains
}

// Min is a function that compares two elements in the heap
func Min(pObj, cObj *member) bool {
	return *pObj.key < *cObj.key
}
