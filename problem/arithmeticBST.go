package problem

import dts "github.com/Mahamadou828/coderpro/datastruct"

const (
	Addition       float64 = 10100
	Subtraction    float64 = 10101
	Multiplication float64 = 10111
	Division       float64 = 11111
)

var mapCommandsFunction = map[float64]func(a, b float64) float64{
	Addition:       func(a, b float64) float64 { return a + b },
	Subtraction:    func(a, b float64) float64 { return a - b },
	Multiplication: func(a, b float64) float64 { return a * b },
	Division:       func(a, b float64) float64 { return a / b },
}

// ArithmeticBST
// Time complexity O(n)
// Space complexity O(n)
// the arithmetic commands are represents in number ( to avoid generic bst )
func ArithmeticBST(tree *dts.Tree) float64 {
	return arithmeticBSTHelper(tree.Root)
}

func arithmeticBSTHelper(leaf *dts.Leaf) float64 {
	if leaf == nil {
		return 0
	}
	fn, ok := mapCommandsFunction[leaf.Value]
	if !ok {
		return leaf.Value
	}
	return fn(arithmeticBSTHelper(leaf.Left), arithmeticBSTHelper(leaf.Right))
}
