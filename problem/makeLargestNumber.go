package problem

import (
	"sort"
	"strings"
)

func MakeLargestNumber(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		a, b := strs[i], strs[j]

		return (a + b) > (b + a)
	})

	return strings.Join(strs, "")
}

//7772452
//77245217
