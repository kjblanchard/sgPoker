package cards

import (
	"fmt"
	"math/rand"
)

type Rank int

const (
	Rank1 = iota
	Rank2
	Rank3
	Rank4
	Rank5
	Rank6
	Rank7
	Rank8
	Rank9
	Rank10
	RankJ
	RankQ
	RankK
	RankA
	RankCount
)

func (r Rank) String() string {
	switch r {
	case Rank1:
		return "1"
	case Rank2:
		return "2"
	case Rank3:
		return "3"
	case Rank4:
		return "4"
	case Rank5:
		return "5"
	case Rank6:
		return "6"
	case Rank7:
		return "7"
	case Rank8:
		return "8"
	case Rank9:
		return "9"
	case Rank10:
		return "10"
	case RankJ:
		return "Jack"
	case RankQ:
		return "Queen"
	case RankK:
		return "King"
	case RankA:
		return "Ace"
	}
	return "unsupported rank"
}

type Suit int

const (
	SuitHearts = iota
	SuitDiamonds
	SuitSpades
	SuitClubs
	SuitCount
)

func (s Suit) String() string {
	switch s {
	case SuitClubs:
		return "Clubs"
	case SuitDiamonds:
		return "Diamonds"
	case SuitHearts:
		return "Hearts"
	case SuitSpades:
		return "Spades"
	}
	return "unsupported suit"
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (c *Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

type Deck []Card

// Creates a deck with 52 cards, not shuffled
func (d *Deck) Init() {
	clear(*d)
	(*d) = make(Deck, 0, RankCount*SuitCount)
	for r := range RankCount {
		for s := range SuitCount {
			*d = append(*d, Card{Rank(r), Suit(s)})
		}
	}
}

func (d *Deck) Draw() (int, Card) {
	if len(*d) == 0 {
		return 0, Card{0, 0}
	}
	c := (*d)[0]
	*d = (*d)[1:]
	return 1, c
}

func (d *Deck) DrawMultiple(n int) (int, []Card) {
	if n <= 0 {
		return 0, nil
	}
	max := len(*d)
	if n > max {
		n = max
	}
	c := (*d)[:n]
	*d = (*d)[n:]
	return n, c
}

func (d *Deck) AddCard(c ...Card) {
	for _, v := range c {
		*d = append(*d, v)
	}
}

func (d *Deck) AddMultipleCards(c []Card) {
	*d = append(*d, c...)
}

// Adds decks from d2 to d, removing all from d2
func (d *Deck) AddDeck(d2 ...*Deck) {
	for _, v := range d2 {
		*d = append(*d, (*v)...)
		clear(*v)
	}
}

func (d *Deck) CountCards(d2 ...*Deck) int {
	c := len(*d)
	for _, v := range d2 {
		c += len(*v)
	}
	return c
}

// Shuffles remaining cards in the deck
func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]

	})

}
