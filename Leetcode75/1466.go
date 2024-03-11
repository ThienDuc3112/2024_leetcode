package main

// Check answer
func dfs(source int, res *int, visited []bool, adjList map[int][]int) {
	visited[source] = true
	for _, next := range adjList[source] {
		safe := next
		if safe < 0 {
			safe = -safe
		}
		if visited[safe] {
			continue
		}
		if next > 0 {
			(*res)++
		}
		dfs(safe, res, visited, adjList)
	}
	visited[source] = false
}

func minReorder(n int, connections [][]int) int {
	visited := make([]bool, n)
	visited[0] = true
	adjList := make(map[int][]int)
	for _, con := range connections {
		adjList[con[0]] = append(adjList[con[0]], con[1])
		adjList[con[1]] = append(adjList[con[1]], -con[0])
	}
	res := 0
	dfs(0, &res, visited, adjList)
	return res
}

// Inefficient
// func minReorder(n int, connections [][]int) int {
// 	arrive := make([]bool, n)
// 	arrive[0] = true
// 	res := 0
// 	hasChanged := true
// 	for i := 0; i < n && hasChanged; i++ {
// 		hasChanged = false
// 		for _, connection := range connections {
// 			if arrive[connection[1]] && arrive[connection[0]] {
// 				continue
// 			}
// 			if arrive[connection[1]] {
// 				arrive[connection[0]] = true
//                 hasChanged = true
// 			} else if arrive[connection[0]] {
// 				hasChanged = true
// 				res++
// 				arrive[connection[1]] = true
// 			}
// 		}
// 	}
// 	return res
// }
