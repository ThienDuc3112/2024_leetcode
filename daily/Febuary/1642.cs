namespace FurthestBuilding;
public class Solution
{
    // Initial thought were correct
    // But somehow the implementation is wrong??
    // Need to look at solution for implementation of 
    // Priority queue
    public int FurthestBuilding(int[] heights, int bricks, int ladders)
    {
        PriorityQueue<int, int> pq = new();
        for (int i = 0; i < heights.Length - 1; i++)
        {
            int heightDiff = heights[i + 1] - heights[i];
            if (heightDiff <= 0) continue;
            pq.Enqueue(heightDiff, -heightDiff);
            if (bricks < heightDiff)
            {
                if (ladders == 0) return i;
                ladders--;
                bricks += pq.Dequeue();
            }
            bricks -= heightDiff;
        }
        return heights.Length - 1;
    }
}