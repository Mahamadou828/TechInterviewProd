package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

func FilterBST(tree *dts.Tree, target float64) *dts.Tree {
	filterBSTHelper(tree.Root, target)
	return tree
}

func filterBSTHelper(node *dts.Leaf, n float64) *dts.Leaf {
	if node == nil {
		return nil
	}
	node.Left = filterBSTHelper(node.Left, n)
	node.Right = filterBSTHelper(node.Right, n)

	if node.Value != n && node.Left != nil && node.Right != nil {
		return nil
	}

	return node
}
