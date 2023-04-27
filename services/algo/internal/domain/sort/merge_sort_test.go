package sort

import (
	"reflect"
	"testing"
	"github.com/korvin3/golang-algo/services/algo/internal/domain/sort/constants"
)

func TestMergeSort(t *testing.T) {
	for _, tc := range _testCases() {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeSort(tc.input, constants.Asc)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, result)
			}
		})
	}
}

func _testCases() []struct {
	name     string
	input    []int
	expected []int
} {
	return []struct {
		name     string
		input    []int
		expected []int
	} {
		{
			name: "sorted array",
			input: []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse sorted array",
			input: []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "unsorted array",
			input: []int{5, 1, 3, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "array with duplicate elements",
			input: []int{5, 1, 3, 2, 4, 2},
			expected: []int{1, 2, 2, 3, 4, 5},
		},
		{
			name: "empty array",
			input: []int{},
			expected: []int{},
		},
		{
			name: "array with single element",
			input: []int{5},
			expected: []int{5},
		},
		{
			name: "array with all same elements",
			input: []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name: "array with large number of elements",
			input: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "array with negative elements",
			input: []int{-3, 5, -1, 0, -2, 4, 2, 1, -4, 3},
			expected: []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
		{
			name: "array with decimal values",
			input: []int{5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
	}
}