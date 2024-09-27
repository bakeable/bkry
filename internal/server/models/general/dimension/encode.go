package dimension

// Convert Key struct to database map
func (x Dimension) Encode() map[string]interface{} {
	return map[string]interface{}{
		"amount": x.Amount,
		"unit": x.Unit,
	}
}