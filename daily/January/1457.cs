namespace PseudoPalindromicPaths;
public class Solution
{
    int[] counts = new int[9];

    public int PseudoPalindromicPaths(TreeNode root)
    {
        counts = new int[9];
        return Dfs(root);
    }

    int Dfs(TreeNode root)
    {
        counts[root.val - 1]++;
        if (root.left == null && root.right == null)
        {
            int res = CanPalindrome(counts) ? 1 : 0;
            counts[root.val - 1]--;
            return res;
        }
        int LeftCount = root.left == null ? 0 : Dfs(root.left);
        int RightCount = root.right == null ? 0 : Dfs(root.right);
        counts[root.val - 1]--;
        return LeftCount + RightCount;
    }

    bool CanPalindrome(int[] counts)
    {
        int tolerance = 1;
        System.Console.WriteLine(string.Join(", ", counts));
        foreach (int count in counts)
        {
            if (count % 2 == 0) continue;
            if (tolerance <= 0) return false;
            tolerance--;
        }
        return true;
    }
}