package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

//InvertBST invert a binary tree
//Time complexity O(n)
//Space complexity O(n)
func InvertBST(tree *dts.Tree) *dts.Tree {
	leaf := invertBSTHelper(tree.Root)

	return &dts.Tree{Root: leaf}
}

func invertBSTHelper(leaf *dts.Leaf) *dts.Leaf {
	nLeaf := &dts.Leaf{Value: leaf.Value}
	if leaf.Left != nil {
		nLeaf.Right = invertBSTHelper(leaf.Left)
	}
	if leaf.Right != nil {
		nLeaf.Left = invertBSTHelper(leaf.Right)
	}
	return nLeaf
}
