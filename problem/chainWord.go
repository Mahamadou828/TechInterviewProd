package problem

//@todo to redo

// ChainWord
// Time complexity O(n!)
// Space complexity O(n!)
func ChainWord(words []string) bool {
	mapSymbol := make(map[string]string)
	for _, word := range words {
		mapSymbol[string(word[0])] = word
	}

	return chainWordHelper(mapSymbol, words[0], words[0], len(words), map[string]bool{})
}

func chainWordHelper(symbol map[string]string, currentWord, startWord string, length int, visited map[string]bool) bool {
	if length == 1 {
		return startWord[0] == currentWord[len(currentWord)-1]
	}
	visited[currentWord] = true

	for _, neighor := range symbol {
		if _, ok := visited[neighor]; !ok {
			return chainWordHelper(symbol, neighor, startWord, length-1, visited)
		}
	}

	delete(visited, currentWord)

	return false
}
