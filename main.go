package main

import (
	"fmt"
	"github.com/Mahamadou828/coderpro/problem"
)

//root := &dts.Leaf{
//	Value: 1,
//	Left:  &dts.Leaf{Value: 2, Right: &dts.Leaf{Value: 4}},
//	Right: &dts.Leaf{Value: 3, Left: &dts.Leaf{Value: 5}, Right: &dts.Leaf{Value: 6}},
//}

func main() {
	//root := &dts.Leaf{Value: 4}
	//leaf1, leaf2, leaf3 := &dts.Leaf{Value: 2}, &dts.Leaf{Value: 1}, &dts.Leaf{Value: 8}
	//leaf4, leaf5, leaf6 := &dts.Leaf{Value: 5}, &dts.Leaf{Value: 9}, &dts.Leaf{Value: 9}
	//
	//root.Left, root.Right = leaf1, leaf3
	//leaf1.Parent, leaf3.Parent = root, root
	//
	//leaf1.Left, leaf2.Parent = leaf2, leaf1
	//
	//leaf3.Left, leaf3.Right, leaf4.Parent, leaf5.Parent = leaf4, leaf5, leaf3, leaf3
	//
	//leaf4.Right, leaf6.Parent = leaf6, leaf4
	//
	//tree := &dts.Tree{Root: root}
	fmt.Println(problem.WordConcatenation([]string{"cat", "cats", "dog", "catsdog"}))
}
