package deck

// Contains checks if the slice contains the card
func Contains(cards []*Card, card *Card) bool {
	if card == nil {
		return false
	}
	for _, c := range cards {
		if c.Equals(*card) {
			return true
		}
	}
	return false
}

// ContainsValue returns if the slice contains a card with the value
func ContainsValue(cards []*Card, card *Card) bool {
	if card == nil {
		return false
	}
	for _, c := range cards {
		if c.Value == card.Value {
			return true
		}
	}
	return false
}

// ContainsSuit returns if the slice contains the suit
func ContainsSuit(cards []*Card, suit Suit) bool {
	for _, c := range cards {
		if c.Suit == suit {
			return true
		}
	}
	return false
}

// SetMinus subtracts sub from cards
func SetMinus(cards, sub []*Card) []*Card {
	subMap := map[string]bool{}
	for _, c := range sub {
		if c != nil {
			subMap[c.String()] = true
		}
	}
	ret := []*Card{}
	for _, c := range cards {
		if c == nil {
			continue
		}
		if _, ok := subMap[c.String()]; !ok {
			ret = append(ret, c)
		}
	}
	return ret
}

// Max returns the maximized card from the given function and its index
func Max(cards []*Card, compare func(i, j *Card) int) (*Card, int) {
	if len(cards) == 0 {
		return nil, -1
	}
	max := cards[0]
	idx := 0
	for i := 0; i < len(cards); i++ {
		if compare(cards[i], max) > 0 {
			max = cards[i]
			idx = i
		}
	}
	return max, idx
}

// IsRed returns if the suit is red
func IsRed(s Suit) bool {
	return s == Hearts || s == Diamonds
}

// IsBlack returns if the suit is black
func IsBlack(s Suit) bool {
	return s == Spades || s == Clubs
}

// SameColor returns if the two suits are the same color
func SameColor(s1, s2 Suit) bool {
	return (IsBlack(s1) && IsBlack(s2)) || (IsRed(s1) && IsRed(s2))
}
