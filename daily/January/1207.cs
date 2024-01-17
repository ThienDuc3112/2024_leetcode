namespace UniqueOccurrences;
public class Solution
{
    public bool UniqueOccurrences(int[] arr)
    {
        Dictionary<int, int> freq = new();
        foreach (int num in arr)
        {
            if (!freq.ContainsKey(num))
            {
                freq[num] = 0;
            }
            freq[num]++;
        }
        var set = freq.Values.ToHashSet();
        return set.Count == freq.Count;
    }
}