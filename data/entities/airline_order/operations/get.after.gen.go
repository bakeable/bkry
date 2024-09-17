package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/data/entities/airline_order"
)

func afterGet(airlineOrderID string, entity airline_order.AirlineOrder) airline_order.AirlineOrder {
	return entity
}