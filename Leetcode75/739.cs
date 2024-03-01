namespace DailyTemperatures;
public class Solution
{
    public int[] DailyTemperatures(int[] temperatures)
    {
        Stack<(int temp, int idx)> monoStack = new();
        int[] res = new int[temperatures.Length];
        for (int i = temperatures.Length - 1; i >= 0; i--)
        {
            int wait = 0;
            while (monoStack.Any() && monoStack.Peek().temp <= temperatures[i])
            {
                monoStack.Pop();
                wait++;
            }
            res[i] = monoStack.Any() ? monoStack.Peek().idx - i: 0;
            monoStack.Push((temperatures[i], i));
        }
        return res;
    }
}
