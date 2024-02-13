namespace CherryPickup;
public class Solution
{
    // Did not work

    // public int CherryPickup(int[][] grid)
    // {
    //     int[][] dp = new int[grid.Length][];
    //     for (int i = grid.Length - 1; i >= 0; i--)
    //     {
    //         dp[i] = new int[grid[i].Length];
    //         for (int j = 0; j < grid[i].Length; j++)
    //         {
    //             if (i == grid.Length - 1)
    //             {
    //                 dp[i][j] = grid[i][j];
    //                 continue;
    //             }
    //             int max = -1;
    //             int maxIdx = 0;
    //             for (int k = -1; k <= 1; k++)
    //             {
    //                 if (j + k < 0 || j + k >= grid[i].Length) continue;
    //                 if (dp[i + 1][j + k] > max)
    //                 {
    //                     max = dp[i + 1][j + k];
    //                     maxIdx = j + k;
    //                 }
    //             }
    //             dp[i][j] = dp[i - 1][maxIdx] + grid[i][j];
    //         }
    //     }
    // }
}