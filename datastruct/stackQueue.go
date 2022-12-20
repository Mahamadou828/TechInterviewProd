package datastructures

// StackQueue is a implementation of a queue using a stack
type StackQueue struct {
	s1 []int
	s2 []int
}

// Dequeue return the first element of the queue.
// Time complexity O(n)
func (s *StackQueue) Dequeue() int {
	if len(s.s2) == 0 && len(s.s1) == 0 {
		return -1
	}

	if len(s.s2) != 0 {
		elt := s.s2[len(s.s2)-1]
		s.s2 = s.s2[:len(s.s2)-1]
		return elt
	}

	for len(s.s1) > 0 {
		elt := s.s1[len(s.s1)-1]
		s.s2 = append(s.s2, elt)
		s.s1 = s.s1[:len(s.s1)-1]
	}

	elt := s.s2[len(s.s2)-1]
	s.s2 = s.s2[:len(s.s2)-1]
	return elt
}

// Enqueue insert an element into the queue
// Time complexity O(1)
func (s *StackQueue) Enqueue(n int) {
	s.s1 = append(s.s1, n)
}

// NewStackQueue creates a new stack queue
// Time complexity O(n)
// where n is the number of element
func NewStackQueue(nums ...int) *StackQueue {
	stack := &StackQueue{}
	for _, n := range nums {
		stack.Enqueue(n)
	}
	return stack
}
