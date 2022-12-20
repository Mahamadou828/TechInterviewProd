package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
	"math"
)

// DepthBinaryTree find the depth of the binary tree
// Time complexity O(n)
// where n is the number of leaf
// Space complexity O(n) worse case, O(log n) for a balanced tree
func DepthBinaryTree(tree *dts.Tree) int {
	return traverseLeaf(tree.Root, 0)
}

func traverseLeaf(leaf *dts.Leaf, score int) int {
	leftScore, rightScore, score := 0, 0, score+1

	if leaf.Left != nil {
		leftScore = traverseLeaf(leaf.Left, score)
	}
	if leaf.Right != nil {
		rightScore = traverseLeaf(leaf.Right, score)
	}

	if leftScore > rightScore && leftScore > score {
		score = leftScore
	}
	if rightScore > leftScore && rightScore > score {
		score = rightScore
	}

	return score
}

// DepthBinaryTreeIterative find the depth of the binary tree
// Time complexity O(n)
// where n is the number of leaf
// Space complexity O(n) worse case
func DepthBinaryTreeIterative(tree *dts.Tree) int {
	type item struct {
		node  *dts.Leaf
		depth int
	}
	var stack []item
	pop := func(stack []item) item {
		x, stack := stack[len(stack)-1], stack[:len(stack)-1]
		return x
	}

	stack = append(stack, item{tree.Root, 1})
	maxDepth := 0
	for len(stack) > 0 {
		elt := pop(stack)
		if elt.node != nil {
			maxDepth = int(math.Max(float64(maxDepth), float64(elt.depth)))
			stack = append(stack, item{elt.node.Left, elt.depth + 1})
			stack = append(stack, item{elt.node.Right, elt.depth + 1})
		}
	}
	return maxDepth
}
