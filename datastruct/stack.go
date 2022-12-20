package datastructures

// Stack is an implementation of a LIFO stack ( Last In First Out ).
type Stack struct {
	max    []int
	Elt    []int
	Length int
}

// Pop return the first element of the stack
// Time complexity O(1)
func (s *Stack) Pop() int {
	if len(s.Elt) == 0 {
		return -1
	}
	elt := s.Elt[s.Length-1]
	s.Elt, s.Length = s.Elt[:s.Length-1], s.Length-1
	if elt == s.Max() {
		s.max = s.max[:len(s.max)-1]
	}
	return elt
}

// Push insert an element into the stack
// Time complexity O(1)
func (s *Stack) Push(n int) {
	s.Elt, s.Length = append(s.Elt, n), s.Length+1
	if n > s.Max() {
		s.max = append(s.max, n)
	}
}

// Peek return the top item of the stack without removing it.
// Time complexity O(1)
func (s *Stack) Peek() int {
	return s.Elt[s.Length-1]
}

// Max return the maximum item of the stack
// Time complexity O(1)
func (s *Stack) Max() int {
	if len(s.max) == 0 {
		return -1
	}
	return s.max[len(s.max)-1]
}

// Dequeue return the last element of the stack
// Time complexity O(1)
func (s *Stack) Dequeue() int {
	if len(s.Elt) == 0 {
		return -1
	}
	elt := s.Elt[0]
	s.Elt, s.Length = s.Elt[1:s.Length], s.Length-1
	if elt == s.Max() {
		s.max = s.max[1:len(s.max)]
	}
	return elt
}

// NewStack creates a new stack
// Time complexity O(n)
// where n is the number of element
func NewStack(nums ...int) *Stack {
	stack := &Stack{}
	for _, n := range nums {
		stack.Push(n)
	}
	return stack
}
