namespace LeastInterval;
public class Solution {
    public int LeastInterval(char[] tasks, int n) {
        Dictionary<char, int> counting = new();
        foreach(char task in tasks){
            counting[task] = counting.GetValueOrDefault(task, 0) + 1;
        }
        int res = 0;
        int finished = 0;
        while(finished < tasks.Length){
            PriorityQueue<char, int> pq = new();
            foreach((char k, int v) in counting){
                if(v == 0) continue;
                pq.Enqueue(k, -v);
            }
            int thisCycle = n + 1;
            while(thisCycle > 0 && pq.Count > 0){
                char task = pq.Dequeue();
                counting[task]--;
                finished++;
                res++;
                thisCycle--;
            }
            System.Console.WriteLine($"{finished}");
            if(finished < tasks.Length) {
                res += thisCycle;
            }
        }
        return res;
    }
}
