package algo

import (
	"errors"
)

// Stack is a generic type that allows to create an array of AllowedTypes
type Stack[T AllowedTypes] struct {
	arr Array[T]
	top int
}

// NewStack is a constructor for the Stack type
func NewStack[T AllowedTypes](n int) Stack[T] {
	return Stack[T]{
		arr: NewArray[T](n),
		top: -1,
	}
}

// Empty is a function that tells if stack is empty
func (s *Stack[T]) Empty() bool {
	return s.top == -1
}

// Push pushes entry down the stack
func (s *Stack[T]) Push(entry T) {
	s.top++
	s.arr[s.top] = entry
}

// Pop pops entry from the top of stack
func (s *Stack[T]) Pop() (T, error) {
	if s.Empty() {
		return T(0), errors.New("empty Stack")
	}
	tmp := s.arr[s.top]
	s.top--
	return tmp, nil
}
