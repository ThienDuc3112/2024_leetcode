namespace ThreeSum;
public class Solution
{
    public IList<IList<int>> ThreeSum(int[] nums)
    {
        Array.Sort(nums);
        List<IList<int>> res = new();
        for (int start = 0; start < nums.Length - 2 && nums[start] <= 0; start++)
        {
            int head = start + 1, tail = nums.Length - 1;
            int curNum = nums[start];
            while (head < tail)
            {
                int sum = nums[head] + nums[tail];
                if (sum < -curNum)
                {
                    head++;
                }
                else if (sum > -curNum)
                {
                    tail--;
                }
                else
                {
                    int curLeft = nums[head];
                    int curRight = nums[tail];
                    res.Add(new int[] { nums[start], curLeft, curRight });
                    while (head < tail && curLeft == nums[head]) head++;
                    while (head < tail && curRight == nums[tail]) tail--;

                }

            }
            while (start + 1 < nums.Length && nums[start + 1] == curNum)
            {
                start++;
            }
        }
        return res;
    }
}