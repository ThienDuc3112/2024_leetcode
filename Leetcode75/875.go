package main

import (
	"math"
)

func binarysearch(piles []int, h, left, right int) int {
    if left == right {
        return left
    }
    mid := (left + right) / 2
    time := 0
    for _, bananas := range piles {
        time += int(math.Ceil(float64(bananas) / float64(mid)))
    }
    if time <= h {
        return binarysearch(piles, h, left, mid)
    } else {
        return binarysearch(piles, h, mid+1, right)
    }
}

func minEatingSpeed(piles []int, h int) int {
    // No point in sorting this since binary search the value of the largest
    // element, not the array itself
	// sort.Slice(piles, func(i, j int) bool {return piles[i] < piles[j]})

    return binarysearch(piles, h, 1, 1_000_000_000)
}
