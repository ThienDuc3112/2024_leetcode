package main

func findMinHeightTrees(n int, edges [][]int) []int {
    if n == 1 {
        return []int{0}
    }
    adj := make([]map[int]bool, n)
    for i := range adj {
        adj[i] = make(map[int]bool)
    }
    
    for _, edge := range edges {
        adj[edge[0]][edge[1]] = true
        adj[edge[1]][edge[0]] = true
    }

    leaves := make([]int, 0)
    for i := range adj {
        if len(adj[i]) == 1 {
            leaves = append(leaves, i)
        }
    }

    for n > 2 {
        n -= len(leaves)
        newLeaves := make([]int, 0)
        for _, leaf := range leaves {
            parent := 0
            for k := range adj[leaf] {
                parent = k
            }
            delete(adj[parent], leaf)
            if len(adj[parent]) == 1 {
                newLeaves = append(newLeaves, parent)
            }
        }
        leaves = newLeaves
    }
    return leaves
}

// BFS every vertex method -> TLE
func BFSFindMinHeightTrees(n int, edges [][]int) []int {
	adjList := make([][]int, n)

	for i := range adjList {
		adjList[i] = make([]int, 0)
	}

	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	maxHeight := n
	res := make([]int, 0)

	var bfs func(int)

	bfs = func(i int) {
		visited := make(map[int]bool)
		q := make([]int, 0)
		curHeight := 0
		visited[i] = true
		q = append(q, i)

		for len(q) > 0 {
			length := len(q)

			for l := 0; l < length; l++ {
				cur := q[0]
				q = q[1:]

				if curHeight > maxHeight {
					return
				}

				for _, child := range adjList[cur] {
					_, ok := visited[child]
					if !ok {
						visited[child] = true
						q = append(q, child)
					}
				}
			}

			curHeight++
		}

		if curHeight == maxHeight {
			res = append(res, i)
		} else if curHeight < maxHeight {
			res = make([]int, 0)
			res = append(res, i)
			maxHeight = curHeight
		}
	}

	for root := 0; root < n; root++ {
		bfs(root)
	}

	return res
}
