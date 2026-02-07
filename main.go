package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type card struct {
	suit  string
	value int
	isAce bool
}

type game struct {
	cards     []card
	hand      []card
	house     []card
	turn      int
	gameOver  bool
	playerWon bool
}

func initialGame() game {
	initialDeck := shuffleDeck(generateDeck())
	initialPlayerHand := []card{initialDeck[0], initialDeck[1]}
	initialHouseHand := []card{initialDeck[2]}
	initialDeck = initialDeck[3:]
	return game{
		cards: initialDeck,
		hand:  initialPlayerHand,
		house: initialHouseHand,
		turn:  0,
	}
}

func (g game) Init() tea.Cmd {
	return nil
}

func (g game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			if len(g.cards) > 0 {
				// pre check for bust
				testHand := append(g.hand, g.cards[0])
				if checkBust(testHand) {
					g.gameOver = true
					g.playerWon = false
				}
				g.hand = append(g.hand, g.cards[0])
				g.cards = g.cards[1:]
			}
			g.turn++
		case "s":
			if len(g.cards) > 0 {
				testHand := append(g.house, g.cards[0])
				if checkBust(testHand) {
					g.gameOver = true
					g.playerWon = true
				}
				if calculateTotal(g.house) < 17 {
					g.house = append(g.house, g.cards[0])
					g.cards = g.cards[1:]
				}
			}
			g.turn++
		case "ctrl+c", "q":
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g game) View() string {
	if g.gameOver {
		if !g.playerWon {
			return fmt.Sprintf("You busted with %d! House wins.\nPress 'q' to quit.", calculateTotal(g.hand))
		}
		return fmt.Sprintf("You win with %d! House had %d.\nPress 'q' to quit.", calculateTotal(g.hand), calculateTotal(g.house))
	}
	total := calculateTotal(g.hand)
	house := calculateTotal(g.house)
	view := ""
	// total
	view += fmt.Sprintf("Your hand: %d \tHouse: %d\n", total, house)
	for _, card := range g.hand {
		view += fmt.Sprintf("%s %d\n", card.suit, card.value)
	}
	view += "\nPress 'h' to hit, 's' to stand, 'q' to quit."
	return view
}

func main() {
	if _, err := tea.NewProgram(initialGame()).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
