package sort

import (
	"reflect"
	"testing"
	"golang-cli-app/sort/constants"
)

func TestMergeSort(t *testing.T) {
	for _, tc := range _testCases() {
		t.Run(tc.name, func(t *testing.T) {
			MergeSort(tc.input, constants.Asc)
			if !reflect.DeepEqual(tc.input, tc.expected) {
				t.Errorf("Expected %v but got %v", tc.expected, tc.input)
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
			name:     "sorted array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "reverse sorted array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "unsorted array",
			input:    []int{5, 1, 3, 2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "array with duplicate elements",
			input:    []int{5, 1, 3, 2, 4, 2},
			expected: []int{1, 2, 2, 3, 4, 5},
		},
	}
}