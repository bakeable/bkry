package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
)

func afterGet(userID string, purchaseID string, entity purchase.Purchase) purchase.Purchase {
	return entity
}