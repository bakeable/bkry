package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/internal/server/models/entities/airline_order"
)

func afterSave(entity airline_order.AirlineOrder, editorID *string) {}