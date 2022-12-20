package problem

import (
	"container/heap"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

// IntervalMinHeap min heap for intervals and sorted by End
type IntervalMinHeap []Interval

func (h IntervalMinHeap) Len() int {
	return len(h)
}

// Less note: the best practice is to use pointer receiver for all methods if some methods need it
// but the actual code is not as readable.. many (*h) in the code
func (h IntervalMinHeap) Less(i, j int) bool {
	// min heap sort by end time, so the top is always the next meeting that will end
	return h[i].End < h[j].End
}

func (h IntervalMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

// Pop this is method that package heap will use for the internal data
// too add and remove things from heap, use heap.Push and heap.Pop
// since this function is for the heap package to use, this function actually remove and
// return the last element in the slice Len() - 1 instead of "pop head of heap"
func (h *IntervalMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

//@todo to-redo not my solution i don't understand what's going on here

func MeetingRooms(intervals []Interval) int {
	if intervals == nil || len(intervals) == 0 {
		return 0
	}

	intervalsOrderedByStart := make([]Interval, len(intervals))
	copy(intervalsOrderedByStart, intervals)
	sort.Slice(intervalsOrderedByStart, func(i, j int) bool {
		return intervalsOrderedByStart[i].Start < intervalsOrderedByStart[j].Start
	})

	h := &IntervalMinHeap{}
	heap.Init(h)

	for i := 0; i < len(intervalsOrderedByStart); i++ {
		// note: when the heap is empty, add to heap and you are done
		// remember to add continue, otherwise later logic will add it twice
		if h.Len() == 0 {
			heap.Push(h, intervalsOrderedByStart[i])
			continue
		}

		// the next meeting that about to stop
		m := (*h)[0]
		// note: we need <= instead of < here, because for back to back meetings
		// we don't need additional room
		if m.End <= intervalsOrderedByStart[i].Start {
			// the next meeting that is about to stop can stop now
			heap.Pop(h)
		}
		heap.Push(h, intervalsOrderedByStart[i])
	}

	// whatever number of meeting in the heap is the number of conference room we need
	// because at the time that meeting is about to start, there is no elems meeting
	// room available for it, we will need additional meeting room, since we didn't pop
	// in this case, the size of heap increased too.
	return h.Len()
}
