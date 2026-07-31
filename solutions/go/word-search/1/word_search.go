package wordsearch

import "errors"

var directions = [8][2]int{
	{0, 1}, {0, -1},
	{1, 0}, {-1, 0},
	{1, 1}, {1, -1},
	{-1, 1}, {-1, -1},
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	out := make(map[string][2][2]int)
	nr := len(puzzle)
	nc := 0
	if nr > 0 {
		nc = len(puzzle[0])
	}

	for _, w := range words {
		found := false
	search:
		for r := 0; r < nr; r++ {
			for c := 0; c < nc; c++ {
				for _, d := range directions {
					if endR, endC, ok := match(puzzle, w, r, c, d[0], d[1]); ok {
						out[w] = [2][2]int{{c, r}, {endC, endR}}
						found = true
						break search
					}
				}
			}
		}
		if !found {
			return nil, errors.New("unsolved")
		}
	}

	return out, nil
}

func match(puzzle []string, w string, r, c, dr, dc int) (endR, endC int, ok bool) {
	nr, nc := len(puzzle), len(puzzle[0])
	row, col := r, c
	for i := 0; i < len(w); i++ {
		if row < 0 || row >= nr || col < 0 || col >= nc || puzzle[row][col] != w[i] {
			return 0, 0, false
		}
		endR, endC = row, col
		row += dr
		col += dc
	}
	return endR, endC, true
}
