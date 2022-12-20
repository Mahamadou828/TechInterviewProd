package problem

type trieNodeWithCount struct {
	count    int
	children map[string]*trieNodeWithCount
}

func (t trieNodeWithCount) Insert(word string) {
	current := &t

	for _, c := range word {
		if _, ok := current.children[string(c)]; !ok {
			current.children[string(c)] = &trieNodeWithCount{
				count:    0,
				children: map[string]*trieNodeWithCount{},
			}
		}
		current = current.children[string(c)]
		current.count += 1
	}
}

func (t trieNodeWithCount) UniquePrefix(word string) string {
	current := &t
	var prefix string

	for _, c := range word {
		if current.count == 1 {
			return prefix
		} else {
			current = current.children[string(c)]
			prefix += string(c)
		}
	}

	return prefix
}

func ShortestUniquePrefix(words ...string) []string {
	var res []string

	trie := trieNodeWithCount{
		count:    0,
		children: map[string]*trieNodeWithCount{},
	}

	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		res = append(res, trie.UniquePrefix(word))
	}

	return res
}
