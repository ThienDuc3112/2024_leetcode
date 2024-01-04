namespace MinOperations;
public class Solution
{
    private int? MinOperationsSingleNum(int amount)
    {
        if (amount == 1) return null;
        if (amount <= 3) return 1;
        if (amount == 4) return 2;
        return 1 + MinOperationsSingleNum(amount - 3);
    }
    public int MinOperations(int[] nums)
    {
        Dictionary<int, int> freqMap = new();
        foreach (int num in nums)
        {
            if (!freqMap.ContainsKey(num))
            {
                freqMap[num] = 0;
            }
            freqMap[num]++;
        }
        int res = 0;
        foreach ((int key, int value) in freqMap)
        {
            int? opNeeded = MinOperationsSingleNum(value);
            if (opNeeded == null) return -1;
            res += (int)opNeeded;
        }
        return res;
    }
}