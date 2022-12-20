package problem

//@todo to-redo, just understand how it works

//ShortColor
//Time complexity O(n)
//Space complexity O(1)
func ShortColor(colors []int) []int {
	blackPtr, purplePtr, currentPtr := 0, len(colors)-1, 0

	for currentPtr <= purplePtr {
		switch colors[currentPtr] {
		case 0:
			colors[blackPtr], colors[currentPtr] = colors[currentPtr], colors[blackPtr]
			blackPtr += 1
			currentPtr += 1
		case 2:
			colors[purplePtr], colors[currentPtr] = colors[currentPtr], colors[purplePtr]
			purplePtr -= 1
		default:
			currentPtr += 1
		}
	}

	return colors
}
