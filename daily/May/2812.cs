namespace MaximumSafenessFactor;

public class Solution
{
    bool valid(int n, int i, int j)
    {
        return i >= 0 && i < n && j >= 0 && j < n;
    }

    public int MaximumSafenessFactor(IList<IList<int>> grid)
    {
        int n = grid.Count;
        List<List<int>> distances = new(n);
        for (int i = 0; i < n; i++)
        {
            distances.Add(new(n));
            for (int j = 0; j < n; j++)
                distances[i].Add(int.MaxValue);
        }

        Queue<(int i, int j, int dist)> q = new();
        for (int i = 0; i < n; i++)
            for (int j = 0; j < n; j++)
                if (grid[i][j] == 1)
                    q.Enqueue((i, j, 0));

        while (q.Any())
        {
            (int i, int j, int dist) = q.Dequeue();
            if (dist >= distances[i][j]) continue;
            distances[i][j] = dist;
            if (valid(n, i + 1, j)) q.Enqueue((i + 1, j, dist + 1));
            if (valid(n, i - 1, j)) q.Enqueue((i - 1, j, dist + 1));
            if (valid(n, i, j + 1)) q.Enqueue((i, j + 1, dist + 1));
            if (valid(n, i, j - 1)) q.Enqueue((i, j - 1, dist + 1));
        }

        /* foreach (var row in distances) */
        /* { */
        /*     Console.WriteLine("\n"); */
        /*     foreach (int num in row) */
        /*         Console.Write($"{num} "); */
        /* } */

        PriorityQueue<(int, int, int), int> pq = new();

        pq.Enqueue((0, 0, distances[0][0]), -distances[0][0]);

        while (pq.Count > 0)
        {
            (int i, int j, int dist) = pq.Dequeue();
            if (distances[i][j] == -1) continue;
            dist = Math.Min(dist, distances[i][j]);
            if(i == n - 1 && j == n - 1) return dist;
            distances[i][j] = -1;
            grid[i][j] = dist;
            if (valid(n, i + 1, j)) pq.Enqueue((i + 1, j, dist), -dist);
            if (valid(n, i - 1, j)) pq.Enqueue((i - 1, j, dist), -dist);
            if (valid(n, i, j + 1)) pq.Enqueue((i, j + 1, dist), -dist);
            if (valid(n, i, j - 1)) pq.Enqueue((i, j - 1, dist), -dist);
        }
        return grid[n - 1][n - 1];
    }
}
