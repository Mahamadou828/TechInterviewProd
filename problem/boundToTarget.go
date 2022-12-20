package problem

//BoundToTarget find the first and last element that are near the target
//Time complexity O(n)
//Space complexity O(1)
func BoundToTarget(arr []int, target int) []int {
	leftPtr, rightPtr := 0, len(arr)-1

	for leftPtr < rightPtr && rightPtr > leftPtr {
		left, right := arr[leftPtr], arr[rightPtr]
		if left == target && right == target {
			break
		}
		if left != target {
			leftPtr += 1
		}
		if right != target {
			rightPtr -= 1
		}
	}

	return []int{leftPtr - 1, rightPtr + 1}
}

//@todo to implement
//func BoundToTargetBS(arr []int, target int) []int {
//	leftPtr, rightPtr, middlePtr := 0, len(arr)-1, (len(arr)-1)/2
//	for leftPtr < rightPtr && rightPtr > leftPtr {
//		left, rigth, middle := arr[leftPtr], arr[rightPtr], arr[middlePtr]
//		if left < middle {
//
//		}
//	}
//}
