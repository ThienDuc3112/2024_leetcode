namespace LargestDivisibleSubset;
public class Solution
{
    class LargestNode
    {
        public int length;
        public int? next;
        public LargestNode(int len, int? next)
        {
            length = len;
            this.next = next;
        }
    }
    public IList<int> LargestDivisibleSubset(int[] nums)
    {
        Array.Sort(nums);
        LargestNode[] memo = new LargestNode[nums.Length];
        int maxSoFar = 0;
        int? maxIndex = 0;
        for (int i = nums.Length - 1; i >= 0; i--)
        {
            LargestNode cur = new(0, null);
            for (int j = i + 1; j < nums.Length; j++)
            {
                if (nums[j] % nums[i] == 0 && memo[j].length > cur.length)
                {
                    cur = new(memo[j].length, j);
                }
            }
            cur.length++;
            memo[i] = cur;
            if (maxSoFar < cur.length)
            {
                maxSoFar = cur.length;
                maxIndex = i;
            }
        }
        List<int> res = new();
        while (maxIndex != null)
        {

            res.Add(nums[(int)maxIndex]);
            maxIndex = memo[(int)maxIndex].next;
        }
        return res;
    }
}