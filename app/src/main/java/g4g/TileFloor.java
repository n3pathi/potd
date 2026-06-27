package g4g;

public class TileFloor {
    static long countWays(int n, int m) {
        if (n == 0 || m == 0) {
            return 0;
        }
        if (n < m) {
            return 1;
        }
        if (n == m) {
            return 2;
        }
        long[] dp = new long[n + 1];
        final long mod = 1_000_000_000_000L + 7;
        for (int i = 1; i < dp.length; i++) {
            if (i < m) {
                dp[i] = 1;
            } else if (i == m) {
                dp[i] = 2;
            } else {
                dp[i] = (dp[i - 1] + dp[i - m]) % mod;
            }
        }
        return dp[n];
    }
    
}
