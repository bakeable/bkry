package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/internal/server/models/entities/airline_order_group"
)

func afterGet(airlineOrderGroupID string, entity airline_order_group.AirlineOrderGroup) airline_order_group.AirlineOrderGroup {
	return entity
}