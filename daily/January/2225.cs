namespace FindWinners;
public class Solution
{
    public IList<IList<int>> FindWinners(int[][] matches)
    {
        Dictionary<int, int> LoseCount = new();
        foreach (var match in matches)
        {
            if (!LoseCount.ContainsKey(match[0]))
            {
                LoseCount.Add(match[0], 0);
            }
            if (!LoseCount.ContainsKey(match[1]))
            {
                LoseCount.Add(match[1], 0);
            }
            LoseCount[match[1]]++;
        }
        List<int>[] res = { new(), new() };
        foreach (var player in LoseCount)
        {
            if (player.Value == 0) res[0].Add(player.Key);
            if (player.Value == 1) res[1].Add(player.Key);
        }
        res[0].Sort();
        res[1].Sort();
        return res;
    }
}