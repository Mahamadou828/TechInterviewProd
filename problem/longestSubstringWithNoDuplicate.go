package problem

import "math"

//@todo a comprendre

func LongestSubstringWithNoDuplicate(strs []string) int {
	start, end := -1, 0
	mapSeenLetter, maxLen := make(map[string]int), math.MinInt

	for end < len(strs) {
		c := strs[end]
		if _, ok := mapSeenLetter[c]; ok {
			start = maxInt(start, mapSeenLetter[c])
		}
		maxLen = maxInt(maxLen, end-start)
		mapSeenLetter[c] = end
		end += 1
	}

	return maxLen
}
