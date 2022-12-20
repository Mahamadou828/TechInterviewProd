package problem

import "container/heap"

type pointHeap [][]int

func (h pointHeap) Len() int { return len(h) }
func (h pointHeap) Less(i, j int) bool {
	return (h[i][0] * h[i][1]) > (h[j][0] * h[j][1])
}
func (h pointHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *pointHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *pointHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *pointHeap) Peek() any {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}

func createNewPointHeap(nums [][]int) *pointHeap {
	h := &pointHeap{}
	for _, num := range nums {
		h.Push(num)
	}
	heap.Init(h)

	return h
}

// ClosestPointsToTheOrigin
// Time complexity kO(log n)
// Space complexity O(n)
func ClosestPointsToTheOrigin(points [][]int, k int) [][]int {
	heap := createNewPointHeap(points)
	var res [][]int
	for i := 0; i < k; i++ {
		point := heap.Pop()
		res = append(res, point.([]int))
	}
	return res
}
