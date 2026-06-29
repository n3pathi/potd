package lc;

import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

class NumStringsTest {
    
    @ParameterizedTest
    @MethodSource("dataProvider")
    void solve(String[] patterns, String word, int expected) {
        assertEquals(expected, NumStrings.solve(patterns, word));
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(new String[]{"a", "abc", "bc", "d"}, "abc", 3),
                Arguments.of(new String[]{"a", "b", "c"}, "aaaaabbbbb", 2),
                Arguments.of(new String[]{"a", "a", "a"}, "ab", 3)
        );
    }
}