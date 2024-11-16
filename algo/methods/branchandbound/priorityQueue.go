package branchandbound

import (
	"errors"
	"sort"
)

// PriorityQueue is a data structure that allows to keep track of elements in a priority queue
type PriorityQueue struct {
	arr  []*Node
	sort func(pObj, cObj *Node) bool
}

// NewPriorityQueue is a constructor for the PriorityQueue type
func NewPriorityQueue(arr []*Node) *PriorityQueue {
	c := make([]*Node, len(arr))
	copy(c, arr)
	pq := PriorityQueue{
		arr:  c,
		sort: Min,
	}
	return &pq
}

// shiftUp is a method that shifts the element up the queue
func (h *PriorityQueue) shiftUp(i int) *PriorityQueue {
	for i > 0 && h.sort(h.arr[i], h.arr[(i-1)/2]) {
		h.Swap((i-1)/2, i)
		i = (i - 1) / 2
	}
	return h
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
		h.arr[p1], h.arr[p2] = h.arr[p2], h.arr[p1]
	}
	return h
}

// IsEmpty is a method that checks if the heap is empty
func (h *PriorityQueue) IsEmpty() bool {
	return len(h.arr) == 0
}

// GetRoot is a method that returns the root of the heap
func (h *PriorityQueue) GetRoot() (*Node, error) {
	if h.IsEmpty() {
		return nil, errors.New("tried to get root from empty priority queue")
	}
	tmp := h.arr[0]
	n := len(h.arr) - 1
	h.Swap(0, n)
	h.arr = h.arr[:n]
	h.shiftDown(n, 0)
	return tmp, nil
}

// Insert is a method that inserts and builds up the heap
func (h *PriorityQueue) Insert(n *Node) *PriorityQueue {
	h.arr = append(h.arr, n)
	h.shiftUp(len(h.arr) - 1)
	return h
}

// SetSort is a method that sets the sorting function for the heap
func (h *PriorityQueue) SetSort(s func(pIdx, cIdx *Node) bool) *PriorityQueue {
	h.sort = s
	return h
}

// Min is a function that compares two elements in the heap
func Min(pObj, cObj *Node) bool {
	return pObj.val < cObj.val
}

// Remove removes on new best
func (h *PriorityQueue) Remove(bound int) *PriorityQueue {
	sort.Slice(h.arr, func(i, j int) bool {
		return Min(h.arr[i], h.arr[j])
	})
	i := 0
	for ; i <= len(h.arr)-1 && h.arr[i].val < bound; i++ {
		continue
	}
	h.arr = h.arr[:i+1]
	h.BuildQueue()
	return h
}

func (h *PriorityQueue) RemoveN() {
	n := 100000
	if len(h.arr) > n {
		h.arr = h.arr[:n]
	}
}
