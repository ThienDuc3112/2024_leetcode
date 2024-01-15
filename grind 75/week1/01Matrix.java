import java.util.ArrayDeque;
import java.util.Queue;

class Solution {
    int[] DIR = new int[] { 0, 1, 0, -1, 0 };

    public int[][] updateMatrix(int[][] mat) {
        int m = mat.length, n = mat[0].length; // The distance of cells is up to (M+N)
        Queue<int[]> q = new ArrayDeque<>();
        for (int r = 0; r < m; ++r)
            for (int c = 0; c < n; ++c)
                if (mat[r][c] == 0)
                    q.offer(new int[] { r, c });
                else
                    mat[r][c] = -1; // Marked as not processed yet!

        while (!q.isEmpty()) {
            int[] curr = q.poll();
            int r = curr[0], c = curr[1];
            for (int i = 0; i < 4; ++i) {
                int nr = r + DIR[i], nc = c + DIR[i + 1];
                if (nr < 0 || nr == m || nc < 0 || nc == n || mat[nr][nc] != -1)
                    continue;
                mat[nr][nc] = mat[r][c] + 1;
                q.offer(new int[] { nr, nc });
            }
        }
        return mat;
    }
}

// Not optimized, needed solution
// namespace UpdateMatrix;
// class Solution
// {
// int[] dir = new int[] { 0, 1, 0, -1, 0 };

// bool validSquare(int[][] mat, int i, int j)
// {
// return i >= 0 && i < mat.Length && j >= 0 && j < mat[i].Length;
// }

// void BFS(int[][] mat, int i, int j)
// {
// Queue<(int i, int j)> queue = new();
// HashSet<(int i, int j)> visited = new();
// queue.Enqueue((i, j));
// visited.Add((i, j));
// int step = 0;
// while (queue.Any())
// {
// int len = queue.Count;
// for (int a = 0; a < len; a++)
// {
// var cur = queue.Dequeue();
// if (mat[cur.i][cur.j] < step) continue;
// mat[cur.i][cur.j] = step;
// for (int k = 0; k < 4; k++)
// {
// int x = cur.i + dir[k];
// int y = cur.j + dir[k + 1];
// if (validSquare(mat, x, y) && mat[x][y] != 0 && !visited.Contains((x, y)))
// {
// queue.Enqueue((x, y));
// visited.Add((x, y));
// }
// }
// }
// step++;
// }
// }

// public int[][] UpdateMatrix(int[][] mat)
// {
// for (int i = 0; i < mat.Length; i++)
// {
// for (int j = 0; j < mat[i].Length; j++)
// {
// if (mat[i][j] != 0) mat[i][j] = int.MaxValue;
// }
// }
// for (int i = 0; i < mat.Length; i++)
// {
// for (int j = 0; j < mat[i].Length; j++)
// {
// if (mat[i][j] == 0)
// {
// BFS(mat, i, j);
// }
// }
// }
// return mat;
// }
// }