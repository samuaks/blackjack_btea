package main

/*
	Mostly game logic helpers here
*/
import "math/rand"

func generateDeck() []card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} // Ace is represented as 11
	deck := []card{}

	for _, suit := range suits {
		for _, value := range values {
			isAce := value == 11
			deck = append(deck, card{suit, value, isAce})
		}
	}
	return deck
}

func shuffleDeck(deck []card) []card {
	shuffled := make([]card, len(deck))
	perm := rand.Perm(len(deck))
	for i, v := range perm {
		shuffled[i] = deck[v]
	}
	return shuffled
}

func checkBust(hand []card) bool {
	return calculateTotal(hand) > 21
}

func calculateTotal(hand []card) int {
	total := 0
	aces := 0
	for _, card := range hand {
		total += card.value
		if card.isAce {
			aces++
		}
	}

	for total > 21 && aces > 0 {
		total -= 10
		aces--
	}
	return total
}
