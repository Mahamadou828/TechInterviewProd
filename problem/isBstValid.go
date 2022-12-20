package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
	"math"
)

// IsBstValid take a binary tree and return true is valid.
// A valid binary tree is a tree where every element superior to the node goes in the Right and every element lower
// than the node goes in the Left.
//
//		5 ✅           5 ❌
//	   / \ 			  / \
//	  3   7			 6   8
//
// Time complexity: O(n), Space complexity: O(n)
func IsBstValid(tree *dts.Tree) bool {
	return traverse(tree.Root, math.Inf(-1), math.Inf(1))
}

func traverse(l *dts.Leaf, low, high float64) bool {
	if l == nil {
		return true
	}
	return l.Value > low && l.Value < high && traverse(l.Left, low, l.Value) && traverse(l.Right, l.Value, high)
}
