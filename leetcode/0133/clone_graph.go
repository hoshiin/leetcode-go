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
