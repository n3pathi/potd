package yacht

func Score(dice []int, category string) int {
	for _, d := range dice {
		if d < 0 || d > 6 {
			return 0
		}
	}

	n := len(dice)
	if n != 5 {
		return 0
	}

	switch category {
	case "ones":
		return sum(dice, 1)
	case "twos":
		return sum(dice, 2)
	case "threes":
		return sum(dice, 3)
	case "fours":
		return sum(dice, 4)
	case "fives":
		return sum(dice, 5)
	case "sixes":
		return sum(dice, 6)
	case "full house":
		f := getFreq(dice)
		if len(f) != 2 {
			return 0
		}
		score := 0
		for k, v := range f {
			switch v{
			case 3:
				score+=k*v
			case 2:
				score+=k*v
			default:
				return 0
			}
		}
		return score
	case "four of a kind":
		f := getFreq(dice)
		for k, v := range f {
			if v >= 4 {
				return k * 4
			}
		}
		return 0
	case "little straight":
		f := getFreq(dice)
		if len(f) == 5 && f[1] == 1 && f[6] == 0 {
			return 30
		}
	case "big straight":
		f := getFreq(dice)
		if len(f) == 5 && f[6] == 1 && f[1] == 0 {
			return 30
		}
	case "choice":
		sum := 0
		for _, d := range dice {
			sum += d
		}
		return sum
	case "yacht":
		for i := 1; i < n; i++ {
			if dice[i] != dice[i-1] {
				return 0
			}
		}
		return 50
	}
	return 0
}

func getFreq(dice []int) map[int]int {
	f := make(map[int]int)
	for _, d := range dice {
		f[d]++
	}
	return f
}

func sum(dice []int, face int) int {
	sum := 0
	for _, d := range dice {
		if d == face {
			sum+=d
		}
	}
	return sum
}
