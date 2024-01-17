namespace TotalCost;
public class Solution
{
    // Technically the same as others solutions
    // Just less efficient
    public long TotalCost(int[] costs, int k, int candidates)
    {
        long res = 0;
        if (costs.Length <= candidates * 2)
        {
            Array.Sort(costs);
            for (int i = 0; i < k; i++)
            {
                res += costs[i];
            }
            return res;
        }
        PriorityQueue<int, int> leftQueue = new();
        PriorityQueue<int, int> rightQueue = new();
        int left = candidates;
        int right = costs.Length - 1 - candidates;
        for (int i = 0; i < left; i++)
        {
            leftQueue.Enqueue(costs[i], costs[i]);
            rightQueue.Enqueue(costs[costs.Length - 1 - i], costs[costs.Length - 1 - i]);
        }

        for (int i = 0; i < k; i++)
        {
            int leftCost = leftQueue.Count > 0 ? leftQueue.Peek() : int.MaxValue;
            int rightCost = rightQueue.Count > 0 ? rightQueue.Peek() : int.MaxValue;
            bool isLeftSmaller = leftCost <= rightCost;
            if (isLeftSmaller)
            {
                res += leftCost;
                leftQueue.Dequeue();
            }
            else
            {
                res += rightCost;
                rightQueue.Dequeue();
            }
            if (left <= right)
            {
                if (isLeftSmaller) leftQueue.Enqueue(costs[left], costs[left++]);
                else rightQueue.Enqueue(costs[right], costs[right--]);
            }
        }
        return res;
    }
}