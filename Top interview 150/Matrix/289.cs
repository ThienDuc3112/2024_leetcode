namespace GameOfLife;
public class Solution
{
    bool Alive(int[][] board, int i, int j)
    {
        int neighbour = 0;
        for (int a = i - 1; a <= i + 1; a++)
        {
            if (a < 0 || a >= board.Length) continue;
            for (int b = j - 1; b <= j + 1; b++)
            {
                if (b < 0 || b >= board[a].Length) continue;
                if (a == i && b == j) continue;
                if (board[a][b] == 1) neighbour++;
            }
        }
        // System.Console.WriteLine($"Neighbour count: {neighbour}\t i,j: ({i},{j})");
        if (board[i][j] == 0)
        {
            return neighbour == 3;
        }
        else
        {
            return neighbour == 2 || neighbour == 3;
        }
    }
    public void GameOfLife(int[][] board)
    {
        int[][] res = new int[board.Length][];
        for (int i = 0; i < board.Length; i++)
        {
            res[i] = new int[board[i].Length];
            for (int j = 0; j < board[i].Length; j++)
            {
                if (Alive(board, i, j)) res[i][j] = 1;
            }
        }
        for (int i = 0; i < board.Length; i++)
        {
            for (int j = 0; j < board[i].Length; j++)
            {
                board[i][j] = res[i][j];
            }
        }
    }
}