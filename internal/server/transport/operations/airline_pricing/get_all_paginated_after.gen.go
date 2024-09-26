package airline_pricing_operations

import (
	airline_pricing "github.com/bakeable/bkry/data/entities/airline_pricing"
)


func afterGetAllPaginated(entities []airline_pricing.AirlinePricing, pageSize int, pageNumber int, orderBy string, ascending bool) []airline_pricing.AirlinePricing {
	return entities
}