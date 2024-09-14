package margin

// Convert Key struct to database map
func (x Margin) Encode() map[string]interface{} {
	return map[string]interface{}{
		"amount": x.Amount,
		"percentage": x.Percentage,
		"percentageType": x.PercentageType,
		"staggered": x.Staggered,
		"type": x.Type,
		"values": x.Values,
	}
}