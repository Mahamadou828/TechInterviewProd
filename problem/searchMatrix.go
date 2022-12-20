package problem

// SearchMatrixBrute
// Time complexity O(n*k)
// Space complexity O(1)
// Where n is the length of the and k the height
func SearchMatrixBrute(matrix [][]int, k int) bool {
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == k {
				return true
			}
		}
	}

	return false
}

func SearchMatrix(matrix [][]int, k int) bool {
	for i := range matrix {
		if k >= matrix[i][0] && k <= matrix[i][len(matrix[i])-1] {
			for j := range matrix[i] {
				if matrix[i][j] == k {
					return true
				}
			}
		}
	}
	return false
}
