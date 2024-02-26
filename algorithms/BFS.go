package algorithms

import "fmt"

func BFS(needle int) {
	rootNode := BinaryTree()

	var queue, visited []*BinaryTreeNode
	queue = append(queue, rootNode)
	res := false
	for len(queue) != 0 {
		fmt.Println("==========================")
		fmt.Println("DeQueue: ", queue[0].Value, len(queue))

		if queue[0].Value == needle {
			fmt.Println("Found Needle!")
			res = true
			break
		}

		visited = append(visited, queue[0])
		fmt.Println(visited)

		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
			visited = append(visited, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
			visited = append(visited, queue[0].Right)
		}

		fmt.Println("queue1", queue)
		//queue = append(queue[:1], queue[:len(queue))
		fmt.Println("queue1:= ", queue)
		fmt.Println("queue1", len(queue), queue[1:], queue[:1])
		queue = queue[1:]
	}
	if !res {
		fmt.Println("Not Found Needle!")
	}

}

// 10,20,5,2,3,4,13,23,30
