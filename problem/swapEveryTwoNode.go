package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// SwapEveryTwoNode
// Time complexity O(n)
// Space complexity O(n)
func SwapEveryTwoNode(ll *dts.LinkedList) *dts.LinkedList {
	res := dts.NewLinkedList()
	var stack []int

	current := ll.Head

	for current != nil {
		stack = append(stack, current.Value)
		if len(stack) == 2 {
			res.Insert(stack[1], stack[0])
			stack = nil
		}
		current = current.Next
	}

	if len(stack) != 0 {
		res.Insert(stack[0])
	}

	return res
}

//SwapEveryTwoNodeOnPlace
//Time complexity O(n)
//Space complexity O(1)
func SwapEveryTwoNodeOnPlace(ll *dts.LinkedList) *dts.LinkedList {
	node := ll.Head

	for node != nil && node.Next != nil {
		node.Value, node.Next.Value = node.Next.Value, node.Value
		node = node.Next.Next
	}

	return ll

}