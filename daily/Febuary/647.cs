using System.Text;

namespace CountSubstrings;
public class Solution
{
    // Extend palindrome
    public int CountSubstrings(string s)
    {
        int res = 0;
        for (int i = 0; i < s.Length; i++)
        {
            res += extendPalindrome(s, i, i);
            res += extendPalindrome(s, i, i + 1);
        }
        return res;
    }
    private int extendPalindrome(string s, int i, int j)
    {
        int res = 0;
        while (i >= 0 && j < s.Length && s[i] == s[j])
        {
            res++;
            i--;
            j++;
        }
        return res;
    }

    // First attempt
    public int BadCountSubstrings(string s)
    {
        int res = 0;
        for (int i = 0; i < s.Length; i++)
        {
            StringBuilder forward = new("");
            StringBuilder backward = new("");
            for (int j = i; j < s.Length; j++)
            {
                forward.Append(s[j]);
                backward.Insert(0, s[j]);
                if (forward.Equals(backward))
                {
                    res++;
                }
            }
        }
        return res;
    }
}