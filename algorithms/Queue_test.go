package algorithms_test

import (
	"testing"

	"github.com/anilpandey071999/go-algo/algorithms"
)

func TestQueueOperations(t *testing.T) {
	ll := &algorithms.LinkList{} // Initialize an empty queue

	// Test adding elements to the queue
	valuesToAdd := []interface{}{1, 2, 3, 4}
	for _, value := range valuesToAdd {
		ll.AddNewNodeAtLast(value)
		if ll.Tail.Value != value {
			t.Errorf("Tail.Value got %v, want %v", ll.Tail.Value, value)
		}
	}

	if ll.Length != len(valuesToAdd) {
		t.Errorf("Length after adding got %d, want %d", ll.Length, len(valuesToAdd))
	}

	// Test removing elements from the queue
	expectedLength := len(valuesToAdd)
	for _, expectedValue := range valuesToAdd {
		if ll.Head.Value != expectedValue {
			t.Errorf("Head.Value got %v, want %v", ll.Head.Value, expectedValue)
		}
		ll.RemoveFromQueue()
		expectedLength--
		if ll.Length != expectedLength {
			t.Errorf("Length after removing got %d, want %d", ll.Length, expectedLength)
		}
	}

	// Ensure the queue is empty
	if ll.Head != nil || ll.Tail != nil {
		t.Errorf("Expected empty queue, got Head: %v, Tail: %v", ll.Head, ll.Tail)
	}

	// Attempt to remove from an empty queue should not change length
	ll.RemoveFromQueue()
	if ll.Length != 0 {
		t.Errorf("Length after attempting to remove from empty got %d, want 0", ll.Length)
	}
}
