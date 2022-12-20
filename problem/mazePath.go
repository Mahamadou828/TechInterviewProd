package problem

// MazePath
// Time complexity O(n)
// Space complexity O(n)
func MazePath(maze [][]int) int {
	var paths [][]int
	for _, row := range maze {
		var res []int
		for range row {
			res = append(res, 0)
		}
		paths = append(paths, res)
	}

	paths[0][0] = 1

	for i, row := range maze {
		for j, val := range row {
			if val == 1 || (i == 0 && j == 0) {
				continue
			}
			var leftPaths, topPaths int
			if i > 0 {
				leftPaths = paths[i-1][j]
			}
			if j > 0 {
				topPaths = paths[i][j-1]
			}
			paths[i][j] = leftPaths + topPaths
		}
	}

	return paths[len(paths)-1][len(paths[0])-1]
}
