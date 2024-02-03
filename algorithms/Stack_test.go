package algorithms_test

import (
	"fmt"
	"testing"

	"github.com/anilpandey071999/go-algo/algorithms"
)

func TestStack(t *testing.T) {
	testCases := []struct {
		StackResult []int // Values to push onto the stack
		Length      int   // Expected length of the stack after pushes
		PopLen      int

	}{
		{[]int{2, 3, 5, 7}, 4, 0},
		{[]int{4, 5, 62}, 3, 0},
	}

	for _, testCase := range testCases {
		stack := &algorithms.StackStruct{} // Assuming StackStruct is defined in the package `algorithms`
		for _, value := range testCase.StackResult {
			fmt.Println("Pushed element is", value)
			stack.Push(value)
		}

		// No need for a Loop function; directly check stack.Length and the top value if needed
		if stack.Length != testCase.Length {
			t.Errorf("After pushing %v, expected stack length %d, got %d", testCase.StackResult, testCase.Length, stack.Length)
		}

		// Example of checking the top value - only if the stack is not empty
		if len(testCase.StackResult) > 0 && stack.Head.Value != testCase.StackResult[len(testCase.StackResult)-1] {
			t.Errorf("Expected top value of %d, got %d", testCase.StackResult[len(testCase.StackResult)-1], stack.Head.Value)
		}

		for _, value := range testCase.StackResult {
			fmt.Println("Poped element is", value)
			stack.Pop()
		}

		// No need for a Loop function; directly check stack.Length and the top value if needed
		if stack.Length != testCase.PopLen {
			t.Errorf("After poping %v, expected stack length %d, got %d", testCase.StackResult, testCase.Length, stack.Length)
		}
	}
}