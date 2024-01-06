class Solution {
    public int removeDuplicates(int[] nums) {
        int curVal = nums[0];
        int empty = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == curVal)
                continue;
            nums[empty] = curVal = nums[i];
            empty++;
        }
        return empty;
    }
}