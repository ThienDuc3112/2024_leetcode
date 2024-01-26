namespace FindPaths;
public class Solution
{
    const int MOD = 1_000_000_007;
    Dictionary<(int, int, int), int> memo;
    int dfs(int m, int n, int maxMove, int startRow, int startColumn)
    {
        if (memo.ContainsKey((startRow, startColumn, maxMove)))
        {
            return memo[(startRow, startColumn, maxMove)];
        }
        if (maxMove == 0)
        {
            return (startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n) ? 1 : 0;
        }
        int res = 0;
        if (startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n)
        {
            return 1;
        }
        res += dfs(m, n, maxMove - 1, startRow + 1, startColumn);
        res %= MOD;
        res += dfs(m, n, maxMove - 1, startRow - 1, startColumn);
        res %= MOD;
        res += dfs(m, n, maxMove - 1, startRow, startColumn + 1);
        res %= MOD;
        res += dfs(m, n, maxMove - 1, startRow, startColumn - 1);
        res %= MOD;
        memo.Add((startRow, startColumn, maxMove), res);
        return res;
    }

    public int FindPaths(int m, int n, int maxMove, int startRow, int startColumn)
    {
        memo = new();
        return dfs(m, n, maxMove, startRow, startColumn);
    }
}