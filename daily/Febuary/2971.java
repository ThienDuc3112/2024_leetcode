import java.util.Arrays;

class Solution {
    // Java don't support heapify, but
    // Using heapify can create a O(n) solution
    public long largestPerimeter(int[] nums) {
        Arrays.sort(nums);
        long sum = 0;
        for (int i : nums) {
            sum += i;
        }
        for (int i = nums.length - 1; i >= 0; i--) {
            sum -= nums[i];
            if (nums[i] < sum && i > 1)
                return sum + nums[i];
        }
        return -1;
    }
}