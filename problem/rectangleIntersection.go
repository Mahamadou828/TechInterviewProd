package problem

type Rectangle struct {
	XMax int
	YMax int
	XMin int
	YMin int
}

func RectangleIntersection(a, b Rectangle) int {
	nR := Rectangle{
		XMin: max(a.XMin, b.XMin),
		YMin: max(a.YMin, b.YMin),
		XMax: min(a.XMax, b.XMax),
		YMax: min(a.YMax, b.YMax),
	}

	if nR.XMin >= nR.XMax || nR.YMin >= nR.YMax {
		return 0
	}

	return (nR.XMax - nR.XMin) * (nR.YMax - nR.YMin)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}