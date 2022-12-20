package problem

// SubArrayTargetSum
// Time complexity O(n)
// Space complexity O(n)
func SubArrayTargetSum(nums []int, target int) []int {
	var res []int

	switch true {
	case len(nums) == 1:
		if nums[0] != target {
			return res
		}
		return nums
	case len(nums) == 2:
		if nums[0]+nums[1] != target {
			return res
		}
		return res
	case len(nums) == 0:
		return res
	}

	leftPtr, rightPtr, addRight := 0, 1, true
	currentSum := nums[leftPtr]

	for rightPtr < len(nums) && leftPtr < len(nums) {
		if addRight {
			currentSum = currentSum + nums[rightPtr]
		}
		if currentSum < target {
			rightPtr += 1
			addRight = true
		}
		if currentSum > target {
			currentSum -= nums[leftPtr]
			leftPtr += 1
			addRight = false
		}
		if currentSum == target {
			res = nums[leftPtr : rightPtr+1]
			break
		}
	}

	return res
}
