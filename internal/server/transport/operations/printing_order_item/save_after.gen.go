package printing_order_item_operations

import (
	printing_order_item "github.com/bakeable/bkry/internal/server/models/entities/printing_order_item"
)

func afterSave(entity printing_order_item.PrintingOrderItem, editorID *string) {}