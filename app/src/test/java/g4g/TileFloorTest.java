package g4g;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

class TileFloorTest {
    @ParameterizedTest
    @MethodSource("dataProvider")
    void countWays(int n, int m, int expected) {
        assertEquals(expected, TileFloor.countWays(n, m));
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(2, 4, 1),
                Arguments.of(4, 4, 2),
                Arguments.of(5, 4, 3),
                Arguments.of(6, 3, 6)
        );
    }
}