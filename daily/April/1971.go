package main

func validPath(n int, edges [][]int, source int, destination int) bool {
    if source == destination {
        return true
    }
    visited := make([]bool, n)
    queue := make([]int,0)
    queue = append(queue, source)
    visited[source] = true
    for len(queue) > 0 {
        cur := queue[0]
        if cur == destination {
            return true
        }
        queue = queue[1:]
        for _, edge := range edges {
            if edge[0] == cur && !visited[edge[1]] {
                visited[edge[1]] = true
                queue = append(queue, edge[1])
            }
            if edge[1] == cur && !visited[edge[0]] {
                visited[edge[0]] = true
                queue = append(queue, edge[0])
            }
        }
    }
    return false
}
