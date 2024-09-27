package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/internal/server/models/entities/delivery_entry"
)


func afterGetAll(entities []delivery_entry.DeliveryEntry) []delivery_entry.DeliveryEntry {
	return entities
}