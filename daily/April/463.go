package main

func islandPerimeter(grid [][]int) int {
	dir := []int{0, 1, 0, -1, 0}
	res := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				for a := 1; a < len(dir); a++ {
                    newI := i + dir[a-1]
                    newJ := j + dir[a]
                    if newI < 0 || newI >= len(grid) || newJ < 0 || newJ >= len(grid[newI]) || grid[newI][newJ] == 0 {
                        res++
                    }
				}
			}
		}
	}
	return res
}
