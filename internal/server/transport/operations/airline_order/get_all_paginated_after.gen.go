package airline_order_operations

import (
	airline_order "github.com/bakeable/bkry/data/entities/airline_order"
)


func afterGetAllPaginated(entities []airline_order.AirlineOrder, pageSize int, pageNumber int, orderBy string, ascending bool) []airline_order.AirlineOrder {
	return entities
}