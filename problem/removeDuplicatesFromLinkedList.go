package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

//RemoveDuplicatesFromLinkedList
//Time complexity O(n)
//Space complexity O(n)
func RemoveDuplicatesFromLinkedList(ll *dts.LinkedList) *dts.LinkedList {
	prev, current := ll.Head, ll.Head
	seen := make(map[int]bool)

	for current != nil {
		_, ok := seen[current.Value]

		switch ok {
		case false:
			seen[current.Value] = true
			break
		case true:
			prev.Next = current.Next
			current = prev
			break
		}

		prev = current
		current = current.Next
	}

	return ll
}
