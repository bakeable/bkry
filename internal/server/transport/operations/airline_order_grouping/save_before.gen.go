package airline_order_grouping_operations

import (
	airline_order_grouping "github.com/bakeable/bkry/internal/server/models/entities/airline_order_grouping"
)

func beforeSave(entity airline_order_grouping.AirlineOrderGrouping, editorID *string) airline_order_grouping.AirlineOrderGrouping {
	// Return AirlineOrderGrouping
	return entity
}