package problem

func ValidParentheses(str string) bool {
	parens := map[string]string{
		"[": "]",
		"{": "}",
		"(": ")",
	}
	invParens := make(map[string]string)
	for k, v := range parens {
		invParens[v] = k
	}

	stack := []string{}

	for _, c := range str {
		if _, ok := parens[string(c)]; ok {
			stack = append(stack, string(c))
		} else if _, ok := invParens[string(c)]; ok {
			invElt := invParens[string(c)]
			if len(stack) == 0 || stack[len(stack)-1] != invElt {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
