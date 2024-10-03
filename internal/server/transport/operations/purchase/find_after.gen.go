package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
)

func afterFind(userID string, entity purchase.Purchase) purchase.Purchase {
	return entity
}