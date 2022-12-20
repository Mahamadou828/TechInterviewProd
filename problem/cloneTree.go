package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// CloneTree find the clone node of a tree
// Time complexity O(n)
// Space complexity O(1)
func CloneTree(a, b, node *dts.Leaf) *dts.Leaf {
	if a == node {
		return nil
	}
	if a.Left != nil && b.Left != nil {
		found := CloneTree(a.Left, b.Left, node)
		if found != nil {
			return found
		}
	}
	if a.Right != nil && b.Right != nil {
		found := CloneTree(a.Right, b.Right, node)
		if found != nil {
			return found
		}
	}
	return nil
}

// CloneTreeItr the iterative version of clone tree
func CloneTreeItr(a, b, node *dts.Leaf) *dts.Leaf {
	stack := []*dts.Leaf{a, b}

	for len(stack) > 0 {
		a, b := stack[0], stack[1]
		if a == node {
			return b
		}
		if a.Left != nil && b.Left != nil {
			stack = append(stack, a.Left, b.Left)
		}
		if a.Right != nil && b.Right != nil {
			stack = append(stack, a.Right, b.Right)
		}
	}
	return nil
}
