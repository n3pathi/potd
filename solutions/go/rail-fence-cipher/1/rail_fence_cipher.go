package railfencecipher

import (
	"strings"
)

func Encode(message string, rails int) string {
	if message == "" || rails < 1 {
		return message
	}
	runes := []rune(message)
	n := len(runes)
	matrix := make([][]rune, 0, rails)
	for range rails {
		row := make([]rune, 0, n)
		for range n {
			row = append(row, ' ')
		}
		matrix = append(matrix, row)
	}

	row, col, dirDown := 0, 0, true
	for i := range n {
		if row < 0 {
			dirDown = true
			row = 1
		} else if row >= rails {
			dirDown = false
			row = rails - 2
		}
		matrix[row][col] = runes[i]
		col++
		if dirDown {
			row++
		} else {
			row--
		}
	}

	var sb strings.Builder
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != ' ' {
				sb.WriteRune(matrix[i][j])
			}
		}
	}

	return sb.String()
}

func Decode(message string, rails int) string {
	runes := []rune(message)
	n := len(runes)

	matrix := make([][]rune, 0, rails)

	for i := 0; i < rails; i++ {
		row := make([]rune, 0, n)
		for j := 0; j < n; j++ {
			row = append(row, ' ')
		}
		matrix = append(matrix, row)
	}

	dirDown := true
	for i, r := 0, 0; i < n; i++ {
		if r < 0 {
			dirDown = true
			r = 1
		} else if r >= rails {
			dirDown = false
			r = rails - 2
		}
		matrix[r][i] = '?'
		if dirDown {
			r++
		} else {
			r--
		}
	}

	k := 0
	for i := 0; i < rails; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '?' {
				matrix[i][j] = runes[k]
				k++
			}
		}
	}
	var sb strings.Builder
	r := 0
	dirDown = true
	for j := 0; j < n; j++ {
		if r < 0 {
			dirDown = true
			r = 1
		} else if r >= rails {
			dirDown = false
			r = rails - 2
		}
		sb.WriteRune(matrix[r][j])
		if dirDown {
			r++
		} else {
			r--
		}
	}
	return sb.String()
}
