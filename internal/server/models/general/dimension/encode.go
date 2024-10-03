package dimension

// Convert Key struct to database map
func (x Dimension) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"amount": x.Amount,
		"unit": x.Unit,
	}
}