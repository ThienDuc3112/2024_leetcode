package main

// Check answer
// Today was a bit burn out alr

func numSubarraysWithSum(nums []int, goal int) int {
    dp := make([]int, len(nums) + 1)
    dp[0] = 1
    sum := 0
    res := 0
    for _, val := range nums {
        sum += val
        if sum >= goal {
            res += dp[sum - goal]
        }
        dp[sum]++
    }
    return res
}
