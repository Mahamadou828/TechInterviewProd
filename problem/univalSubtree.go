package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// UnivalSubtree return the number of unival subtree in this example response is 5
//
//					0
//				   / \
//				  1   0
//	              / \
//		            1   0
//	            / \
//				  1   1
//
// Time complexity O(n)
// Space complexity O(n)
func UnivalSubtree(tree *dts.Tree) int {
	return univalSubtreeHelper(tree.Root)
}

func univalSubtreeHelper(leaf *dts.Leaf) int {
	if leaf == nil {
		return 0
	}
	if leaf.Left == nil && leaf.Right == nil {
		return 1
	}
	score := 0
	score += univalSubtreeHelper(leaf.Left)
	score += univalSubtreeHelper(leaf.Right)
	if leaf.Value == leaf.Left.Value && leaf.Value == leaf.Right.Value {
		score += 1
	}
	return score
}
