package problem

import "sort"

//ThreeSum find three number that sum to zero
//Time complexity O(nlog n)
//Space complexity O(n)
func ThreeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	leftPtr, rightPtr, mapNums := 0, len(nums)-1, make(map[int]int)

	for idx, num := range nums {
		mapNums[num] = idx
	}

	for leftPtr != rightPtr && leftPtr < len(nums) && rightPtr > 0 {
		left, right := nums[leftPtr], nums[rightPtr]
		sum := left + right
		if idx, ok := mapNums[-sum]; ok && idx != leftPtr && idx != rightPtr {
			result = append(result, []int{left, -sum, right})
		}
		if sum < 0 {
			leftPtr += 1
		} else if sum > 0 {
			rightPtr -= 1
		} else {
			leftPtr += 1
			rightPtr -= 1
		}
	}

	return result
}
