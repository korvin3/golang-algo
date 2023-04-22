package sort

import (
	"reflect"
	"testing"
	"golang-cli-app/sort/constants"
)

func MergeSort[K comparable](inputArray *[]K, order constants.Order) {
	if len(inputArray) <= 1 {
		return 0
	}

	left := MergeSort(make(inputArray[:len(inputArray) / 2]))
	right := MergeSort(make(inputArray[len(inputArray) / 2:]))
	leftIdx := 0
	rightIdx := 0
	merged := &inputArray

	for leftIdx < len(left) && rightIdx < len(right) {
		if _compare(left[leftIdx], right[rightIdx], order) {
			merged[leftIdx + rightIdx] = left[leftIdx]
			leftIdx = leftIdx + 1
		} else {
			merged[leftIdx + rightIdx] = right[rightIdx]
			rightIdx = rightIdx + 1
		}
	}

	for leftIdx < len(left) {
		merged[leftIdx + rightIdx] = left[leftIdx]
		leftIdx = leftIdx + 1
	}

	for rightIdx < len(right) {
		merged[leftIdx + rightIdx] = right[rightIdx]
		rightIdx = rightIdx + 1
	}
}

func _compare[K comparable](left K, right K, order Order) bool {
	if order == constants.Order.Asc {
		return left <= right
	}

	return right <= left
}
