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
