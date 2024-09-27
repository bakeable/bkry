package airline_pricing_operations

import (
	airline_pricing "github.com/bakeable/bkry/internal/server/models/entities/airline_pricing"
)

func afterGet(airlinePricingID string, entity airline_pricing.AirlinePricing) airline_pricing.AirlinePricing {
	return entity
}