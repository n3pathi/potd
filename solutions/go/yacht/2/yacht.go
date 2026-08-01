package yacht

func Score(dice []int, category string) int {
	freq := make(map[int]int)
	for _, d := range dice {
		freq[d]++
	}

	switch category {
	case "ones":
		return freq[1] * 1
	case "twos":
		return freq[2] * 2
	case "threes":
		return freq[3] * 3
	case "fours":
		return freq[4] * 4
	case "fives":
		return freq[5] * 5
	case "sixes":
		return freq[6] * 6
	case "full house":
		if len(freq) != 2 {
			return 0
		}
		for _, v := range freq {
			if v != 2 && v != 3 {
				return 0
			}
		}
		return total(dice)
	case "four of a kind":
		for face, count := range freq {
			if count >= 4 {
				return face * 4
			}
		}
		return 0
	case "little straight":
		if len(freq) == 5 && freq[6] == 0 {
			return 30
		}
		return 0
	case "big straight":
		if len(freq) == 5 && freq[1] == 0 {
			return 30
		}
		return 0
	case "choice":
		return total(dice)
	case "yacht":
		if len(freq) == 1 {
			return 50
		}
		return 0
	}
	return 0
}

func total(dice []int) int {
	sum := 0
	for _, d := range dice {
		sum += d
	}
	return sum
}
