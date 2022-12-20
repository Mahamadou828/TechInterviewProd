package problem

// ArrayIntersection
// Time complexity O(n)
// Space complexity O(n*m)
// Where n is the length of the first array and m the length of the second
func ArrayIntersection(a, b []int) []int {
	mapAElt, mapResult := make(map[int]bool), make(map[int]bool)

	for _, num := range a {
		mapAElt[num] = true
	}

	for _, num := range b {
		if _, ok := mapAElt[num]; ok {
			mapResult[num] = true
		}
	}

	var result []int

	for key, _ := range mapResult {
		result = append(result, key)
	}

	return result
}
