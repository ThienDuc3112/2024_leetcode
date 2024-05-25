package main

import "math"

func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	var res int64 = 0
	even := true
	minDiff := math.MaxInt
	for _, a := range nums {
		b := a ^ k
		if b >= a {
			res += int64(b)
			even = !even
			minDiff = min(minDiff, b-a)
		} else {
			res += int64(a)
			minDiff = min(minDiff, a-b)
		}
	}
	if even {
		return res
	}
	return res - int64(minDiff)
}

// func maximumValueSum(nums []int, k int, edges [][]int) int64 {
// 	var res int64 = 0
// 	even := 0
// 	minDiff := math.MaxInt
// 	for _, a := range nums {
// 		b := a ^ k
//         res += int64(max(a,b))
//         if a < b {even ^= 1}
//         if b >= a {
//             minDiff = min(minDiff, (b-a))
//         } else {
//             minDiff = min(minDiff, (a-b))
//         }
// 	}
//     return res - int64(minDiff) * int64(even)
// }
