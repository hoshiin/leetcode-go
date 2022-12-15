package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func CloneGraph(node *Node) *Node {
	return vis(node, make(map[int]*Node))
}

func vis(node *Node, mp map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := mp[node.Val]; !ok {
		nn := &Node{Val: node.Val}
		mp[node.Val] = nn
		for _, nx := range node.Neighbors {
			nn.Neighbors = append(nn.Neighbors, vis(nx, mp))
		}
	}
	return mp[node.Val]
}

func CloneGraphBFS(node *Node) *Node {
	if node == nil {
		return node
	}
	visited := map[*Node]*Node{node: {Val: node.Val, Neighbors: []*Node{}}}
	queue := []*Node{node}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		for _, neighbor := range n.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				visited[neighbor] = &Node{Val: neighbor.Val, Neighbors: []*Node{}}
				queue = append(queue, neighbor)
			}
			visited[n].Neighbors = append(visited[n].Neighbors, visited[neighbor])
		}
	}
	return visited[node]
}
