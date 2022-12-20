package problem

func DistributeBonuses(nums []int) []int {
	count := len(nums)
	var bonuses []int

	for i := 0; i < count; i++ {
		bonuses = append(bonuses, 1)
	}

	for i := 1; i < count; i++ {
		if nums[i-1] < nums[i] {
			bonuses[i] = bonuses[i-1] + 1
		}
	}

	for i := count - 2; i >= 0; i-- {
		if nums[i+1] < nums[i] {
			bonuses[i] = maxInt(bonuses[i], bonuses[i+1]+1)
		}
	}

	return bonuses
}
