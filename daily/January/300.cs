namespace LengthOfLIS;
public class Solution
{
    public int LengthOfLIS(int[] nums)
    {
        int[] maxTillNow = new int[nums.Length];
        maxTillNow[0] = 1;
        for (int i = 1; i < nums.Length; i++)
        {
            maxTillNow[i] = 1;
            for (int j = i - 1; j >= 0; j--)
            {
                if (nums[i] > nums[j]) maxTillNow[i] = Math.Max(maxTillNow[i], maxTillNow[j] + 1);
            }
        }
        return maxTillNow.Max();
    }
}