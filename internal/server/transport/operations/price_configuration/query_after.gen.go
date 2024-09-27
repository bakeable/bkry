package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/internal/server/models/entities/price_configuration"
)


func afterQuery(entities []price_configuration.PriceConfiguration) []price_configuration.PriceConfiguration {
	return entities
}