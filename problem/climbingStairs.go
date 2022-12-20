package problem

//ClimbingStairs
//Time complexity O(n)
//Space complexity O(1)
func ClimbingStairs(n int) int {
	prev, prevprev, current := 1, 1, 0
	for i := 2; i < n+1; i++ {
		current = prev + prevprev
		prevprev = prev
		prev = current
	}
	return current
}
