package problem

// AnagramsInString
// Time complexity O(n)
// Space complexity O(k) where k is the len of b
func AnagramsInString(a string, b string) []int {
	mapChar := map[string]int{}

	for _, char := range b {
		mapChar[string(char)] += 1
	}

	var results []int

	for i := 0; i < len(a); i++ {
		char := string(a[i])

		if i >= len(b) {
			charOld := string(a[i-len(b)])
			mapChar[charOld] += 1
			if mapChar[charOld] == 0 {
				delete(mapChar, charOld)
			}
		}

		mapChar[char] -= 1
		if mapChar[char] == 0 {
			delete(mapChar, char)
		}

		if i+1 >= len(b) && len(mapChar) == 0 {
			results = append(results, i-len(b)+1)
		}
	}

	return results
}
