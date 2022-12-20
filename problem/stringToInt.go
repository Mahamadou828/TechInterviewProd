package problem

// StringToInt
// Time complexity O(n)
// Space complexity O(1)
func StringToInt(str string) int {
	sign, startIndex := 1, 0
	mapStrToInt := map[uint8]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}

	if str[0] == '-' {
		sign, startIndex = -1, 1
	}
	var sum int

	for i := startIndex; i < len(str)-startIndex; i++ {
		num := mapStrToInt[str[i]]
		sum = sum*10 + num
	}

	return sum * sign
}
