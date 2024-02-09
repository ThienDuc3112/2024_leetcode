import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.List;
import java.util.Queue;

class Solution {
    // Check solution
    public int numSquares(int n) {
        int[] dp = new int[n + 1];
        for (int i = 1; i < dp.length; i++)
            dp[i] = Integer.MAX_VALUE;
        for (int i = 1; i * i <= n; i++) {
            for (int j = i * i; j <= n; j++) {
                dp[j] = Math.min(dp[j], dp[j - i * i] + 1);
            }
        }
        return dp[n];
    }

    // First attempt / bad attempt
    public int badNumSquares(int n) {
        List<Integer> listOfSqare = new ArrayList<>();
        while (listOfSqare.size() * listOfSqare.size() <= n) {
            listOfSqare.add(listOfSqare.size() * listOfSqare.size());
        }
        Queue<int[]> q = new ArrayDeque<>();
        q.add(new int[] { n, 0 });
        while (!q.isEmpty()) {
            int[] cur = q.poll();
            for (int j : listOfSqare) {
                if (cur[0] - j == 0)
                    return cur[1] + 1;
                if (cur[0] - j > 0)
                    q.add(new int[] { cur[0] - j, cur[1] + 1 });
            }
        }
        return 0;
    }
}