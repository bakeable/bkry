package purchase_operations

import (
	purchase "github.com/bakeable/bkry/internal/server/models/entities/purchase"
)


func afterGetAllPaginated(userID string, entities []purchase.Purchase, pageSize int, pageNumber int, orderBy string, ascending bool) []purchase.Purchase {
	return entities
}