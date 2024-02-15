class Solution {
    public int[] rearrangeArray(int[] nums) {
        int[] res = new int[nums.length];
        int pos = 0;
        int neg = 1;
        for (int i : nums) {
            if (i > 0) {
                res[pos] = i;
                pos += 2;
            } else {
                res[neg] = i;
                neg += 2;
            }
        }
        return res;
    }
}