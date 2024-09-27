package printing_order_operations

import (
	printing_order "github.com/bakeable/bkry/internal/server/models/entities/printing_order"
)

func afterSave(entity printing_order.PrintingOrder, editorID *string) {}