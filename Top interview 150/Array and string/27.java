class Solution {
    public int removeElement(int[] nums, int val) {
        int availableSpace = 0;
        for (int curSpace = 0; curSpace < nums.length; curSpace++) {
            if (nums[curSpace] == val)
                continue;
            nums[availableSpace] = nums[curSpace];
            availableSpace++;
        }
        return availableSpace;
    }
}