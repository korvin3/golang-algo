package sort

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[K constraints.Ordered](inputArray []K, order Order)[]K {
	if len(inputArray) <= 1 {
		return inputArray
	}

	left := inputArray[:len(inputArray) / 2]
	left = MergeSort(left, order)
	right := inputArray[len(inputArray) / 2:]
	right = MergeSort(right, order)
	leftIdx := 0
	rightIdx := 0
	resultArray := make([]K, len(inputArray))

	for leftIdx < len(left) && rightIdx < len(right) {
		if _compare(left[leftIdx], right[rightIdx], order) {
			resultArray[leftIdx + rightIdx] = left[leftIdx]
			leftIdx += 1
		} else {
			resultArray[leftIdx + rightIdx] = right[rightIdx]
			rightIdx += 1
		}
	}

	for leftIdx < len(left) {
		resultArray[leftIdx + rightIdx] = left[leftIdx]
		leftIdx += 1
	}

	for rightIdx < len(right) {
		resultArray[leftIdx + rightIdx] = right[rightIdx]
		rightIdx += 1
	}
	
	return resultArray
}

func _compare[K constraints.Ordered](left K, right K, order Order) bool {
	if order == Asc {
		return left < right
	}

	return right <= left
}
