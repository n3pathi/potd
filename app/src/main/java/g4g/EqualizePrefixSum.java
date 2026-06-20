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

        long[] prefixSum = new long[arr.length];
        prefixSum[0] = arr[0];
        for (int i = 1; i < arr.length; i++) {
            prefixSum[i] = prefixSum[i - 1] + arr[i];
        }

        for (int i = 1; i < arr.length; i++) {
            int medianIdx = i / 2;
            int median = arr[medianIdx];

            long leftSum = medianIdx > 0 ? prefixSum[medianIdx - 1] : 0;
            int leftCount = medianIdx;

            long rightSum = prefixSum[i] - leftSum;
            int totalElements = i + 1;
            int rightCount = totalElements - medianIdx;

            long ops = ((leftCount * median) - leftSum) + (rightSum - (rightCount * median));
            result.add((int) ops);
        }
        return result;
    }

}
