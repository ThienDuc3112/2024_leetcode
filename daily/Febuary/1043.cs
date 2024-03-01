namespace maxSumAfterPartitioning;
// From solution tab
class Solution
{
    static public int maxSumAfterPartitioning(int[] A, int K)
    {
        int N = A.Length;
        var dp = new int[N + 1];
        for (int i = 1; i <= N; ++i)
        {
            int curMax = 0, best = 0;
            for (int k = 1; k <= K && i - k >= 0; ++k)
            {
                curMax = Math.Max(curMax, A[i - k]);
                best = Math.Max(best, dp[i - k] + curMax * k);
            }
            dp[i] = best;
        }
        
        return dp[N];
    }
}
