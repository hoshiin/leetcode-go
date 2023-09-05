package main

func ValidTree(n int, edges [][]int) bool {
	adjacencyList := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		adjacencyList[u] = append(adjacencyList[u], v)
		adjacencyList[v] = append(adjacencyList[v], u)
	}
	visited := make([]bool, n)

	if hasCycle(adjacencyList, visited, 0, -1) {
		return false
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}

func hasCycle(adjacencyList [][]int, visited []bool, node, parent int) bool {
	visited[node] = true
	for _, neighbor := range adjacencyList[node] {
		if visited[neighbor] && neighbor != parent {
			return true
		}
		if !visited[neighbor] && hasCycle(adjacencyList, visited, neighbor, node) {
			return true
		}
	}
	return false
}
