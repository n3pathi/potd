package g4g;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MinOpsToIdenticalArrayTest {
    @ParameterizedTest
    @MethodSource("dataProvider")
    void testSolve(int[] a, int[] b, int expected) {
        assertEquals(expected, MinOpsToIdenticalArray.solve(a, b));
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(new int[]{1, 2, 5, 3, 1}, new int[]{1, 3, 5}, 4),
                Arguments.of(new int[]{1, 4}, new int[]{1, 4}, 0),
                Arguments.of(new int[]{1, 1, 1, 1, 1}, new int[]{1, 3, 5}, 6)
        );
    }
}