import java.util.ArrayList;
import java.util.List;

class Solution {
    void DFS(TreeNode root, List<Integer> leafs) {
        if (root == null)
            return;
        if (root.left == null && root.right == null) {
            leafs.add(root.val);
            return;
        }
        DFS(root.left, leafs);
        DFS(root.right, leafs);
    }

    public boolean leafSimilar(TreeNode root1, TreeNode root2) {
        List<Integer> leafs1 = new ArrayList<Integer>();
        List<Integer> leafs2 = new ArrayList<Integer>();
        DFS(root1, leafs1);
        DFS(root2, leafs2);
        if (leafs1.size() != leafs2.size())
            return false;
        return leafs1.equals(leafs2);
    }
}