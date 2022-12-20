package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
)

//MergeKSortedLL takes k number of sorted linked list and return a single sorted list
//Time complexity O(k*n)
//Where k is the number of linked list and n is the number of elements in the list
//Space complexity O(n)
func MergeKSortedLL(lls []*dts.LinkedList) *dts.LinkedList {
	var currentLL *dts.LinkedList

	for _, ll := range lls {
		switch currentLL {
		case nil:
			currentLL = ll
		default:
			currentLL = merge2SortedLL([]*dts.LinkedList{currentLL, ll})
		}
	}

	return currentLL
}

func merge2SortedLL(lls []*dts.LinkedList) *dts.LinkedList {
	left, right, result := lls[0].ToArray(), lls[1].ToArray(), make([]int, lls[0].Length+lls[1].Length)
	//merge two sorted array
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return dts.NewLinkedList(result...)
}
