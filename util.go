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
