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

func (ll *LinkList) GetLength() int {
	return ll.Length
}

func (ll *LinkList) increaseLength() {
	ll.Length++
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



func (ll *LinkList) AddNewNodeAtLast(value interface{}) *LinkList {
	current := &Node{value, nil}
	if ll.Head == nil && ll.Tail == nil {
		ll.Head = current
		ll.Tail = current
	}

	ll.Tail.Next = current
	ll.Tail = current

	ll.increaseLength()

	return ll
}

func Queue() {
	list := &LinkList{}
	list = list.AddNewNodeAtLast(1)
	list = list.AddNewNodeAtLast(2)
	list = list.AddNewNodeAtLast(3)
	Loop(list.Head)
	fmt.Println(list.GetLength())
}
