package problem

const (
	right = iota
	down
	left
	up
)

// SpiralTraverse @todo to re-do because it's not my solution
func SpiralTraverse(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	result := []int{}
	spDir := []int{right, down, left, up}
	// initialize the 4 boundaries
	l, r, t, b := 0, n-1, 0, m-1
	// traverse the matrix in spiral order
	for round := 0; l <= r && t <= b; round++ {
		// in each round, we traverse in 1 direction
		switch spDir[round%4] {
		case 0:
			// traverse in right direction
			for i, j := l, t; i <= r; i++ {
				result = append(result, matrix[j][i])
			}
			t++
		case 1:
			// traverse in down direction
			for i, j := r, t; j <= b; j++ {
				result = append(result, matrix[j][i])
			}
			r--
		case 2:
			// traverse in left direction
			for i, j := r, b; i >= l; i-- {
				result = append(result, matrix[j][i])
			}
			b--
		case 3:
			// traverse in up direction
			for i, j := l, b; j >= t; j-- {
				result = append(result, matrix[j][i])
			}
			l++
		}
	}
	for _, num := range result {
		print(" ", num)
	}
}
