package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// FindLinkedListIntersection find the node that intersects with both linked list
//Time Complexity O(n) + O(m)
//Space Complexity O(1)
//Where n is the length of the first linked list
//And m the length of the second linked list
func FindLinkedListIntersection(a, b *dts.LinkedList) *dts.LLNode {
	lenA, lenB := 0, 0
	current := a.Head
	for current != nil {
		lenA += 1
		current = current.Next
	}
	current = b.Head
	for current != nil {
		lenB += 1
		current = current.Next
	}
	currentA, currentB := a.Head, b.Head

	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			currentA = currentA.Next
		}
	}

	if lenB > lenA {
		for i := 0; i < lenB-lenA; i++ {
			currentB = currentB.Next
		}
	}

	for currentA != currentB {
		currentA = currentA.Next
		currentB = currentB.Next
	}

	return currentA
}
