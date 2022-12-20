package problem

//@todo to re-do

// NonDecreasingArray https://www.techseries.dev/products/coderpro/categories/1831147/posts/6228786
// Time complexity O(n)
// Space complexity O(1)
func NonDecreasingArray(nums []int) bool {
	invalidIdx := -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			if invalidIdx != -1 {
				return false
			}
			invalidIdx = i
		}
	}

	if invalidIdx == -1 {
		return true
	}
	if invalidIdx == 0 {
		return true
	}
	if invalidIdx == len(nums)-2 {
		return true
	}
	if nums[invalidIdx] <= nums[invalidIdx+2] {
		return true
	}
	if nums[invalidIdx-1] <= nums[invalidIdx+1] {
		return true
	}
	return false
}
