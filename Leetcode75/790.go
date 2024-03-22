package main

func numTilings(n int) int {
	dp := make([]int, max(n+1, 4))
	dp[1] = 1
	dp[2] = 2
	dp[3] = 5
	for i := 4; i <= n; i++ {
        // I know to use 1d dp, i just don't know 
        // this formula
		dp[i] = dp[i-3] + dp[i-1]*2
	}
	return dp[n]
}
