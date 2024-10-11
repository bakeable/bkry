package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/purchase"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PurchaseRepo = repository.NewFirestoreRepository[*purchase.Purchase]()

// GetPurchase retrieves a single Purchase entity by its ID and purchaseID.
func GetPurchase(UserID string, purchaseID string) (purchase.Purchase, error) {
	entityMap, err := PurchaseRepo.Read(purchase.GetDocumentPath(UserID,  purchaseID))
	return purchase.Decode(entityMap), err
}

// GetPurchaseOrNew retrieves a single Purchase entity by its ID and purchaseID.
func GetPurchaseOrNew(UserID string, purchaseID string) (purchase.Purchase, error) {
	entityMap, err := PurchaseRepo.Read(purchase.GetDocumentPath(UserID,  purchaseID))
	if err != nil || entityMap == nil {
		return purchase.Purchase{}, err
	}
	return purchase.Decode(entityMap), err
}

// GetPurchase retrieves a single Purchase entity by its document path.
func GetPurchaseByPath(path string) (purchase.Purchase, error) {
	entityMap, err := PurchaseRepo.Read(path)
	return purchase.Decode(entityMap), err
}

// FindPurchase retrieves a Purchase entity according to given queries.
func FindPurchase(userID string, queries []database.Query) (purchase.Purchase, error) {
	entityMap, err := PurchaseRepo.Find(purchase.GetCollectionPath(userID), queries)
	if err != nil || entityMap == nil {
		return purchase.Purchase{}, err
	}
	return purchase.Decode(entityMap), err
}

// GetAllPurchases retrieves all Purchase entities.
func GetAllPurchases(userID string) ([]purchase.Purchase, error) {
	entityMaps, err := PurchaseRepo.ReadAll(purchase.GetCollectionPath(userID))
	if err != nil {
		return []purchase.Purchase{}, err
	}
	return purchase.DecodeAll(entityMaps), nil
}


// GetAllPurchasesPaginated retrieves all Purchase entities in a paginated manner.
func GetAllPurchasesPaginated(userID string, pagination database.Pagination) ([]purchase.Purchase, database.Pagination, error) {
	entityMaps, pagination, err := PurchaseRepo.ReadPaginated(purchase.GetCollectionPath(userID), pagination)
	if err != nil {
		return []purchase.Purchase{}, pagination, err
	}
	return purchase.DecodeAll(entityMaps), pagination, nil
}

// QueryPurchases retrieves all Purchase entities according to given queries.
func QueryPurchases(userID string, queries []database.Query, pagination database.Pagination) ([]purchase.Purchase, error) {
	entityMaps, err := PurchaseRepo.Query(purchase.GetCollectionPath(userID), queries, pagination)
	if err != nil {
		return []purchase.Purchase{}, err
	}
	return purchase.DecodeAll(entityMaps), nil
}

// QueryPurchasesGroup retrieves all Purchase entities according to given queries.
func QueryPurchasesGroup(queries []database.Query) ([]purchase.Purchase, error) {
	entityMaps, err := PurchaseRepo.QueryGroup("purchases", queries)
	if err != nil {
		return []purchase.Purchase{}, err
	}
	return purchase.DecodeAll(entityMaps), nil
}

// CreatePurchase creates a new Purchase entity.
func CreatePurchase(UserID string, entity purchase.Purchase, editorID *string) (string, error) {
	entity.UserID = UserID
	return PurchaseRepo.Create(&entity, editorID)
}

// UpdatePurchase updates an existing Purchase entity.
func UpdatePurchase(UserID string, entity purchase.Purchase, editorID *string) (string, error) {
	entity.UserID = UserID
	return PurchaseRepo.Update(&entity, editorID)
}

// SavePurchase saves a Purchase entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePurchase(UserID string, entity purchase.Purchase, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePurchase(UserID, entity, editorID)
	} else {
		return UpdatePurchase(UserID, entity, editorID)
	}
}

// DeletePurchase deletes a Purchase entity by its parents' path and purchaseID.
func DeletePurchase(UserID string, purchaseID string) error {
	return PurchaseRepo.Delete(purchase.GetDocumentPath(UserID, purchaseID))
}
