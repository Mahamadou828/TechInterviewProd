package problem

var mapRomanNumber = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

//RomanNumber
//Time complexity O(n)
//Space complexity O(1)
func RomanNumber(str string) int {
	var prev int
	var sum int
	for j := len(str) - 1; j >= 0; j-- {
		char := string(str[j])
		current, ok := mapRomanNumber[char]
		if !ok {
			continue
		}
		if prev > current {
			sum -= current
		} else {
			sum += current
		}
		prev = current
	}
	return sum
}
