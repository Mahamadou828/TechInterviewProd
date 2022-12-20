package problem

// PythagoreanTriplets find in a given array the triplets of Pythagorean
// The triplets of Pythagorean is a list of 3 number that respect the law of
// a^2 + b^2 = c^2
// Time complexity O(n^2)
// Space complexity O(n)
func PythagoreanTriplets(nums []int) bool {
	mapNumsSquares, numsSquares := make(map[int]bool), []int{}
	for _, num := range nums {
		square := num * num
		mapNumsSquares[square] = true
		numsSquares = append(numsSquares, square)
	}

	for i := 0; i < len(numsSquares); i++ {
		for j := 0; j < len(numsSquares); j++ {
			if i == j {
				continue
			}
			result := numsSquares[j] + numsSquares[i]
			if ok := mapNumsSquares[result]; ok {
				return true
			}
		}
	}

	return false
}
