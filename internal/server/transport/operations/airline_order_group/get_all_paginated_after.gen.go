package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/internal/server/models/entities/airline_order_group"
)


func afterGetAllPaginated(entities []airline_order_group.AirlineOrderGroup, pageSize int, pageNumber int, orderBy string, ascending bool) []airline_order_group.AirlineOrderGroup {
	return entities
}