public class Graph
{
    private int[][] graph;
    public Graph(int n, int[][] edges)
    {
        graph = new int[n][];
        for (int i = 0; i < graph.Length; i++)
        {
            graph[i] = new int[n];
            Array.Fill(graph[i], -1);
        }
        foreach(int[] edge in edges){
            graph[edge[0]][edge[1]] = edge[2];
        }
    }

    public void AddEdge(int[] edge)
    {
        graph[edge[0]][edge[1]] = edge[2];
    }

    public int ShortestPath(int node1, int node2)
    {
        PriorityQueue<(int node, int distance), int> pq = new();
        int[] dst = new int[graph.Length];
        Array.Fill(dst, int.MaxValue);
        pq.Enqueue((node1, 0), 0);
        while(pq.Count > 0 && dst[node2] == int.MaxValue){
            (int node, int distance) = pq.Dequeue();
            if(distance >= dst[node]) continue;
            dst[node] = distance;
            for(int next = 0; next < graph.Length; next++){
                if(graph[node][next] == -1) continue;
                int nextDistance = distance + graph[node][next];
                pq.Enqueue((next, nextDistance), nextDistance);
            }
        }
        return dst[node2] == int.MaxValue ? -1 : dst[node2];
        
    }
}
