package price_layer_operations

import (
	price_layer "github.com/bakeable/bkry/internal/server/models/entities/price_layer"
)

func afterFind(entity price_layer.PriceLayer) price_layer.PriceLayer {
	return entity
}