package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
	"strings"
)

// FindSubtree
// Time complexity O(n)
// Space complexity O(n)
func FindSubtree(a, b *dts.Tree) bool {
	return findSubtreeHelper(a.Root, b.Root)
}

func findSubtreeHelper(a, b *dts.Leaf) bool {
	if a == nil {
		return false
	}

	isMatch := a.Value == b.Value
	if isMatch {
		isMatchLeft := findSubtreeHelper(a.Left, b.Left)
		if isMatchLeft {
			isMatchRight := findSubtreeHelper(a.Right, b.Right)
			if isMatchRight {
				return true
			}
		}
	}

	return findSubtreeHelper(a.Left, b) || findSubtreeHelper(a.Right, b)
}

// FindSubtreeBySerialization
// Time complexity O(m*n)
// Space complexity O(1)
func FindSubtreeBySerialization(a, b *dts.Tree) bool {
	aSer, bSer := TreeSerialization(a), TreeSerialization(b)

	return strings.Contains(aSer, bSer)
}
