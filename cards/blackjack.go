package cards

import (
	"fmt"

	"github.com/kjblanchard/sgPoker/events"
)

const blackjackEventStart = 50

var (
	playDeck    Deck
	discardDeck Deck
	playerHand  Deck
	dealerHand  Deck
	state       blackjackGamestate
)

type blackjackGamestate int

const (
	blackjackGamestateGameStart = iota
	blackjackGamestateWaiting
	blackjackGamestateTurnStart
	blackjackGamestateTurnWait
	blackjackGamestateTurnEnd
)

const (
	BlackjackEventStartGame = iota + blackjackEventStart
	BlackjackEventStartTurn
	BlackjackEventMax
)

func startBlackjackGame() {
	clear(playDeck)
	clear(discardDeck)
	clear(playerHand)
	clear(dealerHand)
	playDeck.Init()
	playDeck.Shuffle()
	e := events.Event{Type: events.EventTypeStatusMessage, Data: "Welcome to the game!"}
	events.EventBus <- e
}

func startBlackjackTurn() {
	discardDeck.AddDeck(&playerHand, &dealerHand)
	n, cards := playDeck.DrawMultiple(4)
	if n < 4 {
		playDeck.AddDeck(&discardDeck)
		d := Deck{}
		d.AddMultipleCards(cards)
		total := playDeck.CountCards(&d)
		if total != 52 {
			//borked
		}
		_, c := playDeck.DrawMultiple(4 - n)
		cards = append(cards, c...)
	}
	dealerHand.AddMultipleCards(cards[:2])
	playerHand.AddMultipleCards(cards[2:])
	e := events.Event{Type: events.EventTypeStatusMessage, Data: fmt.Sprintf("Dealer drew cards %s and %s", &dealerHand[0], &dealerHand[1])}
	e2 := events.Event{Type: events.EventTypeStatusMessage, Data: fmt.Sprintf("Player drew cards %s and %s", &playerHand[0], &playerHand[1])}
	events.EventBus <- e
	events.EventBus <- e2
}

func handleStateChange(s blackjackGamestate) {
	switch s {
	case blackjackGamestateGameStart:
		startBlackjackGame()
	case blackjackGamestateTurnStart:
		startBlackjackTurn()
	}
	state = s

}

func BlackjackHandleEvent(e *events.Event) {
	if e.Type < BlackjackEventStartGame || e.Type > BlackjackEventMax {
		return
	}
	switch e.Type {
	case BlackjackEventStartGame:
		handleStateChange(blackjackGamestateGameStart)
	case BlackjackEventStartTurn:
		handleStateChange(blackjackGamestateTurnStart)
	}
}
