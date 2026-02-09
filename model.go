package main

/*
	Bubble tea app model
	Structs,
	Init,
	Update,
	View
*/

import (
	"fmt"

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
		case " ":
			if len(g.cards) > 0 && !g.gameOver {
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
			for calculateTotal(g.house) < 17 && len(g.cards) > 0 {
				testhand := append(g.house, g.cards[0])
				if checkBust(testhand) {
					g.gameOver = true
					g.playerWon = true
				}
				g.house = append(g.house, g.cards[0])
				g.cards = g.cards[1:]
				if !g.gameOver {
					playerTotal := calculateTotal(g.hand)
					houseTotal := calculateTotal(g.house)
					g.gameOver = true
					g.playerWon = playerTotal > houseTotal
				} else {
					break
				}
			}
			g.turn++
		case "r":
			return initialGame(), tea.ClearScreen
		case "ctrl+c", "q":
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g game) View() string {
	if g.gameOver {
		if !g.playerWon {
			return styles.Render(fmt.Sprintf("You busted with %d! House wins with %d.\nPress 'r' to restart or 'q' to quit.", calculateTotal(g.hand), calculateTotal(g.house)))
		}
		return styles.Render(fmt.Sprintf("You win with %d! House had %d.\nPress 'r' to restart or 'q' to quit.", calculateTotal(g.hand), calculateTotal(g.house)))
	}
	total := calculateTotal(g.hand)
	house := calculateTotal(g.house)
	view := ""
	// total
	view += fmt.Sprintf("Your hand: %d \tHouse: %d\n", total, house)
	for _, card := range g.hand {
		view += fmt.Sprintf("%s %d\n", card.suit, card.value)
	}
	view += "\nPress 'space' to hit, 's' to stand, 'q' to quit."
	return styles.Render(view)
}
