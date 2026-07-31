package wordy

import (
	"strconv"
	"strings"
)

var operators = map[string]byte{
	"plus":       '+',
	"minus":      '-',
	"multiplied": '*',
	"divided":    '/',
}

// term is either a number (isOp == false) or an operator (isOp == true),
// kept in the order it appeared in the question.
type term struct {
	isOp bool
	op   byte
	num  int
}

func Answer(question string) (int, bool) {
	s := strings.TrimSuffix(strings.TrimSpace(question), "?")
	s = strings.TrimPrefix(s, "What is ")
	s = strings.ReplaceAll(s, "multiplied by", "multiplied")
	s = strings.ReplaceAll(s, "divided by", "divided")

	var terms []term
	for t := range strings.FieldsSeq(s) {
		if op, isOp := operators[t]; isOp {
			terms = append(terms, term{isOp: true, op: op})
			continue
		}
		val, err := strconv.Atoi(t)
		if err != nil {
			return 0, false
		}
		terms = append(terms, term{num: val})
	}

	return evaluate(terms)
}

// evaluate walks terms expecting number, op, number, op, ..., number.
// Any deviation from that pattern - two numbers in a row, two operators
// in a row, a leading operator, or a trailing operator - is an error.
func evaluate(terms []term) (int, bool) {
	if len(terms) == 0 || terms[0].isOp {
		return 0, false
	}
	result := terms[0].num

	for i := 1; i < len(terms); i += 2 {
		if !terms[i].isOp || i+1 >= len(terms) || terms[i+1].isOp {
			return 0, false
		}
		operand := terms[i+1].num
		switch terms[i].op {
		case '+':
			result += operand
		case '-':
			result -= operand
		case '*':
			result *= operand
		case '/':
			if operand == 0 {
				return 0, false
			}
			result /= operand
		}
	}
	return result, true
}
