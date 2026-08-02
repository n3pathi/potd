package killersudokuhelper

import "slices"

func Combinations(sum, size int, exclude []int) [][]int {
	output := make([][]int, 0, 1)
	digits := getDigits(exclude)
	
	var search func(start, remaining int, row []int)
	search = func(start, remaining int, row []int) {
		if len(row) == size {
			if remaining == 0 {
				output = append(output, slices.Clone(row))
			}
			return
		}
		for i := start; i < len(digits); i++ {
			d := digits[i]
			if d > remaining {
				break
			}
			search(i+1, remaining-d, append(row, d))
		}
	}
	
	search(0, sum, []int{})
	
	return output
}

func getDigits(exclude []int) []int {
	excluded := make(map[int]bool)
	for _, v := range exclude {
		excluded[v] = true
	}
	set := make([]int, 0, 1)
	for i := 1; i <= 9; i++ {
		if excluded[i] {
			continue
		}
		set = append(set, i)
	}
	return set
}
