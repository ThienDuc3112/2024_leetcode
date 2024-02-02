import java.util.ArrayList;
import java.util.Arrays;

class Solution {
    public int[][] divideArray(int[] nums, int k) {
        Arrays.sort(nums);
        ArrayList<int[]> res = new ArrayList<>();
        for (int i = 0; i < nums.length - 2; i += 3) {
            if (nums[i + 2] - nums[i] <= k) {
                res.add(new int[] { nums[i], nums[i + 1], nums[i + 2] });
            } else {
                return new int[0][];
            }
        }
        return res.toArray(new int[res.size()][]);
    }
}