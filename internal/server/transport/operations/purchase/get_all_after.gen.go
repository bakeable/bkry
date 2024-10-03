package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
)


func afterGetAll(userID string, entities []purchase.Purchase) []purchase.Purchase {
	return entities
}