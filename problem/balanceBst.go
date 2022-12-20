package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
	"math"
)

// IsHeightBalance return if a binary tree is balanced according to the following rules:
// a binary tree is considered balanced if its left and right branches have the same height
// Time complexity O(n)
// Space complexity O(n)
func IsHeightBalance(tree *dts.Tree) bool {
	ok, _ := balancedBstHelper(tree.Root)
	return ok
}

func balancedBstHelper(leaf *dts.Leaf) (bool, int) {
	if leaf == nil {
		return true, 0
	}
	lBalanced, lHeight := balancedBstHelper(leaf.Left)
	rBalanced, rHeight := balancedBstHelper(leaf.Right)

	return lBalanced && rBalanced && math.Abs(float64(lHeight-rHeight)) <= 1, int(math.Max(float64(lHeight), float64(rHeight))) + 1
}
