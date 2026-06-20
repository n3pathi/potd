package g4g;

import java.math.BigInteger;

/**
 * Given two integers a and b in the form of strings. Return the last digit of a^b.
 */
public class LastDigit {
    public int getLastDigit(String a, String b) {
        int lastDigit = lastDigit(a.toCharArray());

        BigInteger bigB = new BigInteger(b);
        if (a.length() == 1 && lastDigit == 0 && bigB.signum() <= 0) {
            throw new ArithmeticException("0^0 is not defined");
        }
        if (bigB.signum() <= 0) {
            throw new IllegalArgumentException("Exponent must be a positive integer");
        }
        int numCycle = getNumCycle(lastDigit);
        int rem = bigB.remainder(BigInteger.valueOf(numCycle)).intValue();
        return getLastDigitAtPosition(lastDigit, rem);
    }

    private int getLastDigitAtPosition(int digit, int position) {
        if (position == 0) {
            position = getNumCycle(digit);
        }
        return (int) Math.pow(digit, position) % 10;
    }


    private int lastDigit(char[] chars) {
        return Integer.parseInt(String.valueOf(chars[chars.length - 1]));
    }

    private int getNumCycle(int lastDigit) {
        switch (lastDigit) {
            case 0, 1, 5, 6:
                return 1;
            case 2, 3, 7, 8:
                return 4;
            case 4, 9:
                return 2;
            default:
                throw new IllegalArgumentException(String.format("bad input: %d", lastDigit));
        }
    }

}
