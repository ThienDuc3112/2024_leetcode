class Solution {
    // Read solution
    public int rangeBitwiseAnd(int left, int right) {
        while (right > left)
            right &= (right - 1);
        return right;
    }
}