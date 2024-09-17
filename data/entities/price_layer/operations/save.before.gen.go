package price_layer_operations

import (
	price_layer "github.com/bakeable/bkry/data/entities/price_layer"
)

func beforeSave(entity price_layer.PriceLayer, editorID *string) price_layer.PriceLayer {
	// Return PriceLayer
	return entity
}