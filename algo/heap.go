package algo

// Heap is a generic type that allows to create an array of AllowedTypes
type Heap[T AllowedTypes] Array[T]

// NewHeapFromArray is a constructor for the Heap type
func NewHeapFromArray[T AllowedTypes](array Array[T]) Heap[T] {
	return Heap[T](array)
}

// NewHeap is a constructor for the Heap type
func NewHeap[T AllowedTypes](n int) Heap[T] {
	return Heap[T](NewArray[T](n))
}

// ShiftDown is a method that shifts the element down the heap
func (h Heap[T]) ShiftDown(n, i int) {
	k := i
	for {
		j := k

		if 2*j+1 < n && (h[2*j+1] > h[k]) {
			k = 2*j + 1
		}
		if 2*j+2 < n && (h[2*j+2] > h[k]) {
			k = 2*j + 2
		}

		h.Swap(j, k)

		if j == k {
			break
		}
	}
}

// BuildHeap is a method that builds the heap
func (h Heap[T]) BuildHeap() {
	for i := (len(h) - 1) / 2; i >= 0; i-- {
		h.ShiftDown(len(h), i)
	}
}

// Swap is a method that swaps two elements in the heap
func (h Heap[T]) Swap(p1, p2 int) {
	Array[T](h).Swap(p1, p2)
}
