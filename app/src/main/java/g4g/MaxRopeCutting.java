package g4g;
/**
 * Given a rope of length n meters, cut it into multiple smaller ropes such that the product
 * of their lengths is maximized. At least one cut is mandatory.
 */
public class MaxRopeCutting {
    public int maxProduct(int n) {
        if (n < 2) return 0; // no valid cut feasible
        if (n == 2) return 1;
        if (n == 3) return 2;

        int[] dp = new int[n + 1];
        dp[0] = 0;
        dp[1] = 1;
        dp[2] = 2;
        dp[3] = 3;
        for (int i = 4; i < dp.length; i++) {
            int max = 0;
            for (int j = 1; j <= i / 2; j++) {
                int prod = dp[j] * dp[i - j];
                if (prod > max) {
                    max = prod;
                }
            }
            dp[i] = max;
        }
        return dp[n];
    }
}
