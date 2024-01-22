namespace CoinChange;

// Have to look at the topic
public class Solution
{
    // Dictionary<int, int> memo;
    // int rCoinChange(int[] coins, int amount)
    // {
    //     if (memo.ContainsKey(amount)) return memo[amount];
    //     if (amount < 0) return -1;
    //     if (amount == 0) return 0;
    //     int res = -1;
    //     foreach (var coin in coins)
    //     {
    //         int curCount = rCoinChange(coins, amount - coin);
    //         // System.Console.WriteLine(curCount);
    //         if (curCount != -1)
    //         {
    //             res = res == -1 ? curCount + 1 : Math.Min(curCount + 1, res);
    //         }
    //     }
    //     memo.Add(amount, res);
    //     return res;
    // }
    // public int CoinChange(int[] coins, int amount)
    // {
    //     memo = new();
    //     return rCoinChange(coins, amount);
    // }

    // Second try with array and loop
    public int CoinChange(int[] coins, int amount)
    {
        int[] dp = new int[amount + 1];
        for (int i = 1; i < amount + 1; i++)
        {
            dp[i] = -1;
            foreach (int coin in coins)
            {
                if (i - coin >= 0 && dp[i - coin] != -1)
                {
                    dp[i] = dp[i] == -1 ? dp[i - coin] + 1 : Math.Min(dp[i], dp[i - coin] + 1);
                }
            }
        }
        return dp.Last();
    }
}