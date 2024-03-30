package main

// Check answer
// Forgot that sliding window also include the 2 atMost function
// Sliding window: 
// 1, Left pointer and right pointer
// 2, AtMost function for "Exact number of k" type questions
func subarraysWithKDistinct(nums []int, k int) int {
    return atMost(nums, k) - atMost(nums, k - 1)
}

func atMost(nums []int, k int) int {
    leftIdx := 0
    res := 0
    freq := make(map[int]int)
    for rightIdx, right := range nums {
        if _, ok := freq[right]; !ok {
            freq[right] = 0
        }
        freq[right]++
        for len(freq) > k {
            freq[nums[leftIdx]]--
            if freq[nums[leftIdx]] == 0 {
                delete(freq, nums[leftIdx])
            }
            leftIdx++
        }
        res += rightIdx - leftIdx + 1
    }
    return res
}
