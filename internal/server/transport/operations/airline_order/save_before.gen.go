package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/data/entities/airline_order"
)

func beforeSave(entity airline_order.AirlineOrder, editorID *string) airline_order.AirlineOrder {
	// Return AirlineOrder
	return entity
}