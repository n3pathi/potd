package g4g;

import java.util.stream.Stream;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.*;

class CountMatchingSubsequencesTest {
    @ParameterizedTest
    @MethodSource("dataProvider")
    void testCountWays(String s1, String s2, int expected) {
        final int actual = CountMatchingSubsequences.countWays(s1, s2);
        assertEquals(expected, actual);
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of("geeksforgeeks", "gks", 4),
                Arguments.of("problemoftheday", "geek", 0),
                Arguments.of("geek", "", 1),
                Arguments.of("", "geek", 0)
        );
    }
}