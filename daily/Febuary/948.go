package main

import (
	"sort"
)

func bagOfTokensScore(tokens []int, power int) int {
	sort.Ints(tokens)
	left := 0
	right := len(tokens) - 1
	score := 0
	curScore := 0
	for left <= right {
		if power < tokens[left] && curScore == 0 {
			break
		}
		for left <= right && power >= tokens[left] {
			power -= tokens[left]
			curScore++
			left++
		}
		score = max(score, curScore)
		for left <= right && curScore > 0 && power < tokens[left] {
			power += tokens[right]
			curScore--
			right--
		}
	}
	return score
}
