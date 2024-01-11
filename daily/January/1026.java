// Why am I so dumb to not think of this?
class Solution {
    int DFS(TreeNode root, int max, int min) {
        if (root == null)
            return max - min;
        max = Math.max(max, root.val);
        min = Math.min(min, root.val);
        int left = DFS(root.left, max, min);
        int right = DFS(root.right, max, min);
        return Math.max(right, left);
    }

    public int maxAncestorDiff(TreeNode root) {
        return DFS(root, root.val, root.val);
    }
}