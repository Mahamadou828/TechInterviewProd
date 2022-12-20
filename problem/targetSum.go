package problem

func TragetSumm(arr []int, target int) (int, int) {
	leftPtr, rightPtr := 0, len(arr)-1
	for leftPtr != rightPtr {
		summ := arr[leftPtr] + arr[rightPtr]
		if summ == target {
			return arr[leftPtr], arr[rightPtr]
		} else if summ < target {
			leftPtr += 1
		} else if summ > target {
			rightPtr -= 1
		}
	}

	return 0, 0
}
