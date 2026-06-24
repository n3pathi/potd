package g4g;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.PriorityQueue;

import static java.util.Comparator.comparingInt;

/**
 * Given a matrix mat[][] of size n × n, where mat[i][j] represents the maximum number of steps a rat can jump either
 * forward (right) or downward from that cell, find a path for the rat to reach from the top-left cell (0, 0) to the
 * bottom-right cell (n - 1, n - 1). A cell containing 0 is blocked and cannot be used in the path. It is guaranteed
 * that the cell mat[n-1][n-1] is not 0.
 * <p>
 * Return an n × n matrix where 1 represents the cells included in the path and 0 represents the remaining cells.
 * If no valid path exists, return [[-1]].
 * <p>
 * Note: If multiple valid paths exist, choose the path with the shortest possible jumps first. For the same jump
 * length, moving forward (right) should be preferred over moving downward.
 */
public class RatMaze {
    public List<List<Integer>> shortestDist(int[][] mat) {
        if (null == mat) {
            return Collections.emptyList();
        }
        int n = mat.length;
        int m = mat[0].length;
        if (n == 0 || m == 0) {
            return Collections.emptyList();
        }
        
        PriorityQueue<Node> queue = new PriorityQueue<>(comparingInt((Node o) -> o.col)
                .thenComparingInt(o -> o.row));
        queue.offer(new Node(0, 0, null));
        boolean[][] visited = new boolean[n][m];
        visited[0][0] = true;
        while (!queue.isEmpty()) {
            Node cur = queue.poll();
            visited[cur.row][cur.col] = true;
            if (cur.row == n - 1 && cur.col == m - 1) {
                return prepareResult(mat, cur);
            }
            
            int reachability = mat[cur.row][cur.col];
            for (int i = cur.row; i < n; i++) {
                for (int j = cur.col; j < m; j++) {
                    int dist = (i - cur.row) + (j - cur.col);
                    if (dist <= reachability && !visited[i][j] && mat[i][j] > 0) {
                        queue.offer(new Node(i, j, cur));
                    }
                }
            }
        }
        return Collections.emptyList();
    }
    
    private static List<List<Integer>> prepareResult(int[][] mat, Node cur) {
        List<List<Integer>> res = new ArrayList<>();
        int[][] t = new int[mat.length][mat[0].length];
        while (cur != null) {
            t[cur.row][cur.col] = 1;
            cur = cur.from;
        }
        for (int i = 0; i < mat.length; i++) {
            res.add(new ArrayList<>());
            for (int j = 0; j < mat[0].length; j++) {
                res.get(i).add(t[i][j]);
            }
        }
        return res;
    }
    
    private static record Node(int row, int col, Node from) {
    }
}
