package problem

// GenerateParentheses generate all possible parenthesis combinations
// Time complexity O(2^n)
// Space complexity O(n)
func GenerateParentheses(n int) []string {
	return generate(n, 0, 0, "")
}

func generate(n, left, right int, str string) []string {
	if left+right == 2*n {
		return []string{str}
	}

	var result []string
	if left < n {
		result = append(result, generate(n, left+1, right, str+"(")...)
	}
	if right < left {
		result = append(result, generate(n, left, right+1, str+")")...)
	}

	return result
}
