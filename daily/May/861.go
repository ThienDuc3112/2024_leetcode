package main

func flipRow(grid [][]int, rowIdx int) {
	for i := range grid[rowIdx] {
		if grid[rowIdx][i] == 1 {
			grid[rowIdx][i] = 0
		} else {
			grid[rowIdx][i] = 1
		}
	}
}

func flipCol(grid [][]int, colIdx int) {
	for i := 0; i < len(grid); i++ {
		if grid[i][colIdx] == 1 {
			grid[i][colIdx] = 0
		} else {
			grid[i][colIdx] = 1
		}
	}
}

func matrixScore(grid [][]int) int {
    for i := 0; i < len(grid); i++ {
        if grid[i][0] == 0 {
            flipRow(grid, i)
        }
    }

    for i := 1; i < len(grid[0]); i++ {
        count := 0
        for j := 0; j < len(grid); j++ {
            if grid[j][i] == 1 { count++ }
        }
        if count < len(grid) - count {
            flipCol(grid, i)
        }
    }

    res := 0
    for i := range(grid) {
        number := 0
        for _, bit := range(grid[i]) {
            number *= 2
            number += bit
        }
        res += number
    }
    return res
}
