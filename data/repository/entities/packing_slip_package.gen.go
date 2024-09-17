package repo

import (
	"github.com/bakeable/bkry/data/entities/packing_slip_package"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PackingSlipPackageRepo = repository.NewRepository[*packing_slip_package.PackingSlipPackage]()

// GetPackingSlipPackage retrieves a single PackingSlipPackage entity by its ID and packingSlipPackageID.
func GetPackingSlipPackage(packingSlipPackageID string) (packing_slip_package.PackingSlipPackage, error) {
	entityMap, err := PackingSlipPackageRepo.Read(packing_slip_package.GetDocumentPath( packingSlipPackageID))
	return packing_slip_package.Decode(entityMap), err
}

// GetPackingSlipPackageOrNew retrieves a single PackingSlipPackage entity by its ID and packingSlipPackageID.
func GetPackingSlipPackageOrNew(packingSlipPackageID string) (packing_slip_package.PackingSlipPackage, error) {
	entityMap, err := PackingSlipPackageRepo.Read(packing_slip_package.GetDocumentPath( packingSlipPackageID))
	if err != nil || entityMap == nil {
		return packing_slip_package.PackingSlipPackage{}, err
	}
	return packing_slip_package.Decode(entityMap), err
}

// GetPackingSlipPackage retrieves a single PackingSlipPackage entity by its document path.
func GetPackingSlipPackageByPath(path string) (packing_slip_package.PackingSlipPackage, error) {
	entityMap, err := PackingSlipPackageRepo.Read(path)
	return packing_slip_package.Decode(entityMap), err
}

// FindPackingSlipPackage retrieves a PackingSlipPackage entity according to given queries.
func FindPackingSlipPackage(queries []datastore.Query) (packing_slip_package.PackingSlipPackage, error) {
	entityMap, err := PackingSlipPackageRepo.Find(packing_slip_package.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return packing_slip_package.PackingSlipPackage{}, err
	}
	return packing_slip_package.Decode(entityMap), err
}

// GetAllPackingSlipPackages retrieves all PackingSlipPackage entities.
func GetAllPackingSlipPackages() ([]packing_slip_package.PackingSlipPackage, error) {
	entityMaps, err := PackingSlipPackageRepo.ReadAll(packing_slip_package.GetCollectionPath())
	if err != nil {
		return []packing_slip_package.PackingSlipPackage{}, err
	}
	return packing_slip_package.DecodeAll(entityMaps), nil
}


// GetAllPackingSlipPackagesPaginated retrieves all PackingSlipPackage entities in a paginated manner.
func GetAllPackingSlipPackagesPaginated(pagination datastore.Pagination) ([]packing_slip_package.PackingSlipPackage, datastore.Pagination, error) {
	entityMaps, pagination, err := PackingSlipPackageRepo.ReadPaginated(packing_slip_package.GetCollectionPath(), pagination)
	if err != nil {
		return []packing_slip_package.PackingSlipPackage{}, pagination, err
	}
	return packing_slip_package.DecodeAll(entityMaps), pagination, nil
}

// QueryPackingSlipPackages retrieves all PackingSlipPackage entities according to given queries.
func QueryPackingSlipPackages(queries []datastore.Query, pagination datastore.Pagination) ([]packing_slip_package.PackingSlipPackage, error) {
	entityMaps, err := PackingSlipPackageRepo.Query(packing_slip_package.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []packing_slip_package.PackingSlipPackage{}, err
	}
	return packing_slip_package.DecodeAll(entityMaps), nil
}

// QueryPackingSlipPackagesGroup retrieves all PackingSlipPackage entities according to given queries.
func QueryPackingSlipPackagesGroup(queries []datastore.Query) ([]packing_slip_package.PackingSlipPackage, error) {
	entityMaps, err := PackingSlipPackageRepo.QueryGroup("packing_slip_packages", queries)
	if err != nil {
		return []packing_slip_package.PackingSlipPackage{}, err
	}
	return packing_slip_package.DecodeAll(entityMaps), nil
}

// CreatePackingSlipPackage creates a new PackingSlipPackage entity.
func CreatePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error) {
	return PackingSlipPackageRepo.Create(&entity, editorID)
}

// UpdatePackingSlipPackage updates an existing PackingSlipPackage entity.
func UpdatePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error) {
	return PackingSlipPackageRepo.Update(&entity, editorID)
}

// SavePackingSlipPackage saves a PackingSlipPackage entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePackingSlipPackage(entity packing_slip_package.PackingSlipPackage, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePackingSlipPackage(entity, editorID)
	} else {
		return UpdatePackingSlipPackage(entity, editorID)
	}
}

// DeletePackingSlipPackage deletes a PackingSlipPackage entity by its parents' path and packingSlipPackageID.
func DeletePackingSlipPackage(packingSlipPackageID string) error {
	return PackingSlipPackageRepo.Delete(packing_slip_package.GetDocumentPath(packingSlipPackageID))
}
