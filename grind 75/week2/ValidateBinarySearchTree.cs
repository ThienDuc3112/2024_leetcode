namespace IsValidBST;
public class Solution
{
    bool BoundedIsValidBST(TreeNode? root, long lowerBound, long upperBound)
    {
        if (root == null) return true;
        if (root.val <= lowerBound || root.val >= upperBound) return false;
        bool left = true, right = true;
        if (root.left != null)
        {
            left = BoundedIsValidBST(root.left, lowerBound, root.val);
            if (left == false) return false;
        }
        if (root.right != null)
        {
            right = BoundedIsValidBST(root.right, root.val, upperBound);
        }
        return left && right;

    }
    public bool IsValidBST(TreeNode root)
    {
        return BoundedIsValidBST(root, long.MinValue, long.MaxValue);
    }
}