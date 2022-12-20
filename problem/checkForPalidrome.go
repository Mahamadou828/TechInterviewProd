package problem

func CheckPalindrome(str string) string {
	charCounts := map[string]int{}

	for _, c := range str {
		charCounts[string(c)] += 1
	}

	var pal, oddChar string

	for c, cnt := range charCounts {
		if cnt%2 == 0 {
			for i := 0; i < cnt/2; i++ {
				pal += c
			}
		} else if oddChar == "" {
			oddChar = c
			for i := 0; i < cnt/2; i++ {
				pal += c
			}
		} else {
			return ""
		}
	}

	return pal + oddChar + reverse(pal)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
