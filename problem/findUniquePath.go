package problem

// FindUniquePath find the number of unique path that allow to go from
// the beginning of a matrix to the end of the matrix
// Time complexity: O(m*n)
// Space complexity O(m*n)
// Where m is the len of the matrix
// And n the height of the matrix
// @todo to re-do understand the solution
func FindUniquePath(m, n int) int {
	cache := make([][]int, m)
	initCell := make([]int, n)
	for i := range initCell {
		initCell[i] = 0
	}
	for i := range cache {
		cache[i] = append(cache[i], initCell...)
	}
	for i := 0; i < m; i++ {
		cache[i][0] = 1
	}
	for j := 0; j < n; j++ {
		cache[0][j] = 1
	}

	for c := 1; c < m; c++ {
		for r := 1; r < n; r++ {
			cache[c][r] = cache[c][r-1] + cache[c-1][r]
		}
	}

	return cache[m-1][n-1]
}
