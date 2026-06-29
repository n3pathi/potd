package g4g;

import java.util.Arrays;

public class MaxDotProduct {
    static int solve(int[] a, int[] b) {
        int n = a.length;
        int m = b.length;
        
        int[][] dp = new int[n + 1][m + 1];
        for (int i = 0; i < dp.length; i++) {
            Arrays.fill(dp[i], Integer.MIN_VALUE);
            dp[i][0] = 0;
        }
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                int currentProduct = a[i - 1] * b[j - 1] + dp[i - 1][j - 1];
                final int skipRow = dp[i - 1][j];
                dp[i][j] = Math.max(currentProduct, skipRow);
            }
        }
        return dp[n][m];
    }
}
