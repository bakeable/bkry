package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
)

func beforeSave(userID string, entity purchase.Purchase, editorID *string) purchase.Purchase {
	// Return Purchase
	return entity
}