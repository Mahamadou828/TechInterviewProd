package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// InorderSuccessor
// Time complexity O(n)
// Space complexity O(n)
func InorderSuccessor(leaf *dts.Leaf) *dts.Leaf {
	if res := inorderSuccessorDownTraversalHelper(leaf.Right, leaf.Value, nil); res != nil && res != leaf {
		return res
	}
	if res := inorderSuccessorUpTraversalHelper(leaf, leaf.Value, nil); res != nil {
		return res
	}
	return nil
}

//InorderSuccessorItr
//Time complexity O(n)
//Space complexity O(1)
func InorderSuccessorItr(leaf *dts.Leaf) *dts.Leaf {
	if leaf.Right != nil {
		curr := leaf.Right
		for curr.Left != nil {
			curr = curr.Left
		}
		return curr
	}

	curr, parent := leaf, leaf.Parent

	for parent != nil && parent.Left != curr {
		curr = parent
		parent = parent.Parent
	}
	return parent
}

func inorderSuccessorUpTraversalHelper(leaf *dts.Leaf, target float64, res *dts.Leaf) *dts.Leaf {
	if leaf == nil {
		return res
	}

	if leaf.Value > target {
		switch res {
		case nil:
			res = leaf
		default:
			if leaf.Value < res.Value {
				res = leaf
			}
		}
	}

	res = inorderSuccessorUpTraversalHelper(leaf.Parent, target, res)

	return res
}

func inorderSuccessorDownTraversalHelper(leaf *dts.Leaf, target float64, res *dts.Leaf) *dts.Leaf {
	if leaf == nil {
		return res
	}

	if leaf.Value > target {
		switch res {
		case nil:
			res = leaf
		default:
			if leaf.Value < res.Value {
				res = leaf
			}
		}
	}

	res = inorderSuccessorDownTraversalHelper(leaf.Left, target, res)

	return res
}
