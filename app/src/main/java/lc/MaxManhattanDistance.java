package lc;

/**
 * You are given a string moves consisting of the characters 'U', 'D', 'L', 'R', and '_'.
 * <p>
 * Starting from the origin (0, 0), each character represents one move on a 2D plane:
 * <ul>
 * <li>'U': Move up by 1 unit.</li>
 * <li>'D': Move down by 1 unit.</li>
 * <li>'L': Move left by 1 unit.</li>
 * <li>'R': Move right by 1 unit.</li>
 * <li>'_': Can be independently replaced with any one of 'U', 'D', 'L', or 'R'.</li>
 * </ul>
 * Return the maximum from the origin that can be achieved after all moves have been performed.
 */
public class MaxManhattanDistance {
    public int maxManhattanDistance(String moves) {
        if (null == moves || moves.trim().length() == 0) {
            return 0;
        }
        int x = 0;
        int y = 0;
        int wildcards = 0;
        char[] chars = moves.toCharArray();
        for (char aChar : chars) {
            switch (aChar) {
                case 'U' -> y++;
                case 'D' -> y--;
                case 'L' -> x--;
                case 'R' -> x++;
                case '_' -> wildcards++;
                default -> throw new IllegalArgumentException("Invalid move " + aChar);
            }
        }
        int absX = Math.abs(x);
        int absY = Math.abs(y);

        if (absX >= absY) {
            absX += wildcards;
        } else {
            absY += wildcards;
        }
        return absX + absY;
    }
}
