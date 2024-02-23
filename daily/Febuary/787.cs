namespace FindCheapestPrice;
public class Solution
{
    // Check ans
    // Bellman-ford sssp algorithm
    public int FindCheapestPrice(int n, int[][] flights, int src, int dst, int k)
    {
        int[] val = new int[n];
        Array.Fill(val, int.MaxValue);
        val[src] = 0;
        for (int i = 0; i < k; i++)
        {
            int[] clone = (int[])val.Clone();
            foreach (var flight in flights)
            {
                if (val[flight[0]] == int.MaxValue) continue;
                clone[flight[1]] = Math.Min(val[flight[0]] + flight[2], val[flight[1]]);
                System.Console.WriteLine($"{flight[1]} updated to: {clone[flight[1]]}");
            }
            val = clone;
        }
        return val[dst] == int.MaxValue ? -1 : val[dst];
    }
}

// import java.util.ArrayDeque;
// import java.util.HashMap;
// import java.util.Map;
// import java.util.Queue;

// class Solution {
//     private class Item {
//         public int id;
//         public int cost;

//         public Item(int id, int cost) {
//             this.id = id;
//             this.cost = cost;
//         }
//     }

//     public int findCheapestPrice(int n, int[][] flights, int src, int dst, int k) {
//         int[][] graph = new int[n][];
//         for (int i = 0; i < graph.length; i++) {
//             graph[i] = new int[n];
//         }
//         for (int[] relation : flights) {
//             graph[relation[0]][relation[1]] = relation[2];
//         }
//         Map<Integer, Integer> map = new HashMap<>();
//         Queue<Item> queue = new ArrayDeque<>();
//         queue.add(new Item(src, 0));
//         while (!queue.isEmpty()) {
//             if (k < -1)
//                 break;
//             int len = queue.size();
//             for (int i = 0; i < len; i++) {
//                 Item cur = queue.poll();
//                 map.put(cur.id, Math.min(cur.cost, map.getOrDefault(cur.id, Integer.MAX_VALUE)));
//                 for (int j = 0; j < n; j++) {
//                     if (graph[cur.id][j] > 0) {
//                         System.out.println("From " + Integer.toString(cur.id) + " to " + Integer.toString(j) + " for "
//                                 + Integer.toString(cur.cost + graph[cur.id][j]));
//                         queue.add(new Item(j, cur.cost + graph[cur.id][j]));
//                     }
//                 }
//             }
//             k--;
//         }
//         return map.getOrDefault(dst, -1);
//     }
// }