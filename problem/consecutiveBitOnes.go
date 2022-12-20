package problem

import "math"

func ConsecutiveBitOnes(n int) int {
	const BITMASK = 1
	maxCount, count := math.MinInt, 0

	for n != 0 {
		if n&BITMASK == 0 {
			maxCount = maxInt(maxCount, count)
			count = 0
		} else {
			count += 1
		}
		n = n >> 1
	}
	maxCount = maxInt(maxCount, count)

	return maxCount
}
