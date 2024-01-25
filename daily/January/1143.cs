namespace LongestCommonSubsequence;
// This is taught in edx DSA IV with 2d array
// I simply forgor :skull:
// Make a int[text1.Length + 1][text2.Length + 1] array
// int[0][x] = 0
// int[x][0] = 0
// if character match, + 1 and go diagonal
// If not, pick the largest from the up, left
// The largest is dp.Last().Last() 
public class Solution
{
    int[][] memo;
    int checkSubSequence(string str1, int len1, string str2, int len2)
    {
        if (len1 >= str1.Length || len2 >= str2.Length) return 0;
        if (memo[len1][len2] != -1) return memo[len1][len2];
        int res;
        if (str1[len1] == str2[len2])
        {
            res = 1 + checkSubSequence(str1, len1 + 1, str2, len2 + 1);
        }
        else
        {
            res = Math.Max(checkSubSequence(str1, len1 + 1, str2, len2), checkSubSequence(str1, len1, str2, len2 + 1));
        }
        memo[len1][len2] = res;
        return res;

    }
    public int LongestCommonSubsequence(string text1, string text2)
    {
        memo = new int[text1.Length][];
        for (int i = 0; i < memo.Length; i++)
        {
            memo[i] = new int[text2.Length];
            for (int j = 0; j < memo[i].Length; j++)
            {
                memo[i][j] = -1;
            }
        }
        return checkSubSequence(text1, 0, text2, 0);
    }
}