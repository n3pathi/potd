package g4g;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.List;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class EqualizePrefixSumTest {

    private static final EqualizePrefixSum SUT = new EqualizePrefixSum();

    @ParameterizedTest
    @MethodSource("dataProvider")
    void optimalArray(int[] input, List<Integer> expected) {
        assertEquals(expected, SUT.optimalArray(input));
    }

    static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(new int[]{1, 6, 9, 12}, List.of(0, 5, 8, 14)),
                Arguments.of(new int[]{1, 1, 1, 7, 7, 10, 19}, List.of(0, 0,
                        0, 6, 12, 21, 33))
        );
    }
}