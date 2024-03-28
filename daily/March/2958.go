package main

func maxSubarrayLength(nums []int, k int) int {
    freqMap := make(map[int]int)
    leftIdx := 0
    res := 0
    for rightIdx, right := range nums {
        if _, ok := freqMap[right]; !ok {
            freqMap[right] = 0
        }
        freqMap[right]++
        for freqMap[right] > k {
            freqMap[nums[leftIdx]]--
            leftIdx++
        }
        res = max(res, rightIdx - leftIdx + 1)
    }
    return res
}
