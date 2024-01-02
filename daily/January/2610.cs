namespace FindMatrix;
public class Solution
{
    public IList<IList<int>> FindMatrix(int[] nums)
    {
        Dictionary<int, int> freqency = new();
        foreach (var num in nums)
        {
            if (!freqency.ContainsKey(num))
            {
                freqency[num] = 0;
            }
            freqency[num]++;
        }
        IList<IList<int>> res = new List<IList<int>>();
        while (freqency.Any())
        {
            List<int> cur = new();
            foreach (var freq in freqency)
            {
                cur.Add(freq.Key);
                if (freq.Value == 1) freqency.Remove(freq.Key);
                else freqency[freq.Key]--;
            }
            res.Add(cur);
        }
        return res;
    }
}