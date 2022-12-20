package problem

// LongestIncreasingSubsequence
// Time complexity O(n)
// Space complexity O(n)
func LongestIncreasingSubsequence(nums []int) int {
	var res int
	var cache []int

	for i := 0; i < len(nums); i++ {
		cache = append(cache, 1)
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				cache[i] = maxInt(cache[i], cache[j]+1)
			}
		}
	}

	for _, val := range cache {
		res = maxInt(val, res)
	}

	return res
}
