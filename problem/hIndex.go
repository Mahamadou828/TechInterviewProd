package problem

//HIndex
//Time complexity O(n)
//Space complexity O(n)
func HIndex(citations []int) int {
	n := len(citations)
	var freqs []int

	for i := 0; i < n+1; i++ {
		freqs = append(freqs, 0)
	}

	for _, c := range citations {
		if c >= n {
			freqs[n] += 1
		} else {
			freqs[c] += 1
		}
	}

	var total int
	i := n

	for i >= 0 {
		total += freqs[i]
		if total >= i {
			return i
		}
		i -= 1
	}

	return 0
}
