package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// ReverseLinkedList take and reverse a linked list.
// Time complexity O(n)
// Space complexity O(1)
func ReverseLinkedList(ll *dts.LinkedList) *dts.LinkedList {
	current, temp := ll.Tail, &dts.LLNode{}

	ll.Tail, ll.Head = ll.Head, ll.Tail

	for current != nil {
		prev := current.Prev
		current.Prev = current.Next
		current.Next, temp = prev, prev
		current = temp
	}

	return ll
}

// ReverseSingleLinkedList take and reverse a single linked list.
// Time complexity O(n)
// Space complexity O(1)
func ReverseSingleLinkedList(ll *dts.LinkedList) *dts.LinkedList {
	current, prev := ll.Head, &dts.LLNode{}

	for current != nil {
		temp := current.Next

		if current == ll.Head {
			current.Next, prev, current = nil, current, temp
			continue
		}

		current.Next = prev

		if temp == nil {
			ll.Head = current
		}

		current, prev = temp, current
	}

	return ll
}
