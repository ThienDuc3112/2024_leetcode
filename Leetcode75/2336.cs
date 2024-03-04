public class SmallestInfiniteSet {
    private int current = 1;
    private PriorityQueue<int, int> pq = new();
    public SmallestInfiniteSet() {}
    
    public int PopSmallest() {
        if(pq.Count > 0 && pq.Peek() < current){
            int res = pq.Dequeue();
            while(pq.Count > 0 && pq.Peek() == res){
                pq.Dequeue();
            }
            return res;
        } else {
            return current++;
        }
    }
    
    public void AddBack(int num) {
        if(num < current){
            pq.Enqueue(num, num);
        }
    }
}
