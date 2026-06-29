package g4g;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MaxDotProductTest {
    @ParameterizedTest
    @MethodSource("dataProvider")
    void testSolve(int[] a, int[] b, int expected) {
        assertEquals(expected, MaxDotProduct.solve(a, b));
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(
                        new int[]{2, 3, 1, 7, 8},
                        new int[]{3, 6, 7},
                        107
                ),
                Arguments.of(
                        new int[]{1, 2, 3},
                        new int[]{4},
                        12
                ),
                Arguments.of(
                        new int[]{5, 8},
                        new int[]{9, 1},
                        53
                ),
                Arguments.of(
                        new int[]{34, 96, 73, 78, 69, 72, 79, 63, 41, 61, 20, 42, 56, 50, 16, 44, 18, 97, 76, 98, 77, 14, 49, 20, 14, 42, 15, 38
                                , 10, 22, 27, 70, 34, 50, 54, 68, 49, 37, 62, 78, 69, 72, 60, 25, 3, 94, 80, 77, 82, 71, 89, 75, 4, 68, 45, 2, 79, 74, 11, 54, 39, 48, 18, 78, 23, 32, 24, 71, 26, 86, 97, 5, 14, 82, 35, 29, 71, 51, 4, 90, 48, 63, 5, 60, 51, 83, 18, 37, 59, 2, 84, 51, 61, 78, 56, 55},
                        new int[]{40, 10, 74, 6, 40, 9, 50, 53, 19, 6, 22, 86, 22, 21, 37, 22, 71, 47, 60, 4, 6, 22, 34, 71, 21, 46, 51, 58, 88,
                                56, 92, 82, 41, 84, 92},
                        125517
                )
        );
    }
}