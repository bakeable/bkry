package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
)

func afterSave(userID string, entity purchase.Purchase, editorID *string) {}