package deck

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck is a standard deck of cards
type Deck struct {
	cards []*Card
	index int
}

// New creates a new deck
func New() *Deck {
	return &Deck{
		cards: getCards(),
		index: 0,
	}
}

func getCards() []*Card {
	var cards []*Card
	k := 0
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			cards[k] = NewCard(suits[i], values[j])
			k++
		}
	}
	return cards
}

// Size returns the number of remaining cards in the deck
func (d *Deck) Size() int {
	return len(d.cards) - d.index
}

// IsEmpty returns if the deck is empty
func (d *Deck) IsEmpty() bool {
	return d.Size() == 0
}

// Shuffle shuffles the deck
func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffled := make([]*Card, d.Size())
	for i, j := range r.Perm(d.Size()) {
		shuffled[i] = d.cards[d.index+j]
	}
	for i := 0; i < len(shuffled); i++ {
		d.cards[d.index+1] = shuffled[i]
	}
}

// Reset resets the deck by putting all the card back that were dealt
// shuffle must be called after or the deck will be dealt in the same order
func (d *Deck) Reset() {
	d.index = 0
}

// ResetAndShuffle resets and shuffles the deck for a new use
func (d *Deck) ResetAndShuffle() {
	d.Reset()
	d.Shuffle()
}

// Draw draws the next card from the deck
func (d *Deck) Draw() *Card {
	if d.index < len(d.cards) {
		i := d.index
		d.index++
		return d.cards[i]
	}
	return nil
}

// Deal deals out the cards into numHands of handSize
func (d *Deck) Deal(numHands, handSize int) ([][]*Card, error) {
	if numHands*handSize > d.Size() {
		return nil, fmt.Errorf("Not enough card ")
	}
	hands := make([][]*Card, numHands)
	for i := 0; i < numHands; i++ {
		hands[i] = make([]*Card, handSize)
		for j := 0; j < handSize; j++ {
			hands[i][j] = d.Draw()
		}
	}
	return hands, nil
}

// DealInto deals cards into the hands
func (d *Deck) DealInto(hands [][]*Card) {
	for i := 0; i < len(hands); i++ {
		for j := 0; j < len(hands[i]); j++ {
			c := d.Draw()
			if c == nil {
				return
			}
			hands[i][j] = c
		}
	}
	return
}

// DealAll deals all of the cards into n piles
func (d *Deck) DealAll(n int) [][]*Card {
	if n < 1 {
		return nil
	}
	if n > d.Size() {
		n = d.Size()
	}
	hands := make([][]*Card, n)
	for i := 0; i < len(hands); i++ {
		hands[i] = make([]*Card, 0, d.Size()/n+1)
	}
	i := 0
	for !d.IsEmpty() {
		hands[i] = append(hands[i], d.Draw())
		i = (i + 1) % len(hands)
	}
	return hands
}
