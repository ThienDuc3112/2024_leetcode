package main

func dfs(grid [][]int, i, j int) int {
	var valid func(int, int) bool
	valid = func(i1, i2 int) bool {
		return i1 >= 0 && i1 < len(grid) && i2 >= 0 && i2 < len(grid[i1]) && grid[i1][i2] != 0
	}
	cur := grid[i][j]
	grid[i][j] = 0

	res := 0

	if valid(i-1, j) && grid[i-1][j] != 0 {
		res = max(res, dfs(grid, i-1, j))
	}
	if valid(i+1, j) && grid[i+1][j] != 0 {
		res = max(res, dfs(grid, i+1, j))
	}
	if valid(i, j-1) && grid[i][j-1] != 0 {
		res = max(res, dfs(grid, i, j-1))
	}
	if valid(i, j+1) && grid[i][j+1] != 0 {
		res = max(res, dfs(grid, i, j+1))
	}

	grid[i][j] = cur
    return res + cur
}

func getMaximumGold(grid [][]int) int {
    res := 0
    for i, row := range grid {
        for j, num := range row {
            if num != 0 {
                res = max(res, dfs(grid, i, j))
            }
        }
    }
    return res
}
