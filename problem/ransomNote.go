package problem

// CanSpell take an array of strings and a word and return if it's possible to spell the word using the
// letter inside the array
// Time complexity: O(n+m), Space complexity: O(1)
// Where n is the length of the word and m is the length of the array
func CanSpell(mgz []string, word string) bool {
	mgzLetterMap := make(map[string]int)
	for _, letter := range mgz {
		mgzLetterMap[letter] += 1
	}
	for _, letter := range word {
		if val, ok := mgzLetterMap[string(letter)]; !ok || val <= 0 {
			return false
		}
		mgzLetterMap[string(letter)] -= 1
	}
	return true
}
