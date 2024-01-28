package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/anilpandey071999/go-algo/algorithms"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		slice    []int
		expected []int
	}{
		{[]int{5, 1, 4, 2, 3}, []int{1, 2, 3, 4, 5}},
		{[]int{10, 9, 8, 7, 6}, []int{6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // Already sorted
		{[]int{}, []int{}},                           // Empty slice
	}

	for _, testCase := range testCases {
		algorithms.BubbleSort(testCase.slice)
		if !reflect.DeepEqual(testCase.slice, testCase.expected) {
			t.Errorf("BubbleSort(%v) = %v; expected %v", testCase.slice, testCase.slice, testCase.expected)
		}
	}
}
