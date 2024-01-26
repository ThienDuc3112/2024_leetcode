namespace ProductExceptSelf;
// Check solution
public class Solution
{
    public int[] ProductExceptSelf(int[] nums)
    {
        int[] res = new int[nums.Length];
        int left = 1;
        for (int i = 0; i < res.Length; i++)
        {
            res[i] = left;
            left *= nums[i];
        }
        int right = 1;
        for (int i = res.Length - 1; i >= 0; i--)
        {
            res[i] *= right;
            right *= nums[i];
        }
        return res;
    }
}