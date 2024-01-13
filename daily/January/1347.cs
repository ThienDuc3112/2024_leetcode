namespace MinSteps;
public class Solution
{
    public int MinSteps(string s, string t)
    {
        Dictionary<char, int> freqMap = new();
        for (int i = 0; i < s.Length; i++)
        {
            if (!freqMap.ContainsKey(s[i])) freqMap.Add(s[i], 0);
            if (!freqMap.ContainsKey(t[i])) freqMap.Add(t[i], 0);
            freqMap[s[i]]++;
            freqMap[t[i]]--;
        }
        int sum = 0;
        foreach ((char key, int value) in freqMap) sum += Math.Abs(value);
        return sum / 2;
    }
}