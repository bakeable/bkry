package price_configuration_operations

import (
	price_configuration "github.com/bakeable/bkry/data/entities/price_configuration"
)


func afterGetAllPaginated(entities []price_configuration.PriceConfiguration, pageSize int, pageNumber int, orderBy string, ascending bool) []price_configuration.PriceConfiguration {
	return entities
}