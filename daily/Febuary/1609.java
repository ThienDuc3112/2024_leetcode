import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    public boolean isEvenOddTree(TreeNode root) {
        Queue<TreeNode> queue = new ArrayDeque<>();
        int level = 0;
        queue.add(root);
        while (!queue.isEmpty()) {
            int len = queue.size();
            int previous = level % 2 == 0 ? Integer.MIN_VALUE : Integer.MAX_VALUE;
            for (int i = 0; i < len; i++) {
                TreeNode cur = queue.poll();
                if (level % 2 == 0) {
                    if (cur.val <= previous || cur.val % 2 == 0)
                        return false;
                    previous = cur.val;
                    if (cur.left != null)
                        queue.add(cur.left);
                    if (cur.right != null)
                        queue.add(cur.right);
                } else {
                    if (cur.val >= previous || cur.val % 2 == 1)
                        return false;
                    previous = cur.val;
                    if (cur.left != null)
                        queue.add(cur.left);
                    if (cur.right != null)
                        queue.add(cur.right);
                }
            }
            level++;
        }
        return true;
    }
}
