package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if i < size-1 {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

func ConnectRecursive(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		if root.Right != nil {
			root.Left.Next = root.Right
		} else {
			root.Left.Next = findNext(root)
		}
	}

	if root.Right != nil {
		root.Right.Next = findNext(root)
	}

	ConnectRecursive(root.Right)
	ConnectRecursive(root.Left)

	return root
}

func findNext(root *Node) *Node {
	for root != nil {
		if root.Next != nil {
			if root.Next.Left != nil {
				return root.Next.Left
			} else if root.Next.Right != nil {
				return root.Next.Right
			}
		}
		root = root.Next
	}
	return nil
}
