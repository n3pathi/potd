package lc;

/**
 * You are given an integer array nums and an integer digit x.
 * <p>
 * A nums[l..r] is considered valid if the sum of its elements satisfies both of the following conditions:
 * <p>
 * The first digit of the sum is equal to x.
 * The last digit of the sum is equal to x.
 * <p>
 * Return the number of valid subarrays.
 */
public class MatchingSubarray {
    public int countValidSubarrays(int[] nums, int x) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        if (x < 0 || x > 9) {
            return 0;
        }
        int ans = 0;
        for (int i = 0; i < nums.length; i++) {
            int sum = 0;
            for (int j = i; j < nums.length; j++) {
                sum += nums[j];
                int firstDigit = Integer.parseInt(String.valueOf(Integer.toString(sum).charAt(0)));
                int lastDigit = sum % 10;
                if (firstDigit == x && lastDigit == x) {
                    ans++;
                }
            }
        }
        return ans;
    }
}
