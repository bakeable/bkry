package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/internal/server/models/entities/price_configuration"
)

func afterGet(priceConfigurationID string, entity price_configuration.PriceConfiguration) price_configuration.PriceConfiguration {
	return entity
}