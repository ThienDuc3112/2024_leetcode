package main

func tribonacci(n int) int {
    if n <= 0 {
        return 0
    }
    if n < 3 {
        return 1
    }
    dp := []int {0,1,1}
    for i := 3; i <= n; i++ {
        dp = append(dp, dp[i-1] + dp[i-2] + dp[i-3])
    }
    return dp[n]
}
