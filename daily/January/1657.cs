// Already solved before
namespace CloseStrings;
public class Solution
{
    public bool CloseStrings(string word1, string word2)
    {
        if (word1.Length != word2.Length)
        {
            return false;
        }
        Dictionary<char, int> w1 = new();
        Dictionary<char, int> w2 = new();
        for (int i = 0; i < word1.Length; i++)
        {
            if (!w1.ContainsKey(word1[i]))
            {
                w1.Add(word1[i], 0);
            }
            if (!w2.ContainsKey(word2[i]))
            {
                w2.Add(word2[i], 0);
            }
            w1[word1[i]]++;
            w2[word2[i]]++;

        }
        HashSet<char> w1Set = w1.Keys.ToHashSet();
        HashSet<char> w2Set = w2.Keys.ToHashSet();
        HashSet<int> w1Values = w1.Values.ToHashSet();
        HashSet<int> w2Values = w2.Values.ToHashSet();
        Dictionary<int, int> w1ValueCount = new();
        Dictionary<int, int> w2ValueCount = new();
        foreach (var entry in w1)
        {
            if (!w1ValueCount.ContainsKey(entry.Value))
            {
                w1ValueCount.Add(entry.Value, 0);
            }
            w1ValueCount[entry.Value]++;
        }
        foreach (var entry in w2)
        {
            if (!w2ValueCount.ContainsKey(entry.Value))
            {
                w2ValueCount.Add(entry.Value, 0);
            }
            w2ValueCount[entry.Value]++;
        }

        return w1Set.SetEquals(w2Set) && w1Values.SetEquals(w2Values) && w1ValueCount.ToHashSet().SetEquals(w2ValueCount.ToHashSet());
    }
}