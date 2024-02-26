package algorithms

import "fmt"

type BinaryTreeNode struct {
	Value    int
	Left     *BinaryTreeNode
	Right    *BinaryTreeNode
	Parent   *BinaryTreeNode
	LeafNode bool
}

func GetLeafNode(node *BinaryTreeNode, childNodes []*BinaryTreeNode) []*BinaryTreeNode {
	// childNodes := []*BinaryTreeNode{}
	println(node, "Node Value: ", node.Value, "Left Node", node.Left, "Right node: ", node.Right, "Parent Node: ", node.Parent)
	if node.Left == nil && node.Right == nil {
		childNodes = append(childNodes, node)
		return childNodes
	}
	if node.Left != nil {
		childNodes = GetLeafNode(node.Left, childNodes)
	}
	if node.Right != nil {
		childNodes = GetLeafNode(node.Right, childNodes)
	}
	return childNodes
}

func GetNodeByValue(node *BinaryTreeNode, value int) *BinaryTreeNode {
	if node.Value > value {
		if node.Left != nil {
			return GetNodeByValue(node.Left, value)
		}
		return node
	} else {
		if node.Right != nil {
			return GetNodeByValue(node.Right, value)
		}
		return node
	}
}

func AddChild(node *BinaryTreeNode, value int) {
	/**
	 * checking the level of tree
	 * if root is not set then set the root
	 * else if level is 0 then add to left
	 * else add to right
	 */
	node = GetNodeByValue(node, value)
	if node.Value > value {
		node.LeafNode = false
		node.Left = &BinaryTreeNode{value, nil, nil, node, true}
	} else {
		node.LeafNode = false
		node.Right = &BinaryTreeNode{value, nil, nil, node, true}
	}
	println("Successfull Add : ", value)
}

func BinaryTree() *BinaryTreeNode {
	/**
	 * TODO
	 * 		? AddChild Node function
	 * 		? Traverse Tree function
	 * 		? RemoveChild Node function
	 * 		? Search Node function
	 */
	FirstNode := &BinaryTreeNode{10, nil, nil, nil, true}
	AddChild(FirstNode, 20)
	AddChild(FirstNode, 5)
	AddChild(FirstNode, 2)
	AddChild(FirstNode, 3)
	AddChild(FirstNode, 4)
	AddChild(FirstNode, 13)
	AddChild(FirstNode, 23)
	AddChild(FirstNode, 30)
	data := GetLeafNode(FirstNode, []*BinaryTreeNode{})
	fmt.Println(data)
	return FirstNode
}
