package algorithms_test

import (
	"testing"

	"github.com/anilpandey071999/go-algo/algorithms"
)

func TestLinearSearch(t *testing.T) {
	testCases := []struct {
		slice    []int
		target   int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 6, false},
	}

	for _, testCase := range testCases {
		result := algorithms.LinearSearch(testCase.slice, testCase.target)
		if result != testCase.expected {
			t.Errorf("LinearSearch(%v, %d) = %v; expected %v", testCase.slice, testCase.target, result, testCase.expected)
		}
	}
}
