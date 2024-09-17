package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/data/entities/airline_order"
)


func afterQueryGroup(entities []airline_order.AirlineOrder) []airline_order.AirlineOrder {
	return entities
}