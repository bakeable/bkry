package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/data/entities/price_configuration"
)

func beforeSave(entity price_configuration.PriceConfiguration, editorID *string) price_configuration.PriceConfiguration {
	// Return PriceConfiguration
	return entity
}