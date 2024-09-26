package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/data/entities/airline_order_group"
)

func afterFind(entity airline_order_group.AirlineOrderGroup) airline_order_group.AirlineOrderGroup {
	return entity
}