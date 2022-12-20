package problem

// Multitasking
// Time complexity O(n)
// Space complexity O(n)
func Multitasking(jobs []int, cooldown int) int {
	lastPos, current := map[int]int{}, 0

	for _, task := range jobs {
		if taskPos, ok := lastPos[task]; ok {
			if current-taskPos <= cooldown {
				current = cooldown + taskPos + 1
			}
		}
		lastPos[task] = current
		current = current + 1
	}
	return current
}
