package problem

// WordSearch search a word inside the matrix
// Time complexity O(n^3)
// Space complexity O(k)
// Wher k is the length of the word
func WordSearch(arr [][]string, word string) bool {
	length, depth := len(arr[0]), len(arr)

	for i := 0; i < depth; i++ {
		for j := 0; j < length; j++ {
			if checkHorizontal(arr, word, i, j) || checkVertical(arr, word, i, j) {
				return true
			}
		}
	}
	return false
}

func checkHorizontal(arr [][]string, word string, i, j int) bool {
	return check(arr[i][j:], word)
}

func checkVertical(arr [][]string, word string, i, j int) bool {
	var toCheck []string

	for k := i; k < len(arr); k++ {
		toCheck = append(toCheck, arr[k][j])
	}

	return check(toCheck, word)
}

func check(toCheck []string, word string) bool {
	if len(toCheck) != len(word) {
		return false
	}
	for i := 0; i < len(toCheck); i++ {
		if toCheck[i] != string(word[i]) {
			return false
		}
	}
	return true
}
