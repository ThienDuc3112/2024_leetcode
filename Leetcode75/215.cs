namespace FindKthLargest;
public class Solution
{
    public int FindKthLargest(int[] nums, int k)
    {
        PriorityQueue<int, int> almostSorted = new();
        foreach (var num in nums)
        {
            almostSorted.Enqueue(num, -num);
        }
        int res = 0;
        while (k > 0)
        {
            res = almostSorted.Dequeue();
            k--;
        }
        return res;
    }
}