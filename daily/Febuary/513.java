import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    public int findBottomLeftValue(TreeNode root) {
        Queue<TreeNode> queue = new ArrayDeque<>();
        queue.add(root);
        int left = root.val;
        while (!queue.isEmpty()) {
            left = queue.peek().val;
            int len = queue.size();
            for (int i = 0; i < len; i++) {
                TreeNode cur = queue.remove();
                if (cur.left != null)
                    queue.add(cur.left);
                if (cur.right != null)
                    queue.add(cur.right);
            }
        }
        return left;
    }
}