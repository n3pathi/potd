package g4g;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;

class MaxPeopleDefeatedTest {
    private final MaxPeopleDefeated sut = new MaxPeopleDefeated();

    @ParameterizedTest
    @MethodSource("dataProvider")
    void solve(int p, int expected) {
        assertEquals(expected, sut.solve(p));
    }

    private static Stream<Arguments> dataProvider() {
        return Stream.of(
                Arguments.of(14, 3),
                Arguments.of(10, 2)
        );
    }
}