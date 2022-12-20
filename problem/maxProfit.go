package problem

import "math"

// MaxProfit
// Time complexity O(n)
// Space complexity O(1)
func MaxProfit(nums []int) int {
	if len(nums) == 2 {
		return nums[0] + nums[1]
	}
	leftPtr, rightPtr := 0, 1
	maxProfit, currentProfit := math.MinInt, 0

	for leftPtr < len(nums) && rightPtr < len(nums) {
		left, right := nums[leftPtr], nums[rightPtr]
		profit := -left + right
		if profit >= maxProfit {
			maxProfit, currentProfit = profit, profit
		}
		if profit < currentProfit {
			leftPtr, currentProfit = rightPtr, profit
		}
		rightPtr += 1
	}

	return maxProfit
}
