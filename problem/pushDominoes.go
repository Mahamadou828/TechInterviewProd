package problem

//@todo to redo

// PushDominoes re-oriente the direction of dominoes.
//Time complexity O(n)
//Space complexity O(n)
func PushDominoes(dominoes string) string {
	forces, maxForce, force := make([]int, len(dominoes)), len(dominoes), 0
	for i, d := range dominoes {
		switch d {
		case 'R':
			force = maxForce
		case 'L':
			force = 0
		default:
			force = max(0, force-1)
		}
		forces[i] = force
	}

	for i := len(dominoes) - 1; i > 0; i-- {
		d := dominoes[i]
		switch d {
		case 'L':
			force = maxForce
		case 'R':
			force = 0
		default:
			force = max(0, force-1)
		}
		forces[i] -= force
	}

	result := ""
	for _, f := range forces {
		if f == 0 {
			result += "."
		} else if f > 0 {
			result += "R"
		} else {
			result += "L"
		}
	}

	return result
}

func max(min, val int) int {
	if val-1 < 0 {
		val = 0
	} else {
		val = val - 1
	}

	return val
}
