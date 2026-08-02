package ocrnumbers

import (
	"errors"
	"strings"
)

var digitPatterns = map[string]string{
	" _ | ||_|": "0",
	"     |  |": "1",
	" _  _||_ ": "2",
	" _  _| _|": "3",
	"   |_|  |": "4",
	" _ |_  _|": "5",
	" _ |_ |_|": "6",
	" _   |  |": "7",
	" _ |_||_|": "8",
	" _ |_| _|": "9",
}

func Recognize(s string) ([]string, error) {
	lines := strings.Split(s, "\n")
	if len(lines) > 0 && lines[0] == "" {
		lines = lines[1:]
	}
	if len(lines)%4 != 0 {
		return nil, errors.New("number of input lines is not a multiple of four")
	}

	ans := make([]string, 0, len(lines)/4)
	for i := 0; i < len(lines); i += 4 {
		val, err := processLine(lines[i : i+4])
		if err != nil {
			return nil, err
		}
		ans = append(ans, val)
	}
	return ans, nil
}

func processLine(block []string) (string, error) {
	width := len(block[0])
	if width%3 != 0 {
		return "", errors.New("number of input columns is not a multiple of three")
	}

	var sb strings.Builder
	for i := 0; i < width; i += 3 {
		sb.WriteString(recognizeDigit(block, i, i+3))
	}
	return sb.String(), nil
}

func recognizeDigit(block []string, start, end int) string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString(block[i][start:end])
	}
	if digit, ok := digitPatterns[sb.String()]; ok {
		return digit
	}
	return "?"
}
