package price_layer_operations

import (
	price_layer "github.com/bakeable/bkry/data/entities/price_layer"
)


func afterGetAllPaginated(entities []price_layer.PriceLayer, pageSize int, pageNumber int, orderBy string, ascending bool) []price_layer.PriceLayer {
	return entities
}