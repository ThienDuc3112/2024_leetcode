import java.util.HashMap;

class Solution {
    HashMap<Integer, Integer> minuteMap;

    Integer rAmountOfTime(TreeNode node, Integer parent, int start) {
        if (node == null)
            return null;
        // if (minuteMap.get(node.val) != null)
        // return minuteMap.get(node.val);
        if (node.val == start) {
            minuteMap.put(node.val, 0);
            rAmountOfTime(node.left, node.val, start);
            rAmountOfTime(node.right, node.val, start);
            return 0;
        }
        if (parent == null || minuteMap.get(parent) == null) {
            Integer left = rAmountOfTime(node.left, node.val, start);
            Integer right = rAmountOfTime(node.right, node.val, start);
            int res = 0;
            if (left == null && right == null)
                return null;
            if (left == null) {
                res = right + 1;
                minuteMap.put(node.val, res);
                rAmountOfTime(node.left, node.val, start);
            } else {
                res = left + 1;
                minuteMap.put(node.val, res);
                rAmountOfTime(node.right, node.val, start);
            }
            return res;
        } else {
            minuteMap.put(node.val, minuteMap.get(parent) + 1);
            rAmountOfTime(node.left, node.val, start);
            rAmountOfTime(node.right, node.val, start);
            return minuteMap.get(parent) + 1;
        }
    }

    public int amountOfTime(TreeNode root, int start) {
        minuteMap = new HashMap<>();
        rAmountOfTime(root, null, start);
        int res = 0;
        for (var value : minuteMap.values()) {
            res = Math.max(value, res);
        }
        return res;
    }
    // Quicker way would've been to just turn the thing to graph
    // Best way (1 pass DFS) is to store both nodeDepth and subDepth
    // Then return Math.max(subDepth.get(start), nodeDepth.get(start) +
    // subDepth.get(root) - 1)
}