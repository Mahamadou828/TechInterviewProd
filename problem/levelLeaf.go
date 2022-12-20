package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

//LevelLeaf
//Time complexity O(n)
//Space complexity O(n)
func LevelLeaf(tree *dts.Tree, level int) []int {
	mapLevelLeaf := make(map[int][]int)
	levelLeafHelper(tree.Root, 1, mapLevelLeaf)
	return mapLevelLeaf[level]
}

func levelLeafHelper(leaf *dts.Leaf, level int, mapLevelLeaf map[int][]int) {
	mapLevelLeaf[level] = append(mapLevelLeaf[level], int(leaf.Value))
	if leaf.Left != nil {
		levelLeafHelper(leaf.Left, level+1, mapLevelLeaf)
	}
	if leaf.Right != nil {
		levelLeafHelper(leaf.Right, level+1, mapLevelLeaf)
	}
}
