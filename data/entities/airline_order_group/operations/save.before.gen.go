package airline_order_group_operations

import (
	airline_order_group "github.com/bakeable/bkry/data/entities/airline_order_group"
)

func beforeSave(entity airline_order_group.AirlineOrderGroup, editorID *string) airline_order_group.AirlineOrderGroup {
	// Return AirlineOrderGroup
	return entity
}