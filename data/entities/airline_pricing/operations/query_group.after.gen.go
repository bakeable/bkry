package airline_pricing_operations

import (
	airline_pricing "github.com/bakeable/bkry/data/entities/airline_pricing"
)


func afterQueryGroup(entities []airline_pricing.AirlinePricing) []airline_pricing.AirlinePricing {
	return entities
}