package main

// Better dfs way
func numIslands(grid [][]byte) int {
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
    res := 0
	for i := range grid {
		for j := range grid[i] {
            if grid[i][j] == '1' {
                res++
                dfs(i, j)
            }
		}
	}
    return res
}

// The UFDS way
// func numIslands(grid [][]byte) int {
// 	dir := []int{0, 1, 0, -1, 0}
// 	parent := make(map[int]int)
// 	weight := make(map[int]int)
// 	for i, row := range grid {
// 		for j, land := range row {
// 			if land != '1' {
// 				continue
// 			}
// 			id := i*len(grid[i]) + j
// 			parent[id] = id
// 			weight[id] = 1
// 			for a := 1; a < len(dir); a++ {
// 				newI := i + dir[a-1]
// 				newJ := j + dir[a]
//                 if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[i]) || grid[newI][newJ] == '0' {
// 					continue
// 				}
//                 if _, ok := parent[newI * len(grid[newI]) + newJ]; !ok {
//                     continue
//                 }
// 				// Add to the ufds
// 				root1 := newI*len(grid[newI]) + newJ
// 				root2 := id
// 				for root1 != parent[root1] {
// 					root1 = parent[root1]
// 				}
// 				for root2 != parent[root2] {
// 					root2 = parent[root2]
// 				}
// 				if root1 == root2 {
// 					continue
// 				}
// 				if weight[root1] >= weight[root2] {
// 					parent[root2] = root1
// 					weight[root1] += weight[root2]
// 				} else {
// 					parent[root1] = root2
// 					weight[root2] += weight[root1]
// 				}
// 			}
// 		}
// 	}
// 	res := 0
// 	for k, v := range parent {
// 		if k == v {
// 			res++
// 		}
// 	}
// 	return res
// }
