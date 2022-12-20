package problem

import "math"

// FirstPositiveMissingInt find the first missing positive integer
// Time complexity O(n)
// Space complexity O(n)
func FirstPositiveMissingInt(nums []int) int {
	mapNums := make(map[int]bool)

	for _, num := range nums {
		mapNums[num] = true
	}

	for i := 1; i < math.MaxInt; i++ {
		if _, ok := mapNums[i]; !ok {
			return i
		}
	}

	return 0
}
