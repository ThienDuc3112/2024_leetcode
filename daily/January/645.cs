namespace FindErrorNums;
public class Solution
{
    public int[] FindErrorNums(int[] nums)
    {
        HashSet<int> set = new();
        int[] res = new int[2];
        foreach (var num in nums)
        {
            if (set.Contains(num)) res[0] = num;
            else set.Add(num);

        }
        for (int i = 1; i <= nums.Length; i++)
        {
            if (!set.Contains(i)) res[1] = i;
        }
        return res;
    }
}