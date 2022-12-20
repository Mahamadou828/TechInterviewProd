package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

// LevelByLevelTree
// Time complexity O(n)
// Space complexity O(n)
func LevelByLevelTree(tree *dts.Tree) map[int][]float64 {
	mapLevelElt := make(map[int][]float64)
	levelByLevelTreeHelper(tree.Root, 0, mapLevelElt)

	return mapLevelElt
}

func levelByLevelTreeHelper(leaf *dts.Leaf, currentLevel int, mapLevelElt map[int][]float64) {
	switch elt, ok := mapLevelElt[currentLevel]; ok {
	case true:
		mapLevelElt[currentLevel] = append(elt, leaf.Value)
	default:
		mapLevelElt[currentLevel] = []float64{leaf.Value}
	}
	if leaf.Left != nil {
		levelByLevelTreeHelper(leaf.Left, currentLevel+1, mapLevelElt)
	}
	if leaf.Right != nil {
		levelByLevelTreeHelper(leaf.Right, currentLevel+1, mapLevelElt)
	}
}

// LevelByLevelTree2
// Time complexity O(n)
// Space complexity O(n)
func LevelByLevelTree2(tree *dts.Tree) []float64 {
	res := []float64{tree.Root.Value}
	levelByLevelTreeHelper2(tree.Root, func(val float64) {
		res = append(res, val)
	})
	return res
}

func levelByLevelTreeHelper2(leaf *dts.Leaf, append func(float64)) {
	if leaf == nil {
		return
	}
	//append(leaf.Value)
	if leaf.Left != nil {
		append(leaf.Left.Value)
	}
	if leaf.Right != nil {
		append(leaf.Right.Value)
	}
	levelByLevelTreeHelper2(leaf.Left, append)
	levelByLevelTreeHelper2(leaf.Right, append)
}
