package leetcode;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MatchingSubarrayTest {
    private MatchingSubarray sut = new MatchingSubarray();

    @ParameterizedTest
    @MethodSource("dataProvider")
    void countValidSubarrays(int[] nums, int x, int expected) {
        assertEquals(expected, sut.countValidSubarrays(nums, x));
    }

    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(new int[]{1, 100, 1}, 1, 4),
                Arguments.of(new int[]{1}, 2, 0),
                Arguments.of(new int[]{15,5,25}, 5, 1)
        );
    }
}