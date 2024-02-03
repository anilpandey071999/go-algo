package algorithms

type StackStruct struct {
	Head   *Node
	Length int
}

func (st *StackStruct) Push(value int) *Node {
	node := &Node{
		Value: value,
		Next:  st.Head,
	}
	st.Head = node

	st.Length++

	return st.Head
}

func (st *StackStruct) Pop() *Node {
	newHead := st.Head.Next
	st.Head = newHead

	st.Length--

	return st.Head
}

func StackMain() {
	stack := StackStruct{}
	stack.Push(10)
	stack.Push(11)
	Loop(stack.Head)
	stack.Pop()
	stack.Push(12)
	Loop(stack.Head)
}
