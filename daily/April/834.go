package main

// Need to check answer

// Double dfs

// First one calculate weight of subtree
// and the distance between root and all the nodes in 
// the subtree

// Second is to calculate the distance with the distance from
// the first dfs, where everytime we shift to a child node,
// the distance of every other nodes +1, while the distance of 
// the nodes inside the child subtree -1

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph := make(map[int][]int)
	count := make([]int, n)
	res := make([]int, n)

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs1 func(cur, parent int)
	dfs1 = func(cur, parent int) {
		count[cur] = 1
		for _, child := range graph[cur] {
			if child != parent {
				dfs1(child, cur)
				count[cur] += count[child]
				res[cur] += res[child] + count[child]
			}
		}
	}

	var dfs2 func(cur, parent int)
	dfs2 = func(cur, parent int) {
		for _, child := range graph[cur] {
			if child != parent {
				res[child] = res[cur] + n - 2*count[child]
				dfs2(child, cur)
			}
		}
	}

	dfs1(0, -1)
	dfs2(0, -1)

	return res
}

// Constant bfs, TLE
func bfsSumOfDistancesInTree(n int, edges [][]int) []int {
	adjSet := make([]map[int]bool, n)
    for i := range adjSet {
        adjSet[i] = make(map[int]bool)
    }

	for _, edge := range edges {
		adjSet[edge[0]][edge[1]] = true
		adjSet[edge[1]][edge[0]] = true
	}

	var bfs func(int) int
	bfs = func(i int) int {
		visitedNode := make(map[int]bool)
		visitedNode[i] = true
		queue := make([]int, 0)
		queue = append(queue, i)
		distance := 0
        res := 0

		for len(queue) > 0 {
			length := len(queue)

			for j := 0; j < length; j++ {
				cur := queue[0]
				queue = queue[1:]
                res += distance
                for neighbour  := range adjSet[cur] {
                    _, visited := visitedNode[neighbour]
                    if !visited {
                        queue = append(queue, neighbour)
                        visitedNode[neighbour] = true
                    }
                }
			}

            distance++
		}
        return res
	}

    res := make([]int, n)
	for i := 0; i < n; i++ {
        res[i] = bfs(i)
	}
    return res
}
