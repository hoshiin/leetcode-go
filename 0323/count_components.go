package leetcode

func CountComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	visited := make([]bool, n)
	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true
		neighbors := graph[node]
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			dfs(i)
		}
	}
	return count
}
