package problem

// Permutation compose all the possible permutation of a given array of strings
// Can be change to operate with integers
// Time complexity: O(n!)
// Space complexity: O(n!)
// @todo to re-do, i don't really understand how it works
func Permutation(nums []string, f func([]string)) {
	perm(nums, f, 0)
}

func perm(nums []string, f func([]string), i int) {
	if i > len(nums) {
		f(nums)
		return
	}
	perm(nums, f, i+1)
	for j := i + 1; j < len(nums); j++ {
		nums[i], nums[j] = nums[j], nums[i]
		perm(nums, f, i+1)
		nums[i], nums[j] = nums[j], nums[i]
	}

}
