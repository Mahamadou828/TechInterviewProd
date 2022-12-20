package problem

func SortedSquareNbrs(nums []int) []int {
	negI, i := 1, 0
	var res []int

	for _, n := range nums {
		if n >= 0 {
			if negI == -1 {
				negI = i - 1
			}

			for negI >= 0 && nums[negI] < 0 && -nums[negI] < nums[i] {
				val := nums[negI]
				res = append(res, val*val)
				negI -= 1
			}

			val := nums[i]
			res = append(res, val*val)
		}
		i += 1
	}
	if len(res) == 0 && negI == -1 {
		negI = len(nums) - 1
	}

	for negI >= 0 && nums[negI] < 0 {
		val := nums[negI]
		res = append(res, val*val)
		negI -= 1
	}

	return res
}
