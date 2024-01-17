// First attempt: TLE, due to a single line
// Check answer => Realized that single line mistake
// Further optimization: just use visited, it's gurantee to not overlap anyway
// Note to self: Only remove if finding unique path, otherwise it's no difference
namespace FindCircleNum;
public class Solution
{
    HashSet<int> visited;
    void DFS(int[][] matrix, int current, HashSet<int> provinceVisited)
    {
        visited.Add(current);
        for (int i = 0; i < matrix[current].Length; i++)
        {
            if (matrix[current][i] == 1 && i != current && !provinceVisited.Contains(i))
            {
                provinceVisited.Add(i);
                DFS(matrix, i, provinceVisited);
                // provinceVisited.Remove(i);
            }
        }

    }
    public int FindCircleNum(int[][] isConnected)
    {
        visited = new();
        int res = 0;
        for (int i = 0; i < isConnected.Length; i++)
        {
            if (!visited.Contains(i))
            {
                res++;
                DFS(isConnected, i, new());
            }
        }
        return res;
    }
}