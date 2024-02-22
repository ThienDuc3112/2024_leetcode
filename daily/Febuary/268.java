class Solution {
    public int missingNumber(int[] nums) {
        int sum = 0;
        int trueSum = 0;
        for (int i = 0; i < nums.length; i++) {
            sum += i;
            trueSum += nums[i];
        }
        sum += nums.length;
        return sum - trueSum;
    }
}