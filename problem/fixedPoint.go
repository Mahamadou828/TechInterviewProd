package problem

// FixedPoint
// Time Complexity O(log n)
// Space Complexity O(1)
func FixedPoint(nums []int) int {
	low, high := 0, len(nums)

	for low != high {
		mid := (low + high) / 2
		switch true {
		case nums[mid] == mid:
			return mid
		case nums[mid] < mid:
			low = mid + 1
		default:
			high = mid
		}
	}

	return -1
}
