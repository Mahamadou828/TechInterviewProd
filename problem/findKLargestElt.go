package problem

import (
	dts "github.com/Mahamadou828/coderpro/datastruct"
	"math"
)

// FindKLargestElt return k largest element in an array
// Time complexity O(n + k log n)
// O(n) to construct the heap
// O(k log n ) because pop take log(n) and we do that k time
// Space complexity O(n)
func FindKLargestElt(nums []int, k int) []int {
	heap, res := dts.CreateNewIntHeap(nums), []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop().(int))
	}
	return res
}

// FindKSmallestElt return k smallest elements in an array
// Time complexity O(k.n)
// Space complexity O(k)
func FindKSmallestElt(nums []int, k int) int {
	var res []int

	for i := 0; i < k; i++ {
		currentMin, minIdx := math.MaxInt, -1
		for idx, num := range nums {
			if num < currentMin {
				currentMin, minIdx = num, idx
			}
		}

		res = append(res, currentMin)
		nums[minIdx] = math.MaxInt
	}

	return res[len(res)-1]
}

func FindKLargestEltWithQS(nums []int, k int) int {
	left, right := 0, len(nums)-1
	k = len(nums) - k
	partition := func(nums []int, low, high int) int {
		pivot := nums[high]
		i := low
		for j := low; j < high; j++ {
			if nums[j] <= pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i += 1
			}
		}
		nums[i], nums[high] = nums[high], nums[i]
		return i
	}

	for left <= right {
		pivotIdx := partition(nums, left, right)
		if pivotIdx == k {
			return nums[pivotIdx]
		} else if pivotIdx > k {
			right = pivotIdx - 1
		} else {
			left = pivotIdx + 1
		}
	}
	return -1
}
