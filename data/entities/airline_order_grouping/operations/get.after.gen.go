package airline_order_grouping_operations

import (
	airline_order_grouping "github.com/bakeable/bkry/data/entities/airline_order_grouping"
)

func afterGet(airlineOrderGroupingID string, entity airline_order_grouping.AirlineOrderGrouping) airline_order_grouping.AirlineOrderGrouping {
	return entity
}