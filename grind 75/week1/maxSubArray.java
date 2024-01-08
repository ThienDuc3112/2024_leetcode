class Solution {
    // Have some help with solution
    // Subproblem: subMaxSubArray(int[] nums, int i)
    // Return max array that INCLUDES nums[i]
    public int maxSubArray(int[] nums) {
        int[] dp = new int[nums.length];
        int res = dp[0] = nums[0];
        for (int i = 1; i < dp.length; i++) {
            dp[i] = Math.max(dp[i - 1] + nums[i], nums[i]);
            res = Math.max(res, dp[i]);
        }
        return res;
    }
}