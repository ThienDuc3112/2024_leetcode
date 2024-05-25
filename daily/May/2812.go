package main

import (
	"fmt"
	"math"
)

// Incomplete
func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return 0
    }
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt
		}
	}

    // BFS
	q := make([][]int, 0)
	for i, row := range grid {
		for j, num := range row {
			if num == 1 {
				q = append(q, []int{i, j, 0})
			}
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]

		if cur[2] >= dp[cur[0]][cur[1]] {
			continue
		}
		dp[cur[0]][cur[1]] = cur[2]

		if cur[0]+1 < len(grid) {
			q = append(q, []int{cur[0] + 1, cur[1], cur[2] + 1})
		}
		if cur[0]-1 >= 0 {
			q = append(q, []int{cur[0] - 1, cur[1], cur[2] + 1})
		}
		if cur[1]+1 < len(grid[cur[0]]) {
			q = append(q, []int{cur[0], cur[1] + 1, cur[2] + 1})
		}
		if cur[1]-1 >= 0 {
			q = append(q, []int{cur[0], cur[1] - 1, cur[2] + 1})
		}
	}



	fmt.Print(dp)
	return 0
}
