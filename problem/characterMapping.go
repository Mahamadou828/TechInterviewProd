package problem

//CharacterMapping
//Time complexity O(n)
//Space complexity O(n)
func CharacterMapping(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	chars := map[uint8]uint8{}

	for i := 0; i < len(str1); i++ {
		if _, ok := chars[str1[i]]; !ok {
			chars[str1[i]] = str2[i]
		} else {
			if chars[str1[i]] != str2[i] {
				return false
			}
		}
	}

	return true
}
