package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/data/entities/delivery_entry"
)

func beforeSave(entity delivery_entry.DeliveryEntry, editorID *string) delivery_entry.DeliveryEntry {
	// Return DeliveryEntry
	return entity
}