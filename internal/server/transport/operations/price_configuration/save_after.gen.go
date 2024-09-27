package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/internal/server/models/entities/price_configuration"
)

func afterSave(entity price_configuration.PriceConfiguration, editorID *string) {}