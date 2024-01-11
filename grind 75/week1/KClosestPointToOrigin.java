import java.util.ArrayList;
import java.util.Comparator;
import java.util.PriorityQueue;

// Surprisingly, I didn't have to check the solution
// But I still need to check how to implements comparator
class Solution {
    private class PointComparer implements Comparator<int[]> {
        @Override
        public int compare(int[] a, int[] b) {
            double ad = Math.sqrt(a[0] * a[0] + a[1] * a[1]);
            double bd = Math.sqrt(b[0] * b[0] + b[1] * b[1]);
            if (ad > bd)
                return 1;
            if (ad < bd)
                return -1;
            return 0;
        }

    }

    public int[][] kClosest(int[][] points, int k) {
        PriorityQueue<int[]> priQ = new PriorityQueue<>(new PointComparer());
        for (int[] point : points) {
            priQ.add(point);
        }
        ArrayList<int[]> res = new ArrayList<>();
        for (int i = 0; i < k; i++) {
            res.add(priQ.poll());
        }
        return res.toArray(new int[res.size()][]);

    }
}