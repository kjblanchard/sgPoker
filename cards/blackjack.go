package cards

import (
	"fmt"
	"strconv"
	"strings"

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

func getCardValue(c Card) int {
	if c.Rank <= Rank10 {
		return int(c.Rank) + 1
	} else if c.Rank < RankA {
		return 10
	}
	return 11
}

func getHandValues(d *Deck) []int {
	sum := 0
	aces := 0
	for _, c := range *d {
		v := getCardValue(c)
		if v == 11 {
			aces++
		}
		sum += v
	}
	values := []int{sum}
	//For each ace, you can have two values, so account for each ace being both values
	//Seems like for each ace, we can subtract total by 10 based on my back of napkin math
	for i := 0; i < aces; i++ {
		values = append(values, sum-(10*(i+1)))
	}
	return values
}

func printHandValues(v []int) string {
	var vs strings.Builder
	for i, n := range v {
		vs.WriteString(strconv.Itoa(n))
		if i < len(v)-1 {
			vs.WriteString(" or ")
		}
	}
	return vs.String()
}

func startBlackjackGame() {
	clear(playDeck)
	clear(discardDeck)
	clear(playerHand)
	clear(dealerHand)
	playDeck.Init()
	playDeck.Shuffle()
	e := events.Event{Type: events.EventTypeStatusMessage, Data: "Welcome to the game! Send 'deal' to start"}
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
	v := getHandValues(&playerHand)
	e2 := events.Event{Type: events.EventTypeStatusMessage, Data: fmt.Sprintf("Player drew cards %s and %s for a total of %s", &playerHand[0], &playerHand[1], printHandValues(v))}
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
