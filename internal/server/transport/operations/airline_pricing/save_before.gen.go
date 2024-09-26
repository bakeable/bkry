package airline_pricing_operations

import (
	airline_pricing "github.com/bakeable/bkry/data/entities/airline_pricing"
)

func beforeSave(entity airline_pricing.AirlinePricing, editorID *string) airline_pricing.AirlinePricing {
	// Return AirlinePricing
	return entity
}