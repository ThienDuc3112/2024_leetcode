package main

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	var backtrack func(i, j, curLen int) bool
	backtrack = func(i, j, curLen int) bool {
		if curLen == len(word) {
			return true
		}
		if word[curLen] != board[i][j] {
			return false
		}
		curLen++
		if curLen == len(word) {
			return true
		}
		visited[i][j] = true
		found := false
		if i-1 >= 0 && !visited[i-1][j] {
			found = backtrack(i-1, j, curLen)
		}
		if found {
			return true
		}
		if i+1 < len(board) && !visited[i+1][j] {
			found = backtrack(i+1, j, curLen)
		}
		if found {
			return true
		}
		if j-1 >= 0 && !visited[i][j-1] {
			found = backtrack(i, j-1, curLen)
		}
		if found {
			return true
		}
		if j+1 < len(board[i]) && !visited[i][j+1] {
			found = backtrack(i, j+1, curLen)
		}
        if !found {
            visited[i][j] = false
        }
		return found
	}
	for i := range board {
		for j := range board[i] {
			if backtrack(i, j, 0) {
				return true
			}
		}
	}
	return false
}
