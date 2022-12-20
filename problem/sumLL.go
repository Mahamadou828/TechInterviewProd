package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// SumLL takes two linked list and return a linked list of the sum of both
// more detail here https://www.techseries.dev/products/coderpro/categories/1831147/posts/6176705
// Time complexity O(n)
// Space complexity O(1)
func SumLL(ll1 *dts.LinkedList, ll2 *dts.LinkedList) *dts.LinkedList {
	//for now let's treat only list that are same in size
	if ll1.Length != ll2.Length {
		return nil
	}

	var remaining int
	ll := &dts.LinkedList{}
	for i := ll1.Length; i > 0; i-- {
		val := ll1.Pop() + ll2.Peak() + remaining
		if val >= 10 {
			val, remaining = val%10, val/10
			//handle remaining
		}
		ll.InsertAtTail(val)
	}

	return ll
}

// SumLLRec is the recursive version of SumLL
// Time complexity O(n)
// Space complexity O(n)
func SumLLRec(ll1 *dts.LinkedList, ll2 *dts.LinkedList) *dts.LinkedList {
	ll := &dts.LinkedList{}
	sumrec(ll1, ll2, ll, 0)
	return ll
}

func sumrec(ll1 *dts.LinkedList, ll2 *dts.LinkedList, dest *dts.LinkedList, remaining int) {
	if ll1.Length == 0 && ll2.Length == 0 {
		return
	}
	val := ll1.Pop() + ll2.Peak() + remaining
	if val >= 10 {
		val, remaining = val%10, val/10
	}
	dest.InsertAtTail(val)
	sumrec(ll1, ll2, dest, remaining)
}
