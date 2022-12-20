package problem

import (
	"math"
)

// JumpToTheEnd
// Time complexity O(n^2)
// Space complexity O(n)
func JumpToTheEnd(nums []int) int {
	var hops []int

	for i := 0; i < len(nums); i++ {
		hops = append(hops, math.MaxInt)
	}

	hops[0] = 0
	for i, n := range nums {
		for j := 1; j < n+1; j++ {
			if i+j < len(hops) {
				hops[i+j] = min(hops[i+j], hops[i]+1)
			} else {
				break
			}
		}
	}

	return hops[len(hops)-1]
}
