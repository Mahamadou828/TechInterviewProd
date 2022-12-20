package problem

// NumberOfIslands
// Time complexity O(n*k)
// Space complexity O(n*k)
// Where n is the len and k is the depth
func NumberOfIslands(matrix [][]int) int {
	var numberOfIsland int

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			island := matrix[i][j]
			if island == 1 || island == 2 {
				continue
			}
			numberOfIsland++
			explore(matrix, i, j)
		}
	}

	return numberOfIsland
}

func explore(matrix [][]int, i int, j int) {
	if i >= len(matrix) || j >= len(matrix[0]) {
		return
	}
	//if given cell is water then return
	if matrix[i][j] == 1 || matrix[i][j] == 2 {
		return
	}
	//if it's not we have to mark it as explore
	matrix[i][j] = 2

	//then we have to ask
	//can we explore left ?
	if j-1 > 0 {
		explore(matrix, i, j-1)
	}
	//can we explore right ?
	if j+1 < len(matrix[0]) {
		explore(matrix, i, j+1)
	}
	//can we explore down ?
	if i+1 < len(matrix) {
		explore(matrix, i+1, j)
	}
	//can we explore up ?
	if i-1 > 0 {
		explore(matrix, i+1, j)
	}
}
