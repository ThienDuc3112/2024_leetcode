import java.util.LinkedList;
import java.util.Queue;

class Solution {
    private class Coordinate {
        public int x;
        public int y;

        Coordinate(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }

    final int[][] directions = new int[][] { new int[] { 1, 0 }, new int[] { -1, 0 }, new int[] { 0, -1 },
            new int[] { 0, 1 } };

    public int orangesRotting(int[][] grid) {
        int res = 0;
        Queue<Coordinate> q = new LinkedList<Coordinate>();
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] == 2)
                    q.add(new Coordinate(i, j));
            }
        }
        while (!q.isEmpty()) {
            int len = q.size();
            for (int i = 0; i < len; i++) {
                Coordinate coord = q.poll();
                for (int[] dir : directions) {
                    int newX = coord.x + dir[0];
                    int newY = coord.y + dir[1];
                    if (newX < grid.length && newX >= 0 && newY < grid[newX].length && newY >= 0
                            && grid[newX][newY] == 1) {
                        grid[newX][newY] = 2;
                        q.add(new Coordinate(newX, newY));
                    }
                }
            }
            // for (int[] row : grid) {
            // System.out.println(Arrays.toString(row));
            // }
            // System.out.println("==========");
            res++;
        }
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] == 1)
                    return -1;
            }
        }
        if (res == 0)
            return 0;
        return res - 1;
    }
}