package g4g;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

class ZeroCoverageTest {
    private static final ZeroCoverage sut = new ZeroCoverage();

    @ParameterizedTest
    @MethodSource("dataProvider")
    void findCoverage(int[][] mat, int expected) {
        assertEquals(expected, sut.findCoverage(mat));
    }

    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(
                        new int[][]{
                                {1, 1, 1, 0},
                                {1, 0, 0, 1}
                        },
                        8
                ),
                Arguments.of(
                        new int[][]{},
                        0
                ),
                Arguments.of(
                        new int[][]{
                                {1, 1, 1},
                                {1, 1, 1}},
                        0
                ),
                Arguments.of(
                        new int[][]{
                                {0, 1, 0},
                                {0, 1, 1},
                                {0, 0, 0}
                        },
                        6
                ),
                Arguments.of(
                        new int[][]{
                                {0, 0, 0},
                                {0, 0, 0},
                                {0, 0, 0}
                        },
                        0
                )
        );
    }
}