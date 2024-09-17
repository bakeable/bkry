package repo

import (
	"github.com/bakeable/bkry/data/entities/packing_slip"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PackingSlipRepo = repository.NewRepository[*packing_slip.PackingSlip]()

// GetPackingSlip retrieves a single PackingSlip entity by its ID and packingSlipID.
func GetPackingSlip(packingSlipID string) (packing_slip.PackingSlip, error) {
	entityMap, err := PackingSlipRepo.Read(packing_slip.GetDocumentPath( packingSlipID))
	return packing_slip.Decode(entityMap), err
}

// GetPackingSlipOrNew retrieves a single PackingSlip entity by its ID and packingSlipID.
func GetPackingSlipOrNew(packingSlipID string) (packing_slip.PackingSlip, error) {
	entityMap, err := PackingSlipRepo.Read(packing_slip.GetDocumentPath( packingSlipID))
	if err != nil || entityMap == nil {
		return packing_slip.PackingSlip{}, err
	}
	return packing_slip.Decode(entityMap), err
}

// GetPackingSlip retrieves a single PackingSlip entity by its document path.
func GetPackingSlipByPath(path string) (packing_slip.PackingSlip, error) {
	entityMap, err := PackingSlipRepo.Read(path)
	return packing_slip.Decode(entityMap), err
}

// FindPackingSlip retrieves a PackingSlip entity according to given queries.
func FindPackingSlip(queries []datastore.Query) (packing_slip.PackingSlip, error) {
	entityMap, err := PackingSlipRepo.Find(packing_slip.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return packing_slip.PackingSlip{}, err
	}
	return packing_slip.Decode(entityMap), err
}

// GetAllPackingSlips retrieves all PackingSlip entities.
func GetAllPackingSlips() ([]packing_slip.PackingSlip, error) {
	entityMaps, err := PackingSlipRepo.ReadAll(packing_slip.GetCollectionPath())
	if err != nil {
		return []packing_slip.PackingSlip{}, err
	}
	return packing_slip.DecodeAll(entityMaps), nil
}


// GetAllPackingSlipsPaginated retrieves all PackingSlip entities in a paginated manner.
func GetAllPackingSlipsPaginated(pagination datastore.Pagination) ([]packing_slip.PackingSlip, datastore.Pagination, error) {
	entityMaps, pagination, err := PackingSlipRepo.ReadPaginated(packing_slip.GetCollectionPath(), pagination)
	if err != nil {
		return []packing_slip.PackingSlip{}, pagination, err
	}
	return packing_slip.DecodeAll(entityMaps), pagination, nil
}

// QueryPackingSlips retrieves all PackingSlip entities according to given queries.
func QueryPackingSlips(queries []datastore.Query, pagination datastore.Pagination) ([]packing_slip.PackingSlip, error) {
	entityMaps, err := PackingSlipRepo.Query(packing_slip.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []packing_slip.PackingSlip{}, err
	}
	return packing_slip.DecodeAll(entityMaps), nil
}

// QueryPackingSlipsGroup retrieves all PackingSlip entities according to given queries.
func QueryPackingSlipsGroup(queries []datastore.Query) ([]packing_slip.PackingSlip, error) {
	entityMaps, err := PackingSlipRepo.QueryGroup("packing_slips", queries)
	if err != nil {
		return []packing_slip.PackingSlip{}, err
	}
	return packing_slip.DecodeAll(entityMaps), nil
}

// CreatePackingSlip creates a new PackingSlip entity.
func CreatePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error) {
	return PackingSlipRepo.Create(&entity, editorID)
}

// UpdatePackingSlip updates an existing PackingSlip entity.
func UpdatePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error) {
	return PackingSlipRepo.Update(&entity, editorID)
}

// SavePackingSlip saves a PackingSlip entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePackingSlip(entity packing_slip.PackingSlip, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePackingSlip(entity, editorID)
	} else {
		return UpdatePackingSlip(entity, editorID)
	}
}

// DeletePackingSlip deletes a PackingSlip entity by its parents' path and packingSlipID.
func DeletePackingSlip(packingSlipID string) error {
	return PackingSlipRepo.Delete(packing_slip.GetDocumentPath(packingSlipID))
}
