package camicia

import "fmt"

type Outcome struct {
	finishes bool
	cards    int
	tricks   int
}

var penalties = map[string]int{
	"J": 1,
	"Q": 2,
	"K": 3,
	"A": 4,
}

func SimulateGame(playerA, playerB []string) Outcome {
	hands := [2][]int{getHand(playerA), getHand(playerB)}
	var pile []int
	turn := 0
	owed := 0
	attacker := 0
	cardsPlayed, tricks := 0, 0

	seen := map[string]bool{roundKey(hands, turn): true}
	for {
		if len(hands[turn]) == 0 {
			other := 1 - turn
			hands[other] = append(hands[other], pile...)
			return Outcome{finishes: true, cards: cardsPlayed, tricks: tricks + 1}
		}

		card := hands[turn][0]
		hands[turn] = hands[turn][1:]
		pile = append(pile, card)
		cardsPlayed++

		switch {
		case card > 0:
			attacker = turn
			owed = card
			turn = 1 - turn
		case owed == 0:
			turn = 1 - turn
		default:
			owed--
			if owed > 0 {
				continue
			}

			hands[attacker] = append(hands[attacker], pile...)
			pile = nil
			tricks++
			turn = attacker

			if len(hands[0]) == 0 || len(hands[1]) == 0 {
				return Outcome{finishes: true, cards: cardsPlayed, tricks: tricks}
			}

			key := roundKey(hands, turn)
			if seen[key] {
				return Outcome{finishes: false, cards: cardsPlayed, tricks: tricks}
			}
			seen[key] = true
		}
	}
}

func roundKey(hands [2][]int, turn int) string {
	return fmt.Sprint(hands, turn)
}

func getHand(player []string) []int {
	hand := make([]int, len(player))
	for i, c := range player {
		hand[i] = penalties[c]
	}
	return hand
}
