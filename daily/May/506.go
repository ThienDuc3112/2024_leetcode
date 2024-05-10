package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(scores []int) []string {
	type item struct {
		score int
		index int
	}
	newScores := make([]item, len(scores))
	for i, score := range scores {
		newScores[i] = item{score: score, index: i}
	}

	sort.Slice(newScores, func(i, j int) bool {return newScores[i].score < newScores[j].score})

    res := make([]string, len(scores))

    for i, score := range newScores {
        rank := len(newScores) - i
        if rank == 1 {
            res[score.index] = "Gold Medal"
        } else if rank == 2 {
            res[score.index] = "Silver Medal"
        } else if rank == 3 {
            res[score.index] = "Bronze Medal"
        } else {
            s := strconv.Itoa(rank)
            res[score.index] = s
        }
    }

    return res
}
