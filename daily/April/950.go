package main

import "sort"

// The solution (just simulate the process but played backward)
func deckRevealedIncreasing(deck []int) []int {
	sort.Slice(deck, func(i, j int) bool { return deck[i] < deck[j] })
    queue := make([]int, 0)
    for i := len(deck) - 1; i >= 0; i-- {
        if len(queue) > 0 {
            queue = append(queue, queue[0])
            queue = queue[1:]
        }
        queue = append(queue, deck[i])
    }
    for i, j := 0, len(queue)-1; i < j; i, j = i+1, j-1 {
        queue[i], queue[j] = queue[j], queue[i]
    }
    return queue
}

// The first method
// func deckRevealedIncreasing(deck []int) []int {
// 	sort.Slice(deck, func(i, j int) bool { return deck[i] < deck[j] })
// 	res := make([]int, len(deck))
// 	ptr := 0
// 	freeSpace := 0
// 	add := true
//
// 	for ptr < len(deck) {
// 		for res[freeSpace] != 0 {
// 			freeSpace++
// 			freeSpace %= len(deck)
// 		}
// 		if !add {
// 			freeSpace++
// 			freeSpace %= len(deck)
// 			add = true
// 			continue
// 		}
// 		res[freeSpace] = deck[ptr]
// 		ptr++
// 		add = false
// 		freeSpace++
// 		freeSpace %= len(deck)
// 	}
// 	return res
// }
