package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

//@todo to-test

// RotateLinkedList
// Time complexity O(n)
// Space complexity O(1)
func RotateLinkedList(ll *dts.LinkedList, n int) *dts.LLNode {
	var length int
	curr := ll.Head
	for curr != nil {
		length += 1
		curr = curr.Next
	}
	n = n % length

	slow, fast := ll.Head, ll.Head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	fast.Next = ll.Head
	head := slow.Next
	slow.Next = nil

	return head
}
