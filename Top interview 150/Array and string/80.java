class Solution {
    public int removeDuplicates(int[] nums) {
        int curVal = nums[0];
        int empty = 1;
        int apperance = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == curVal && apperance > 1)
                continue;
            if (nums[i] == curVal)
                apperance++;
            if (nums[i] != curVal)
                apperance = 1;
            nums[empty] = curVal = nums[i];
            empty++;
        }
        return empty;
    }
}
