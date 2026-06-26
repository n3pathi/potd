package g4g;

public class CountMatchingSubsequences {
    public static int countWays(String s1, String s2) {
        if (null == s1 || null == s2) {
            return 0;
        }
        int n = s1.length();
        int m = s2.length();
        int[] dp = new int[m + 1];
        dp[0] = 1;
        for (int i = 1; i <= n; i++) {
            char c = s1.charAt(i - 1);
            for (int j = 1; j <= m; j++) {
                if (s2.charAt(j - 1) == c) {
                    dp[j] += dp[j - 1];
                }
            }
        }
        return dp[m];
    }
}
