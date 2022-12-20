package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// NumberOfCousins
// Time Complexity O(n)
// Space Complexity O(n)
func NumberOfCousins(tree *dts.Tree, target *dts.Leaf) []float64 {
	mapLevelLeaf := make(map[int]map[*dts.Leaf]bool)
	mapLeafSibling := make(map[*dts.Leaf]*dts.Leaf)

	numberOfCousinsHelper(tree.Root, 1, mapLevelLeaf, mapLeafSibling)

	for _, set := range mapLevelLeaf {
		if ok := set[target]; ok {
			delete(set, target)
			var res []float64
			sibling, hasSibling := mapLeafSibling[target]
			for leaf := range set {
				if hasSibling && leaf == sibling {
					continue
				}
				res = append(res, leaf.Value)
			}
			return res
		}
	}

	return []float64{}
}

func numberOfCousinsHelper(leaf *dts.Leaf, level int, mapLevelLeaf map[int]map[*dts.Leaf]bool, mapLeafSibling map[*dts.Leaf]*dts.Leaf) {
	if leaf == nil {
		return
	}
	set, ok := mapLevelLeaf[level]

	switch ok {
	case true:
		set[leaf] = true
		mapLevelLeaf[level] = set
	case false:
		mapLevelLeaf[level] = map[*dts.Leaf]bool{leaf: true}
	}

	mapLeafSibling[leaf.Left] = leaf.Right
	mapLeafSibling[leaf.Right] = leaf.Left

	numberOfCousinsHelper(leaf.Left, level+1, mapLevelLeaf, mapLeafSibling)
	numberOfCousinsHelper(leaf.Right, level+1, mapLevelLeaf, mapLeafSibling)
}
