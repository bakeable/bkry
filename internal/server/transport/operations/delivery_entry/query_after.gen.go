package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/data/entities/delivery_entry"
)


func afterQuery(entities []delivery_entry.DeliveryEntry) []delivery_entry.DeliveryEntry {
	return entities
}