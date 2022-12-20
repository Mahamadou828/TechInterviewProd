package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

//@todo to re-do

func RemoveZeroSumConsecutiveNodes(ll *dts.LinkedList) *dts.LinkedList {
	mapIndexNode, index, sum := make(map[int]*dts.LLNode), 0, 0
	ll.Traverse(func(node *dts.LLNode) {
		mapIndexNode[index] = node
		index, sum = index+1, sum+node.Value
	})

	if sum == 0 {
		return nil
	}

	leftPtr, rightPtr := 0, ll.Length-1

	for leftPtr != rightPtr && leftPtr < ll.Length && rightPtr > 0 {
		left, right := mapIndexNode[leftPtr], mapIndexNode[rightPtr]

		//it's mean everything between left and right ( excluding left ) sum to zero
		if sum-left.Value == 0 {
			if leftPtr == 0 {
				ll.Head = left.Next
				left.Next.Prev = nil
				rightNode := mapIndexNode[rightPtr-1]
				ll.Head.Next = rightNode
				break
			}
			leftNode, rightNode := mapIndexNode[leftPtr], mapIndexNode[rightPtr-1]
			leftNode.Next = rightNode
		}
		//it's mean everything between left and right ( excluding right ) sum to zero
		if sum-right.Value == 0 {
			if rightPtr == ll.Length {
				ll.Tail = right.Prev
				right.Prev.Next = nil
				leftNode := mapIndexNode[leftPtr]
				leftNode.Next = ll.Tail
				break
			}
			leftNode, rightNode := mapIndexNode[leftPtr-1], mapIndexNode[rightPtr]
			leftNode.Next = rightNode
		}
		//it's mean everything between left and right ( excluding left and right ) sum to zero
		if sum-(left.Value+right.Value) == 0 {
			leftNode, rightNode := mapIndexNode[leftPtr], mapIndexNode[rightPtr]
			leftNode.Next = rightNode
		}

		leftPtr, rightPtr, sum = leftPtr+1, rightPtr-1, sum-left.Value-right.Value
	}

	return ll

}
