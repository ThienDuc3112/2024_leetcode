namespace LevelOrder;
public class Solution
{
    public IList<IList<int>> LevelOrder(TreeNode root)
    {
        List<IList<int>> res = new();
        Queue<TreeNode?> queue = new();
        queue.Enqueue(root);
        while (queue.Any())
        {
            int len = queue.Count;
            List<int> list = new();
            for (int i = 0; i < len; i++)
            {
                var cur = queue.Dequeue();
                if (cur != null)
                {
                    list.Add(cur.val);
                    queue.Enqueue(cur.left);
                    queue.Enqueue(cur.right);
                }
            }
            if (list.Count > 0) res.Add(list);
        }
        return res;
    }
}