package airline_order_grouping_operations

import (
	airline_order_grouping "github.com/bakeable/bkry/internal/server/models/entities/airline_order_grouping"
)

func afterFind(entity airline_order_grouping.AirlineOrderGrouping) airline_order_grouping.AirlineOrderGrouping {
	return entity
}