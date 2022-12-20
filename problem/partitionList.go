package problem

// PartitionList
// Time complexity O(n)
// Space complexity O(n)
func PartitionList(nums []int, k int) []int {
	var left, right []int
	for _, num := range nums {
		switch num > k {
		case true:
			right = append(right, num)
		case false:
			left = append(left, num)
		}
	}

	for _, num := range right {
		left = append(left, num)
	}

	return left
}

func PartitionListOptimal(nums []int, k int) []int {
	low, high := 0, len(nums)-1
	var i int
	for i <= high {
		n := nums[i]
		switch true {
		case n > k:
			nums[high], nums[i] = nums[i], nums[high]
		case n < k:
			nums[low], nums[i] = nums[i], nums[low]
		case n == k:
			i += 1
		}
	}

	return nums
}
