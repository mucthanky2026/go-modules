package search

type Graph struct {
	AjdList map[int][]int
}

func (g *Graph) DFS(node int, visited map[int]bool) {
	if visited[node] {
		return
	}
	visited[node] = true

	for _, neighbor := range g.AjdList[node] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func (g *Graph) BFS(start int) []int {
	visited := make(map[int]bool)
	queue := []int{start}
	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			visited[node] = true
			result = append(result, node)

			for _, neighbor := range g.AjdList[node] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return result
}

func dijkstra(graph map[int]map[int]int, start int) map[int]int {
	distances := make(map[int]int)
	for node := range graph {
		distances[node] = 1<<31 - 1 // Initialize with a large value (infinity)
	}
	distances[start] = 0

	visited := make(map[int]bool)

	for len(visited) < len(graph) {
		// Find the unvisited node with the smallest distance
		minNode := -1
		minDistance := 1<<31 - 1
		for node, distance := range distances {
			if !visited[node] && distance < minDistance {
				minDistance = distance
				minNode = node
			}
		}

		if minNode == -1 {
			break // All reachable nodes have been visited
		}

		visited[minNode] = true

		for neighbor, weight := range graph[minNode] {
			if !visited[neighbor] {
				newDistance := distances[minNode] + weight
				if newDistance < distances[neighbor] {
					distances[neighbor] = newDistance
				}
			}
		}
	}

	return distances
}

func topologicalSort(graph map[int][]int) ([]int, bool) {
	visited := make(map[int]bool)
	tempMarked := make(map[int]bool)
	result := []int{}
	hasCycle := false

	var visit func(node int)
	visit = func(node int) {
		if tempMarked[node] {
			hasCycle = true
			return
		}
		if !visited[node] {
			tempMarked[node] = true
			for _, neighbor := range graph[node] {
				visit(neighbor)
			}
			tempMarked[node] = false
			visited[node] = true
			result = append(result, node)
		}
	}

	for node := range graph {
		if !visited[node] {
			visit(node)
		}
	}

	if hasCycle {
		return nil, false // Cycle detected, topological sort not possible
	}

	// Reverse the result to get the correct order
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result, true
}	