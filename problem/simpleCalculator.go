package problem

import (
	"strconv"
	"unicode"
)

func SimpleCalculator(expression string) int {
	r, _ := simpleCalculatorHelper(expression, 0)
	return r
}

func simpleCalculatorHelper(expression string, index int) (int, int) {
	op, result := '+', 0
	for index < len(expression) {
		char := expression[index]
		if char == '+' || char == '-' {
			op = int32(char)
		} else {
			value := 0
			if unicode.IsDigit(rune(char)) {
				letter, _ := strconv.Atoi(string(char))
				value = letter
			} else if char == '(' {
				value, index = simpleCalculatorHelper(expression, index+1)
			}
			if op == '+' {
				result += value
			}
			if op == '-' {
				result -= value
			}
		}
		index += 1
	}
	return result, index
}
