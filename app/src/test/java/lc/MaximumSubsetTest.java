package lc;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.*;

class MaximumSubsetTest {
    private static final MaximumSubset sut = new MaximumSubset();
    
    @ParameterizedTest
    @MethodSource("dataProvider")
    void maximumLength(int[] nums, int expected) {
        assertEquals(expected, sut.maximumLength(nums));
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(new int[]{5, 4, 1, 2, 2}, 3),
                Arguments.of(new int[]{1, 3, 2, 4}, 1),
                Arguments.of(new int[]{14, 14, 196, 196, 38416, 38416}, 5),
                Arguments.of(new int[]{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}, 9)
        );
    }
}