class Solution {
    public int minFallingPathSum(int[][] matrix) {
        for (int i = 1; i < matrix.length; i++) {
            for (int j = 0; j < matrix[i].length; j++) {
                matrix[i][j] += Math.min(matrix[i - 1][j],
                        Math.min(matrix[i - 1][Math.max(0, j - 1)], matrix[i - 1][Math.min(matrix.length - 1, j + 1)]));
            }
        }
        int res = Integer.MAX_VALUE;
        for (int sum : matrix[matrix.length - 1]) {
            res = Math.min(sum, res);
        }
        return res;
    }

    // First solution
    // Could've done inplace

    // public int minFallingPathSum(int[][] matrix) {
    // int[][] dp = new int[matrix.length][matrix[0].length];
    // for (int i = 0; i < dp.length; i++) {
    // for (int j = 0; j < dp[i].length; j++) {
    // if (i == 0)
    // dp[i][j] = matrix[i][j];
    // else {
    // dp[i][j] = Integer.MAX_VALUE;
    // for (int k = j - 1; k <= j + 1 && k < dp[i].length; k++) {
    // if (k < 0)
    // continue;
    // dp[i][j] = Math.min(dp[i][j], dp[i - 1][k] + matrix[i][j]);
    // }
    // }
    // }
    // }
    // int res = Integer.MAX_VALUE;
    // for (int sum : dp[dp.length - 1]) {
    // res = Math.min(sum, res);
    // }
    // return res;
    // }
}