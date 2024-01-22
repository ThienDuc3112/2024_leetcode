namespace Rob;
public class Solution
{
    public int Rob(int[] nums)
    {
        if (nums.Length <= 1)
        {
            return nums[0];
        }
        else if (nums.Length == 2)
        {
            return Math.Max(nums[0], nums[1]);
        }
        int[] dp = new int[nums.Length];
        dp[0] = nums[0];
        dp[1] = nums[1];
        dp[2] = nums[2] + nums[0];
        int maxSoFar = Math.Max(dp[0], dp[1]);
        for (int i = 3; i < nums.Length; i++)
        {
            dp[i] = maxSoFar + nums[i];
            maxSoFar = Math.Max(maxSoFar, dp[i - 1]);
        }

        System.Console.WriteLine(string.Join(", ", dp));
        return Math.Max(dp[^1], dp[^2]);
    }
}