import java.util.HashMap;

// Inefficient
// When check depth can check the diameter also
class Solution {
    private HashMap<TreeNode, Integer> height;

    private int maxDepth(TreeNode root) {
        if (root == null)
            return -1;
        if (height.containsKey(root))
            return height.get(root);
        height.put(root, Math.max(maxDepth(root.left), maxDepth(root.right)) + 1);
        return height.get(root);
    }

    private int rDiameterOfBinaryTree(TreeNode root) {
        if (root == null)
            return 0;
        return Math.max(maxDepth(root.left) + maxDepth(root.right) + 2,
                Math.max(rDiameterOfBinaryTree(root.left), rDiameterOfBinaryTree((root.right))));
    }

    public int diameterOfBinaryTree(TreeNode root) {
        height = new HashMap<>();
        return rDiameterOfBinaryTree(root);
    }

    private int result;

    private int depth(TreeNode root) {
        if (root == null)
            return 0;
        int left = depth(root.left);
        int right = depth(root.right);
        result = Math.max(result, left + right);
        return Math.max(left, right) + 1;
    }

    public int betterDiameterOfBinaryTree(TreeNode root) {
        result = Integer.MIN_VALUE;
        depth(root);
        return result;
    }
}