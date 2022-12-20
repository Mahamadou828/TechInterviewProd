package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
)

func OptimiseListSum(nums []int, from, to int) int {
	ll := dts.NewLinkedList(nums...)
	mapIndexNode := make(map[int]*dts.LLNode)
	var idx int
	ll.Traverse(func(node *dts.LLNode) {
		mapIndexNode[idx] = node
		idx++
	})
	current := mapIndexNode[from]
	var sum int

	for i := 0; i < to-from; i++ {
		sum += current.Value
		current = current.Next
	}

	return sum
}

func OptimiseListSumPerf(nums []int, from, to int) int {
	accumulatedSum := []int{0}
	var accumulatedSumCalculus int

	for _, num := range nums {
		accumulatedSumCalculus += num
		accumulatedSum = append(accumulatedSum, accumulatedSumCalculus)
	}

	return accumulatedSum[to] - accumulatedSum[from]
}
