package problem

///@todo to redo

type trieNode struct {
	isWord   bool
	children map[string]*trieNode
}

func buildTries(words []string) *trieNode {
	trie := &trieNode{isWord: false, children: map[string]*trieNode{}}
	for _, word := range words {
		current := trie
		for _, char := range word {
			if _, ok := current.children[string(char)]; !ok {
				current.children[string(char)] = &trieNode{isWord: false, children: map[string]*trieNode{}}
			}
			current = current.children[string(char)]
		}
		current.isWord = true
	}
	return trie
}

func Autocompletion(words []string, word string) []string {
	trie := buildTries(words)
	current := trie
	var res []string

	for _, char := range word {
		if _, ok := current.children[string(char)]; !ok {
			return res
		}
		current = current.children[string(char)]
	}
	autocompletionHelper(current, word, func(s string) {
		res = append(res, s)
	})
	return res
}

func autocompletionHelper(node *trieNode, prefix string, append func(string)) {
	if node.isWord {
		append(prefix)
	}
	for char := range node.children {
		autocompletionHelper(node.children[char], prefix+char, append)
	}
}

func autocompletionHelperItr(node *trieNode, prefix string, words []string) {
	type item struct {
		node   *trieNode
		prefix string
	}
	stack := []item{{node, prefix}}

	for len(stack) > 0 {
		popedItem, stack := stack[len(stack)-1], stack[:len(stack)-1]
		if popedItem.node.isWord {
			words = append(words, popedItem.prefix)
		}
		for char := range popedItem.node.children {
			child := node.children[char]
			stack = append(stack, item{child, prefix + char})
		}
	}
}
