package problem

import "math"

// MinSubArrayLen return the minimum length of the sub array where the sum of all elements is >= to k
// Time complexity O(n)
// Space complexity O(1)
func MinSubArrayLen(nums []int, k int) int {
	leftPtr, rightPtr := 0, 0
	sum, minLength := 0, math.MaxInt

	for rightPtr < len(nums) {
		sum += nums[rightPtr]

		for sum >= k {
			minLength = int(math.Min(float64(minLength), float64(rightPtr-leftPtr+1)))
			sum -= nums[leftPtr]
			leftPtr += 1
		}

		rightPtr += 1
	}

	if minLength == math.MaxInt {
		return 0
	}

	return minLength
}
