package g4g;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * Given an integer n, return all the n digit numbers in increasing order, such that their digits are in strictly
 * increasing order(from left to right).
 */
public class IncreasingDigits {
    public static List<Integer> getIncreasingDigits(int n) {
        if (n <= 0 || n > 9) {
            return Collections.emptyList();
        }
        List<Integer> res = new ArrayList<>();
        int[] digits = new int[n];
        backtrack(res, digits, 0);
        return res;
    }
    
    private static void backtrack(final List<Integer> res, final int[] digits, final int k) {
        if (k == digits.length) {
            int value = 0;
            for (int digit : digits) {
                value = value * 10 + digit;
            }
            res.add(value);
            return;
        }
        int start = 0;
        if (k != 0) {
            start = digits[k - 1];
        }
        for (int i = start + 1; i <= 9; i++) {
            digits[k] = i;
            backtrack(res, digits, k + 1);
        }
    }
}
