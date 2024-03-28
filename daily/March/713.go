package main

func numSubarrayProductLessThanK(nums []int, k int) int { 
    if k <= 1 {return 0}
    res := 0
    leftIdx := 0
	product := 1
    for rightIdx, right := range nums {
        product *= right
        for product >= k {
            product /= nums[leftIdx]
            leftIdx++
        }
        res += rightIdx - leftIdx + 1
	}
	return res
}
