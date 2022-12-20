package problem

//MaxSubArraySum return the max sum in a sub-array
//Time complexity O(n)
//Space complexity O(1)
func MaxSubArraySum(nums []int) int {
	var totalSum int
	for _, num := range nums {
		totalSum += num
	}
	left, right, maxSum := 0, len(nums)-1, totalSum
	for left != right {
		if nums[left] < nums[right] {
			newSum := totalSum - nums[left]
			maxSum = maxInt(maxSum, newSum)
			left += 1
		}
		if nums[left] >= nums[right] {
			newSum := totalSum - nums[right]
			maxSum = maxInt(maxSum, newSum)
			right -= 1
		}
	}
	return maxSum
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
