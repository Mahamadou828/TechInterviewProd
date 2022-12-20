package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

//removeKthFromLL remove kth element from linked list not using the double linked list.
//Time complexity O(n)
//Space complexity O(1)
func removeKthFromLL(ll *dts.LinkedList, k int) {
	slow, fast := ll.Head, ll.Head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	prev := &dts.LLNode{}
	for fast != nil {
		prev = slow
		fast = fast.Next
		slow = slow.Next
	}
	prev.Next = slow.Next
}
