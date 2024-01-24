namespace MaxLength;
// Look at the topic
public class Solution
{
    bool contained(string str, HashSet<char> appeared)
    {
        var strSet = str.ToCharArray().ToHashSet();
        if (strSet.Count < str.Length) return true;
        foreach (char c in str)
        {
            if (appeared.Contains(c)) return true;
        }
        return false;
    }
    int rMaxLength(IList<string> arr, int from, HashSet<char> appeared)
    {
        int res = 0;
        if (from >= arr.Count) return res;
        if (contained(arr[from], appeared))
        {
            return rMaxLength(arr, from + 1, appeared);
        }
        foreach (char c in arr[from])
        {
            appeared.Add(c);
        }
        int with = arr[from].Length + rMaxLength(arr, from + 1, appeared);
        foreach (char c in arr[from])
        {
            appeared.Remove(c);
        }
        int without = rMaxLength(arr, from + 1, appeared);
        return Math.Max(with, without);
    }
    public int MaxLength(IList<string> arr)
    {
        return rMaxLength(arr, 0, new());
    }
}