package deck

// Deck is a standard deck of cards
type Deck struct {
	Cards [52]*Card
}

// NewDeck creates a new deck
func NewDeck() *Deck {
	return &Deck{
		Cards: getCards(),
	}
}

func getCards() [52]*Card {
	var cards [52]*Card
	k := 0
	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			cards[k] = NewCard(suits[i], values[j])
			k++
		}
	}
	return cards
}
