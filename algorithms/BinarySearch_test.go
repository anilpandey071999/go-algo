package algorithms_test

import (
	"testing"

	"github.com/anilpandey071999/go-algo/algorithms"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		slice    []int
		target   int
		expected bool
	}{
		{[]int{5, 1, 4, 2, 3}, 1, true}, // Target (1) is present at a random position
		{[]int{5, 1, 4, 2, 3}, 6, false}, // Target (6) is not present
		{[]int{5, 1, 4, 2, 3}, 4, true}, // Target (4) is present
	}

	for _, testCase := range testCases {
		result := algorithms.BinarySearch(testCase.slice, testCase.target)
		if result != testCase.expected {
			t.Errorf("LinearSearch(%v, %d) = %v; expected %v", testCase.slice, testCase.target, result, testCase.expected)
		}
	}
}