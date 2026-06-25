package g4g;

import java.util.Collections;
import java.util.List;
import java.util.stream.Stream;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

class IncreasingDigitsTest {
    private final IncreasingDigits sut = new IncreasingDigits();
    
    @ParameterizedTest
    @MethodSource("dataProvider")
    void getIncreasingDigits(int n, List<Integer> expected) {
        assertEquals(expected, sut.getIncreasingDigits(n));
    }
    
    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(1, List.of(1, 2, 3, 4, 5, 6, 7, 8, 9)),
                Arguments.of(2, List.of(12, 13, 14, 15, 16, 17, 18, 19, 23, 24, 25, 26, 27, 28,
                        29, 34, 35, 36, 37, 38, 39, 45, 46, 47, 48, 49, 56, 57, 58, 59, 67, 68, 69, 78, 79, 89)),
                Arguments.of(15, Collections.emptyList())
        );
    }
}