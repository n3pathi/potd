package g4g;

/**
 * There are infinitely many people standing in a row, indexed from 1. The strength of the person at index i is i².
 * <p>
 * Given a strength p, determine the maximum number of people that can be defeated.
 * A person with strength x can be defeated only if p ≥ x, after which the strength p decreases by x.
 */
public class MaxPeopleDefeated {
    int solve(int p) {
        long lo = 0;
        long hi = 1;
        while (sumOfSquare(hi) <= p) {
            hi *= 2;
        }

        while (lo < hi) {
            long mid = (lo + hi + 1) / 2;
            if (sumOfSquare(mid) <= p) {
                lo = mid;
            } else {
                hi = mid - 1;
            }
        }
        return (int) lo;
    }

    private long sumOfSquare(long k) {
        return (k * (k + 1) * (2 * k + 1)) / 6;
    }
}
