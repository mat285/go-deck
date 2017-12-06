package deck

import "fmt"

// Card is a playing card
type Card struct {
	Suit
	Value
}

// Suit is the suit of a card
type Suit string

// Color is the color of the card
type Color string

const (
	// Red is the color red
	Red Color = "Red"
	// Black is the color black
	Black Color = "Black"
)

const (
	// Hearts is the hearts suit
	Hearts Suit = "Hearts"
	// Spades is the spades suit
	Spades Suit = "Spades"
	// Diamonds is the diamonds suit
	Diamonds Suit = "Diamonds"
	// Clubs is the clubs suit
	Clubs Suit = "Clubs"
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

func valueString(v Value) string {
	switch v {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return ""
	}
}

// String returns a string representation of this card
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", valueString(c.Value), c.Suit)
}

// Color returns the color of the card
func (c Card) Color() Color {
	if c.Suit == Spades || c.Suit == Clubs {
		return Black
	}
	return Red
}

// IsBlack returns if the card is black
func (c Card) IsBlack() bool {
	return c.Color() == Black
}

// IsRed returns if the card is red
func (c Card) IsRed() bool {
	return c.Color() == Red
}

// Compare returns negative is c < o, 0 if c = o and positive of c > o
func (c Card) Compare(o Card, acesLow ...bool) int {
	if o.Value == c.Value {
		return 0
	}
	aceHigh := true
	if len(acesLow) > 0 {
		aceHigh = acesLow[0]
	}
	if (c.Value == Ace && aceHigh) || (o.Value == Ace && !aceHigh) || c.Value > o.Value {
		return 1
	}
	return -1
}

// Equals returns if c is the same card as o
func (c Card) Equals(o Card) bool {
	return c.Value == o.Value && c.Suit == o.Suit
}
