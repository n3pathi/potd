package lc;

import java.util.Arrays;

public class NumStrings {
    static int solve(String[] patterns, String word){
        return (int) Arrays.stream(patterns).filter(word::contains).count();
    }
}
