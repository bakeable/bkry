package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/data/entities/airline_order"
)

func afterFind(entity airline_order.AirlineOrder) airline_order.AirlineOrder {
	return entity
}