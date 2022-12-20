package problem

var mapLetterToNumber = map[string]int{
	"A": 2,
	"B": 2,
	"C": 2,
	"D": 3,
	"E": 3,
	"F": 3,
	"G": 4,
	"H": 4,
	"I": 4,
	"J": 5,
	"K": 5,
	"L": 5,
	"M": 6,
	"N": 6,
	"O": 6,
	"P": 7,
	"Q": 7,
	"R": 7,
	"S": 7,
	"T": 8,
	"U": 8,
	"V": 8,
	"W": 9,
	"Y": 9,
	"X": 9,
	"Z": 9,
}

// PhoneNumber
// Time complexity O(n^2)
// Space complexity O(n)
func PhoneNumber(nums []int, dict []string) []string {
	var res []string
	mapNumsCounter := make(map[int]int)
	for _, num := range nums {
		mapNumsCounter[num] += 1
	}

	for _, word := range dict {
		seen := map[int]int{}
		isOk := true
		for _, char := range word {
			if ok := phoneNumberHelper(string(char), mapNumsCounter, seen); !ok {
				isOk = false
			}
		}
		if isOk {
			res = append(res, word)
		}
	}

	return res
}

func phoneNumberHelper(char string, mapNums map[int]int, seen map[int]int) bool {
	charNum, ok := mapLetterToNumber[char]
	if !ok {
		return false
	}
	if _, ok := mapNums[charNum]; !ok {
		return false
	}
	if count, ok := seen[charNum]; ok {
		mapNumsCounter := mapNums[charNum]
		if count > mapNumsCounter {
			return false
		}
	}
	seen[charNum] += 1
	return true
}
