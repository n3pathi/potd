package g4g;

import java.util.ArrayList;

/**
 * Given a sorted array arr[]. For each i(0 ≤ i ≤ n-1), make all the elements of the array from index 0 to i equal, using the minimum number of operations.
 * <p>
 * In one operation you either increase or decrease the array element by 1. Return an array that contains the minimum number of operations for each i, to accomplish the above task.
 * <p>
 * Note:
 * <p>
 * For each index i, consider the original array without applying modifications made for previous indices.
 * Try to solve the problem using O(1) extra space (excluding the resultant array).
 *
 * @see https://www.geeksforgeeks.org/problems/optimal-array--170647/1
 */
public class EqualizePrefixSum {
    public ArrayList<Integer> optimalArray(int[] arr) {
        ArrayList<Integer> result = new ArrayList<>();
        result.add(0); // single entry is always equal

        for (int i = 1; i < arr.length; i++) {
            int ops = 0;
            int medianIdx = (0 + i) / 2;
            int median = arr[medianIdx];
            for (int j = 0; j <= i; j++) {
                ops += Math.abs(arr[j] - median);
            }
            result.add(ops);
        }
        return result;
    }

}
