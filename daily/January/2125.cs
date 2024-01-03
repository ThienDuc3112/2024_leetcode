namespace NumberOfBeams;
public class Solution
{
    public int NumberOfBeams(string[] bank)
    {
        int lastCount = 0;
        int res = 0;
        foreach (string row in bank)
        {
            int current = 0;
            foreach (char c in row)
            {
                if (c == '1') current++;
            }
            if (current != 0)
            {
                res += current * lastCount;
                lastCount = current;
            }
        }
        return res;
    }
}