package solution

import dts "github.com/Mahamadou828/coderpro/datastruct"

func removeZeroSumSublists(head *dts.LLNode) *dts.LLNode {
	curr := head
	finalHead := head

	done := false
	for !done {
		cleaned, changed := _removeZeroSumSublists(curr)

		if !changed {
			done = true
			break
		}

		if cleaned == nil && !changed {
			done = true
			break
		}
		// means everything was removed
		if cleaned == nil && changed {
			return nil
		}
		if cleaned != finalHead {
			finalHead = cleaned
			curr = finalHead
			continue
		}
	}

	return finalHead

}

func _removeZeroSumSublists(head *dts.LLNode) (*dts.LLNode, bool) {
	var count int
	countTracker := make(map[int]*dts.LLNode)
	curr := head
	for curr != nil {
		count += curr.Value
		// if count == 0, there is zero sum starting with head. return next node
		if count == 0 {
			return curr.Next, true
		}
		prevNode, ok := countTracker[count]
		if ok {
			prevNode.Next = curr.Next
			return head, true
		}
		countTracker[count] = curr
		curr = curr.Next
	}
	return head, false

}
