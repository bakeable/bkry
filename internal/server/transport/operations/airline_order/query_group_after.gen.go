package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/internal/server/models/entities/airline_order"
)


func afterQueryGroup(entities []airline_order.AirlineOrder) []airline_order.AirlineOrder {
	return entities
}