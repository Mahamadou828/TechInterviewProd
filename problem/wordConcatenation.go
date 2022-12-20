package problem

import "strings"

//WordConcatenation
//Time complexity O(n^2)
//Space complexity O(n)
func WordConcatenation(words []string) []string {
	var res []string
	mapWord, mapRes := make(map[string]string), make(map[string]string)

	for _, word := range words {
		mapWord[word] = word
	}

	for i, word := range words {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}

			word2 := words[j]

			if strings.Contains(word2, word) {
				subWord := strings.Split(word2, word)
				isWord := true
				for _, word := range subWord {
					if _, ok := mapWord[word]; !ok && len(word) > 0 {
						isWord = false
						break
					}
				}
				if isWord {
					mapRes[word2] = word2
				}
			}
		}
	}

	for _, val := range mapRes {
		res = append(res, val)
	}

	return res
}
