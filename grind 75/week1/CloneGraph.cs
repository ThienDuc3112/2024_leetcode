namespace CloneGraph;
public class Solution
{
    // Solution with hint from answer
    public Node CloneGraph(Node node)
    {
        if (node == null) return null;
        Dictionary<Node, (Node copy, bool Visited)> Visited = new();
        Queue<Node> q = new();
        q.Enqueue(node);
        Visited[node] = (new(node.val), false);
        while (q.Any())
        {
            Node cur = q.Dequeue();
            if (Visited[cur].Visited) continue;
            foreach (var neighbour in cur.neighbors)
            {
                Node copy = new(neighbour.val);
                if (!Visited.ContainsKey(neighbour)) Visited[neighbour] = (copy, false);
                copy = Visited[neighbour].copy;
                Visited[cur].copy.neighbors.Add(copy);
                q.Enqueue(neighbour);
            }
            Visited[cur] = (Visited[cur].copy, true);
        }
        return Visited[node].copy;
    }
}