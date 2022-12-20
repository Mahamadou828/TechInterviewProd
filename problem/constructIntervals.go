package problem

import "fmt"

//ConstructIntervals constructs all existing intervals from a list of number.
//Time complexity O(n)
//Space complexity O(n)
func ConstructIntervals(nums []int) []string {
	left, right := 0, 1
	var res []string
	for right < len(nums) {
		if nums[right]-nums[right-1] > 1 {
			res = append(res, fmt.Sprintf("%d-%d", nums[left], nums[right-1]))
			left = right
		}
		if right == len(nums)-1 {
			res = append(res, fmt.Sprintf("%d-%d", nums[left], nums[right]))
		}
		right += 1
	}

	return res
}
