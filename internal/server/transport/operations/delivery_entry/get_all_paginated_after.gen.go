package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/internal/server/models/entities/delivery_entry"
)


func afterGetAllPaginated(entities []delivery_entry.DeliveryEntry, pageSize int, pageNumber int, orderBy string, ascending bool) []delivery_entry.DeliveryEntry {
	return entities
}