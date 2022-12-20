package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
)

// FrequentSubtreeSum
// Time complexity O(n)
// Space complexity O(n)
func FrequentSubtreeSum(tree *dts.Tree) float64 {
	var res float64
	mapFrequencySum := make(map[float64]float64)
	frequentSubtreeSumHelper(tree.Root, mapFrequencySum)
	var maxFrequency float64
	for sum, frequency := range mapFrequencySum {
		if frequency > maxFrequency {
			res, maxFrequency = sum, frequency
		}
	}

	return res
}

func frequentSubtreeSumHelper(leaf *dts.Leaf, mapFrequencySum map[float64]float64) float64 {
	if leaf == nil {
		return 0
	}
	sum := leaf.Value + frequentSubtreeSumHelper(leaf.Left, mapFrequencySum) + frequentSubtreeSumHelper(leaf.Right, mapFrequencySum)
	mapFrequencySum[sum] += 1
	return sum
}
