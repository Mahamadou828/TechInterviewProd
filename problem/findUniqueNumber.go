package problem

//FindUniqueNumber find the unique number in the provided list
//Space complexity O(n)
//Time complexity O(n)
func FindUniqueNumber(nums []int) int {
	numberCountMap := make(map[int]int)

	for _, num := range nums {
		numberCountMap[num] = numberCountMap[num] + 1
	}

	for key, val := range numberCountMap {
		if val == 1 {
			return key
		}
	}

	return -1
}

//FindUniqueNumber2 is a variation of FindUniqueNumber
//Time complexity O(n)
//Space complexity O(1)
func FindUniqueNumber2(nums []int) int {
	var unique int
	for _, k := range nums {
		unique ^= k
	}
	return unique
}
