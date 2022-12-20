package problem

import (
	"fmt"
)

// MaximumConnectedColors
// Time complexity O(n*k)
// Space complexity O(n)
// Where k is the length and n the depth of the grid
func MaximumConnectedColors(grid [][]int) int {
	var max int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			elt := grid[i][j]
			var currentMax int
			maximumConnectedColorsHelper(grid, i, j, elt, &currentMax, make(map[string]bool))
			max = maxInt(max, currentMax)
		}
	}

	return max
}

func maximumConnectedColorsHelper(grid [][]int, i, j, target int, counter *int, seenMap map[string]bool) {
	if i >= len(grid) || j >= len(grid[0]) {
		return
	}

	switch grid[i][j] {
	case target:
		seenMap[fmt.Sprintf("%d-%d", i, j)] = true
		*counter += 1
	default:
		return
	}

	//Can we explore left ?
	if j-1 > 0 && !seenHelper(seenMap, i, j-1) {
		maximumConnectedColorsHelper(grid, i, j-1, target, counter, seenMap)
	}
	//Can we explore right ?
	if j+1 < len(grid[0]) && !seenHelper(seenMap, i, j+1) {
		maximumConnectedColorsHelper(grid, i, j+1, target, counter, seenMap)
	}
	//Can we explore top ?
	if i-1 > 0 && !seenHelper(seenMap, i-1, j) {
		maximumConnectedColorsHelper(grid, i-1, j, target, counter, seenMap)
	}
	//Can we explore bottom ?
	if i+1 < len(grid) && !seenHelper(seenMap, i+1, j) {
		maximumConnectedColorsHelper(grid, i+1, j, target, counter, seenMap)
	}
}

func seenHelper(seen map[string]bool, i, j int) bool {
	return seen[fmt.Sprintf("%d-%d", i, j)]
}
