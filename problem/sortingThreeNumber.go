package problem

// SortingThreeNumber take an array compose of 3 unique numbers and sort the array
// Time complexity O(n)
// Space complexity O(1)
func SortingThreeNumber(nums []int, num1, num2, num3 int) []int {
	leftIdx, index, rightIdx := 0, 0, len(nums)-1
	for index <= rightIdx {
		switch nums[index] {
		case num1:
			nums[index], nums[leftIdx] = nums[leftIdx], nums[index]
			leftIdx, index = leftIdx+1, index+1
		case num2:
			index = index + 1
		case num3:
			nums[index], nums[rightIdx] = nums[rightIdx], nums[index]
			rightIdx -= 1
		}
	}
	return nums
}
