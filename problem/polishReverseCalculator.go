package problem

import "strconv"

func ReversePolishCalculator(inputs ...string) int {
	var stack []int

	pop := func(stack []int) (int, []int) {
		poppedVal := stack[len(stack)-1]
		return poppedVal, stack[:len(stack)-1]
	}

	for _, i := range inputs {
		switch i {
		case "-", "+", "*", "/":
			var a, b int
			a, stack = pop(stack)
			b, stack = pop(stack)
			switch i {
			case "-":
				stack = append(stack, a-b)
			case "+":
				stack = append(stack, a+b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		default:
			cvrSymbol, _ := strconv.Atoi(i)
			stack = append(stack, cvrSymbol)
		}
	}

	return stack[0]
}
