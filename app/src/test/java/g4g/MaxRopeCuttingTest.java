package g4g;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.*;

class MaxRopeCuttingTest {
    private MaxRopeCutting sut = new MaxRopeCutting();

    @ParameterizedTest
    @MethodSource("provideTestData")
    void testMaxProduct(int n, int expected) {
        assertEquals(expected, sut.maxProduct(n));
    }

    private static Stream<Arguments> provideTestData() {
        return Stream.of(
                Arguments.of(0, 0),
                Arguments.of(1, 0),
                Arguments.of(2, 1),
                Arguments.of(3, 2),
                Arguments.of(4, 4),
                Arguments.of(5, 6),
                Arguments.of(10, 36)
        );
    }
}
