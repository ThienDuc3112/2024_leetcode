namespace FindContentChildren;
public class Solution
{
    public int FindContentChildren(int[] g, int[] s)
    {
        Array.Sort(g);
        Array.Sort(s);
        int i = 0;
        int j = 0;
        int res = 0;
        while (i < g.Length && j < s.Length)
        {
            if (g[i] <= s[j])
            {
                i++;
                j++;
                res++;
            }
            else
            {
                j++;
            }
        }
        return res;
    }
}