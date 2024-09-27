package delivery_entry_operations

import (
	delivery_entry "github.com/bakeable/bkry/internal/server/models/entities/delivery_entry"
)

func afterSave(entity delivery_entry.DeliveryEntry, editorID *string) {}