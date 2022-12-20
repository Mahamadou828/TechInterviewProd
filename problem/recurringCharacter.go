package problem

// RecurringCharacter
// Time complexity O(n)
// Space complexity O(n)
func RecurringCharacter(str string) string {
	mapChar := make(map[string]int)

	for _, c := range str {
		if _, ok := mapChar[string(c)]; ok {
			return string(c)
		}
		mapChar[string(c)] += 1
	}

	return ""
}
