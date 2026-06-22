package leetcode;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MaxManhattanDistanceTest {
    private final MaxManhattanDistance sut = new MaxManhattanDistance();

    @ParameterizedTest
    @MethodSource("dataProvider")
    void maxManhattanDistance(String moves, int expected) {
        assertEquals(expected, sut.maxManhattanDistance(moves));
    }

    static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of("L_D_", 4),
                Arguments.of("U_R", 3),
                Arguments.of("___", 3),
                Arguments.of(" ", 0),
                Arguments.of(null, 0)
        );
    }
}