package algorithms

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (ll *LinkList) IncreaseLength() {
	ll.Length++
}
func (ll *LinkList) DecreaseLength() {
	ll.Length--
}

func Loop(head *Node) {
	for head != nil {
		fmt.Println(head.Value)
		if head.Next == nil {
			break
		}
		head = head.Next
	}
}

func (ll *LinkList) RemoveFromQueue() *LinkList {
	if ll.Length == 0 {
		fmt.Println("Nothing to remove")
		return ll
	}
	if ll.Length == 1 {
		ll.Tail = nil
	}
	fmt.Println("Old Head:- ", ll.Head)
	newHead := ll.Head.Next
	ll.Head.Next = nil
	ll.Head = newHead
	fmt.Println("New Head:- ", ll.Head)



	ll.DecreaseLength()
	return ll
}

func (ll *LinkList) AddNewNodeAtLast(value interface{}) *LinkList {
	current := &Node{value, nil}
	if ll.Head == nil && ll.Tail == nil {
		ll.Head = current
		ll.Tail = current
	}

	ll.Tail.Next = current
	ll.Tail = current

	ll.IncreaseLength()

	return ll
}

func Queue() {
	list := &LinkList{}
	list = list.AddNewNodeAtLast(1)
	list = list.AddNewNodeAtLast(2)
	list = list.AddNewNodeAtLast(3)
	Loop(list.Head)
	fmt.Println("Length After Adding:- ",list.Length, "\n Start Deletion")
	list = list.RemoveFromQueue()
	list = list.RemoveFromQueue()
	list = list.RemoveFromQueue()
	list = list.RemoveFromQueue()
	Loop(list.Head)
	fmt.Println("Length After deleting:- ",list.Length)
}
