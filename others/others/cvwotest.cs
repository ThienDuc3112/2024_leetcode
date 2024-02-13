namespace Solution;
class Solution {
    public int solution(int[] A) {
        // Implement your solution here
        HashSet<int> set = new();
        foreach(int num in A){
            set.Add(num);
        }
        int i = 1;
        while(true){
            if(!set.Contains(i)) return i;
            i++;
        }
    }
}
