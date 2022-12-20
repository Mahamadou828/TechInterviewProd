package problem

var mapCharacterCode = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"g": 6,
	"f": 7,
}

//GroupAnagram groups all anagram together. Actually working only for letter between a-f
//Time complexity  O(n * m )
//Space complexity O(n)
//Where n is the number of word and m is the number of character in each word
func GroupAnagram(words []string) [][]string {
	encodedWords := make([]int, len(words))
	mapEncodedWords := make(map[int][]string)

	for i, word := range words {
		encodedWords[i] = hashString(word)
	}

	for i, point := range encodedWords {
		res, ok := mapEncodedWords[point]
		if !ok {
			mapEncodedWords[point] = []string{words[i]}
		}
		mapEncodedWords[point] = append(res, words[i])
	}
	var result [][]string

	for _, anagram := range mapEncodedWords {
		result = append(result, anagram)
	}

	return result
}

func hashString(s string) int {
	var total int
	for _, c := range s {
		point := mapCharacterCode[string(c)]
		total += point
	}
	return total
}
