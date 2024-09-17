package price_layer_operations

import (
	price_layer "github.com/bakeable/bkry/data/entities/price_layer"
)

func afterGet(priceLayerID string, entity price_layer.PriceLayer) price_layer.PriceLayer {
	return entity
}