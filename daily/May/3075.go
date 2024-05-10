package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
    var res int64 = 0
    for i := 0; i < k; i++ {
        happyValue := happiness[len(happiness) - 1 - i] - i
        if happyValue <= 0 {break}
        res += int64(happyValue)

    }
    return res
}
