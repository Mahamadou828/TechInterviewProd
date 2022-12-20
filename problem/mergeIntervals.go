package problem

import (
	"sort"
)

//@todo to re-do

// MergeIntervals merge all intervals
// Time complexity O(nlog n )
// Space complexity O(n)
func MergeIntervals(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		current, prev := intervals[i], intervals[i-1]

		if prev[0] < current[0] && prev[1] > current[1] {
			continue
		}

		res = append(res, current)
	}

	return res
}
