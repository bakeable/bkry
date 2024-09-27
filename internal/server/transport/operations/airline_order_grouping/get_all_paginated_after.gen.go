package airline_order_grouping_operations

import (
	airline_order_grouping "github.com/bakeable/bkry/internal/server/models/entities/airline_order_grouping"
)


func afterGetAllPaginated(entities []airline_order_grouping.AirlineOrderGrouping, pageSize int, pageNumber int, orderBy string, ascending bool) []airline_order_grouping.AirlineOrderGrouping {
	return entities
}