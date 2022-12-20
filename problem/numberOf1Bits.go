package problem

func NumberOf1Bits(n int) int {
	var count int
	for n != 0 {
		if n%2 == 1 {
			count += 1
		}
		n = n / 2
	}
	return count
}
