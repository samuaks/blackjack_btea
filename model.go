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

type score struct {
	wins  int
	games int
}

func (s *score) update(playerWon bool) {
	s.games++
	if playerWon {
		s.wins++
	}
}

type game struct {
	cards     []card
	hand      []card
	house     []card
	turn      int
	gameOver  bool
	playerWon bool
	score     score
}

func (g game) startGame(fresh bool) game {
	initialDeck := shuffleDeck(generateDeck())
	initialPlayerHand := []card{initialDeck[0], initialDeck[1]}
	initialHouseHand := []card{initialDeck[2]}
	initialDeck = initialDeck[3:]
	if fresh {
		return game{
			cards: initialDeck,
			hand:  initialPlayerHand,
			house: initialHouseHand,
			turn:  0,
		}
	} else {
		return game{
			cards:     initialDeck,
			hand:      initialPlayerHand,
			house:     initialHouseHand,
			turn:      0,
			score:     g.score,
			playerWon: false,
			gameOver:  false,
		}
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

				g.hand = append(g.hand, g.cards[0])
				g.cards = g.cards[1:]
				if checkBust(g.hand) {
					g.gameOver = true
					g.playerWon = false
					g.score.update(g.playerWon)

				}
			}
			g.turn++
		case "s":
			for calculateTotal(g.house) < 17 && len(g.cards) > 0 {
				g.house = append(g.house, g.cards[0])
				g.cards = g.cards[1:]
				if checkBust(g.house) {
					g.gameOver = true
					g.playerWon = true
					g.score.update(g.playerWon)
					return g, nil
				}
			}
			total := calculateTotal(g.hand)
			house := calculateTotal(g.house)
			g.gameOver = true
			g.playerWon = total > house
			g.score.update(g.playerWon)
			g.turn++
		case "r":
			if g.gameOver {
				return g.startGame(false), tea.ClearScreen
			}
		case "ctrl+c", "q":
			return g, tea.Quit
		}
	}
	return g, nil
}

func (g game) View() string {
	if g.gameOver {
		var gameOverMsg string
		if !g.playerWon {
			gameOverMsg = fmt.Sprintf("🤯 BUST! You busted with %d\n🏠 House wins with %d.", calculateTotal(g.hand), calculateTotal(g.house))
		} else {
			gameOverMsg = fmt.Sprintf("🎉 You win with %d!\n🏠 House had %d.", calculateTotal(g.hand), calculateTotal(g.house))
		}
		//gameOverMsg += fmt.Sprintf("\n\nYour score is %d/%d", g.score.wins, g.score.games)
		gameOverMsg += fmt.Sprintf("\n\nWins: %d \t Win rate: %.2f%%", g.score.wins, float64(g.score.wins)/float64(g.score.games)*100)
		gameOverMsg += "\n\nPress 'r' to play again or 'q' to quit."
		return styles.Render(gameOverMsg)
	}

	total := calculateTotal(g.hand)
	house := calculateTotal(g.house)
	view := ""
	// total
	view += fmt.Sprintf("Your hand: %d \tHouse: %d\n", total, house)
	for _, card := range g.hand {
		view += fmt.Sprintf("%s %d\n", card.suit, card.value)
	}
	//for _, card := range g.house {
	//view += fmt.Sprintf("%s %d\n", card.suit, card.value)
	//}
	view += "\nPress 'space' to hit, 's' to stand, 'q' to quit."
	return styles.Render(view)
}
