package lc;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * You are given an array of positive integers nums.
 * <p>
 * You need to select a of nums which satisfies the following condition:
 * <p>
 * You can place the selected elements in a 0-indexed array such that it follows
 * the pattern: [x, x2, x4, ..., xk/2, xk, xk/2, ..., x4, x2, x] (Note that k can be any non-negative power of 2).
 * For example, [2, 4, 16, 4, 2] and [3, 9, 3] follow the pattern while [2, 4, 8, 4, 2] does not.
 * <p>
 * Return the maximum number of elements in a subset that satisfies these conditions.
 * <p>
 * Constraints:
 * <p>
 * 2 <= nums.length <= 105
 * 1 <= nums[i] <= 109
 */
public class MaximumSubset {
    public int maximumLength(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        
        Map<Long, Integer> freq = new HashMap<>();
        for (final int num : nums) {
            freq.merge((long) num, 1, Integer::sum);
        }
        
        int ones = freq.getOrDefault(1L, 0);
        int res = (ones & 1) == 1 ? ones : ones - 1;
        freq.remove(1L);
        
        List<Long> list = new ArrayList<>(freq.keySet());
        Collections.sort(list);
        for (final Long num : list) {
            int t = 0;
            long x = num;
            while (freq.containsKey(x) && freq.get(x) > 1) {
                t += 2;
                x *= x;
            }
//            top element was never added if single but is counted twice if duplicate
            final int adjustedResult = t + (freq.containsKey(x) ? 1 : -1);
            res = Math.max(res, adjustedResult);
        }
        return res;
    }
}
