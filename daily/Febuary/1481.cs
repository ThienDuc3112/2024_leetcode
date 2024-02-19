namespace FindLeastNumOfUniqueInts;
public class Solution
{
    public int FindLeastNumOfUniqueInts(int[] arr, int k)
    {
        Dictionary<int, int> freqMap = new();
        for (int i = 0; i < arr.Length; i++)
        {
            if (!freqMap.ContainsKey(arr[i])) freqMap.Add(arr[i], 0);
            freqMap[arr[i]]++;
        }
        List<int> freqList = new();
        foreach (var kvpair in freqMap)
            freqList.Add(kvpair.Value);

        freqList.Sort();
        foreach (var thing in freqList) System.Console.WriteLine(thing);
        int index = 0;
        while (k > 0 && index < freqList.Count)
        {
            freqList[index]--;
            if (freqList[index] <= 0)
            {
                index++;
            }
            k--;
        }
        return freqList.Count - index;
    }
}