package algo

// Queue is a data structure that allows to keep track of elements in a first-in-first-out manner
type Queue []int

// NewQueue creates a new Queue data structure
func NewQueue() *Queue {
	return &Queue{}
}

// Dequeue removes the first element from the queue and returns it
func (q *Queue) Dequeue() int {
	val := (*q)[0]
	if !q.IsEmpty() {
		*q = (*q)[1:]
	}

	return val
}

// Enqueue adds a new element to the queue
func (q *Queue) Enqueue(value int) {
	*q = append(*q, value)
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
