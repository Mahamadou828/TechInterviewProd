package problem

func Fibonacci(n int) int {
	a, b := 0, 1
	switch true {
	case a == n:
		return a
	case b == n:
		return b
	}

	var value int

	for i := 2; i < n+1; i++ {
		value = a + b
		a = b
		b = value
	}
	
	return value
}
