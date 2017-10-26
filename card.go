package deck

import "fmt"

// Card is a playing card
type Card struct {
	Suit
	Value
}

// Suit is the suit of a card
type Suit string

const (
	// Hearts is the hearts suit
	Hearts Suit = "hearts"
	// Spades is the spades suit
	Spades Suit = "spades"
	// Diamonds is the diamonds suit
	Diamonds Suit = "diamonds"
	// Clubs is the clubs suit
	Clubs Suit = "clubs"
)

var suits = [4]Suit{Hearts, Spades, Diamonds, Clubs}

// Value is the value of a card
type Value int

const (
	// Ace is an ace
	Ace Value = iota
	// Two is a 2
	Two
	// Three is a three
	Three
	// Four is a four
	Four
	// Five is a five
	Five
	// Six is a six
	Six
	// Seven is a seven
	Seven
	// Eight is an eight
	Eight
	// Nine is a nine
	Nine
	// Ten is a ten
	Ten
	// Jack is a jack
	Jack
	// Queen is a queen
	Queen
	// King is a king
	King
)

var values = [13]Value{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// NewCard creates a new card of the given suit
// if the suit and value are not valid, a panic occurs
func NewCard(s Suit, v Value) *Card {
	validate(s, v)
	return &Card{
		Suit:  s,
		Value: v,
	}
}

func validate(s Suit, v Value) {
	if s != Hearts && s != Spades && s != Diamonds && s != Clubs {
		panic(fmt.Sprintf("Unknown suit %s", s))
	}
	if v > King || v < Ace {
		panic(fmt.Sprintf("Unknown value %d", v))
	}
}
