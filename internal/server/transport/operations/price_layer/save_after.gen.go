package price_layer_operations

import (
	price_layer "github.com/bakeable/bkry/internal/server/models/entities/price_layer"
)

func afterSave(entity price_layer.PriceLayer, editorID *string) {}