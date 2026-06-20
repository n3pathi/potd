package g4g;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrowsExactly;

class LastDigitTest {
    private final LastDigit lastDigit = new LastDigit();

    @ParameterizedTest
    @MethodSource("dataProvider")
    void testLastDigit(String a, String b, int expected) {
        assertEquals(expected, lastDigit.getLastDigit(a, b));
    }

    static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of("3", "10", 9),
                Arguments.of("6", "2", 6),
                Arguments.of("2", "4", 6),
                Arguments.of("10", "1", 0)
        );
    }

    @Test
    void testLastDigitWithZero() {
        assertThrowsExactly(ArithmeticException.class, () -> lastDigit.getLastDigit("0", "0"));
    }

    @Test
    void testWithNegativeExponent() {
        assertThrowsExactly(IllegalArgumentException.class, () -> lastDigit.getLastDigit("-1", "-1"));
    }
}