namespace MinFlips;
public class Solution
{
    public int MinFlips(int a, int b, int c)
    {
        int res = 0;
        for (int i = 0; i < 32; i++)
        {
            int aBit = a & (1 << i);
            int bBit = b & (1 << i);
            int cBit = c & (1 << i);
            if (cBit > 0)
            {
                if (aBit == 0 && bBit == 0) res++;
            }
            else
            {
                if (aBit > 0) res++;
                if (bBit > 0) res++;
            }
        }
        return res;
    }
}