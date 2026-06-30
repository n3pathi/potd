package g4g;

/**
 * Given two arrays a[] and b[], find the minimum number of insertions and deletions on the array a[], required to make both the arrays identical.
 * Note: Array b[] is sorted and all its elements are distinct, operations can be performed at any index not necessarily at the end.
 */
public class MinOpsToIdenticalArray {
    static int solve(int[] a, int[] b) {
        int n = a.length;
        int m = b.length;
        int[][] dp = new int[n+1][m+1];
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= m; j++) {
                if (a[i - 1] == b[j - 1]) {
                    dp[i][j] = dp[i - 1][j - 1] + 1;
                } else {
                    dp[i][j] = Math.max(dp[i - 1][j], dp[i][j - 1]);
                }
            }
        }
        int lcs = dp[n][m];
        return (n - lcs) + (m - lcs);
    }
}
