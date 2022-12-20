package problem

import (
	"fmt"
	dts "github.com/Mahamadou828/coderpro/datastruct"
)

// ZigZagBST
// Time complexity O(n)
// Space complexity O(n)
func ZigZagBST(tree *dts.Tree) []float64 {
	var res []float64
	mapLevelLeaf := make(map[int][]float64)
	zigZagBSTHelper(tree.Root, mapLevelLeaf, 0)
	fmt.Println(mapLevelLeaf)
	for level, values := range mapLevelLeaf {
		switch level % 2 {
		case 0:
			for i := len(values) - 1; i >= 0; i-- {
				res = append(res, values[i])
			}
		default:
			for _, v := range values {
				res = append(res, v)
			}

		}
	}

	return res
}

func zigZagBSTHelper(leaf *dts.Leaf, mapLevelLeaf map[int][]float64, level int) {
	if leaf == nil {
		return
	}
	levelVals, ok := mapLevelLeaf[level]
	switch ok {
	case true:
		levelVals = append(levelVals, leaf.Value)
		mapLevelLeaf[level] = levelVals
	case false:
		mapLevelLeaf[level] = []float64{leaf.Value}
	}

	if leaf.Left != nil {
		zigZagBSTHelper(leaf.Left, mapLevelLeaf, level+1)
	}
	if leaf.Right != nil {
		zigZagBSTHelper(leaf.Right, mapLevelLeaf, level+1)
	}
}
