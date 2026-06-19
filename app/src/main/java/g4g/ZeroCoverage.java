package g4g;

/**
 * Given a binary matrix mat[][] containing only 0s and 1s, find the total coverage of all 0's.
 * The coverage of a particular 0 cell is defined by checking 1's in its four directions (left, right, up, and down).
 * For each direction, if there is at least one 1 anywhere between the 0 and the boundary of the matrix,
 * the coverage increases by one.
 * <p>
 * Return the sum of the coverage values for all 0 cells in the matrix.
 *
 * @see 'https://www.geeksforgeeks.org/problems/coverage-of-all-zeros-in-a-binary-matrix4024/1'
 */
public class ZeroCoverage {
    public int findCoverage(int[][] mat) {
        int result = 0;
        for (int i = 0; i < mat.length; i++) {
            for (int j = 0; j < mat[0].length; j++) {
                if (mat[i][j] == 0) {
                    result += searchRow(mat, i, j);
                    result += searchCol(mat, i, j);
                }
            }
        }
        return result;
    }

    private int searchCol(int[][] mat, int r, int c) {
        int result = 0;
        // search above
        for (int i = r - 1; i >= 0; i--) {
            if (mat[i][c] == 1) {
                result += 1;
                break;
            }
        }
        // search below
        for (int i = r + 1; i < mat.length; i++) {
            if (mat[i][c] == 1) {
                result += 1;
                break;
            }
        }
        return result;
    }

    private int searchRow(int[][] mat, int r, int c) {
        int result = 0;
        // search left
        for (int i = c - 1; i >= 0; i--) {
            if (mat[r][i] == 1) {
                result += 1;
                break;
            }
        }
        // search right
        for (int i = c + 1; i < mat[r].length; i++) {
            if (mat[r][i] == 1) {
                result += 1;
                break;
            }
        }
        return result;
    }
}
