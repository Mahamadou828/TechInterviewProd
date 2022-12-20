package problem

// ProductOfArrayExceptSelf https://www.techseries.dev/products/coderpro/categories/1831147/posts/6228785
// Time complexity O(n)
// Space complexity O(n)
func ProductOfArrayExceptSelf(nums []int) []int {
	product, right := 1, make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		right[i] = 1
	}
	for i := len(nums) - 2; i >= 0; i-- {
		product *= nums[i+1]
		right[i] = product
	}

	product = 1
	for i := 1; i < len(nums); i++ {
		product *= nums[i-1]
		right[i] *= product
	}

	return right
}
