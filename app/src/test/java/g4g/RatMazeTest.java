package g4g;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.Collections;
import java.util.List;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class RatMazeTest {
    private final RatMaze sut = new RatMaze();
    
    @ParameterizedTest
    @MethodSource("dataProvider")
    void testShortestDist(int[][] mat, List<List<Integer>> expected) {
        assertEquals(expected, sut.shortestDist(mat));
    }
    
    static Stream<Arguments> dataProvider() {
        return Stream.of(
//                1
                Arguments.of(
                        new int[][]{
                                {2, 1, 0, 0},
                                {3, 0, 0, 1},
                                {0, 1, 0, 1},
                                {0, 0, 0, 1}
                        },
                        List.of(
                                List.of(1, 0, 0, 0),
                                List.of(1, 0, 0, 1),
                                List.of(0, 0, 0, 1),
                                List.of(0, 0, 0, 1)
                        )
                ),
//                2
                Arguments.of(
                        new int[][]{
                                {2, 1, 0, 0},
                                {2, 0, 0, 1},
                                {0, 1, 0, 1},
                                {0, 0, 0, 1},
                        },
                        Collections.emptyList()
                ),
//                3
                Arguments.of(
                        new int[][]{
                                {2, 0, 0},
                                {0, 2, 0},
                                {0, 0, 1}
                        },
                        List.of(
                                List.of(1, 0, 0),
                                List.of(0, 1, 0),
                                List.of(0, 0, 1)
                        )
                ),
//                4
                Arguments.of(
                        new int[][]{
                                {3, 2, 0, 0},
                                {0, 0, 2, 0},
                                {0, 0, 0, 1},
                                {0, 0, 0, 1}
                        },
                        List.of(
                                List.of(1, 0, 0, 0),
                                List.of(0, 0, 1, 0),
                                List.of(0, 0, 0, 1),
                                List.of(0, 0, 0, 1)
                        )
                ),
//                5
                Arguments.of(
                        new int[][]{
                                {2, 0, 0, 0},
                                {0, 2, 0, 0},
                                {0, 0, 2, 0},
                                {0, 0, 0, 1}
                        },
                        List.of(
                                List.of(1, 0, 0, 0),
                                List.of(0, 1, 0, 0),
                                List.of(0, 0, 1, 0),
                                List.of(0, 0, 0, 1)
                        )
                ),
//                6
                Arguments.of(
                        new int[][]{
                                {1, 0, 0},
                                {0, 2, 0},
                                {0, 0, 1}
                        },
                        Collections.emptyList()
                ),
//                7
                Arguments.of(
                        new int[][]{
                                {2, 0, 0, 0},
                                {0, 1, 0, 0},
                                {0, 0, 1, 0},
                                {0, 0, 0, 1}
                        },
                        Collections.emptyList()
                ),
//                8
                Arguments.of(
                        new int[][]{
                                {3, 1, 1, 0},
                                {0, 0, 2, 0},
                                {0, 0, 0, 2},
                                {0, 0, 0, 1}
                        },
                        List.of(
                                List.of(1, 0, 0, 0),
                                List.of(0, 0, 1, 0),
                                List.of(0, 0, 0, 1),
                                List.of(0, 0, 0, 1)
                        )
                ),
//                9
                Arguments.of(
                        new int[][]{
                                {3, 0, 0, 0},
                                {0, 3, 0, 0},
                                {0, 0, 1, 0},
                                {0, 0, 0, 2}
                        },
                        Collections.emptyList()
                )
        );
    }
}