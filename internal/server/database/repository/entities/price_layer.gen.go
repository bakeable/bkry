package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/price_layer"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var PriceLayerRepo = repository.NewRepository[*price_layer.PriceLayer]()

// GetPriceLayer retrieves a single PriceLayer entity by its ID and priceLayerID.
func GetPriceLayer(priceLayerID string) (price_layer.PriceLayer, error) {
	entityMap, err := PriceLayerRepo.Read(price_layer.GetDocumentPath( priceLayerID))
	return price_layer.Decode(entityMap), err
}

// GetPriceLayerOrNew retrieves a single PriceLayer entity by its ID and priceLayerID.
func GetPriceLayerOrNew(priceLayerID string) (price_layer.PriceLayer, error) {
	entityMap, err := PriceLayerRepo.Read(price_layer.GetDocumentPath( priceLayerID))
	if err != nil || entityMap == nil {
		return price_layer.PriceLayer{}, err
	}
	return price_layer.Decode(entityMap), err
}

// GetPriceLayer retrieves a single PriceLayer entity by its document path.
func GetPriceLayerByPath(path string) (price_layer.PriceLayer, error) {
	entityMap, err := PriceLayerRepo.Read(path)
	return price_layer.Decode(entityMap), err
}

// FindPriceLayer retrieves a PriceLayer entity according to given queries.
func FindPriceLayer(queries []datastore.Query) (price_layer.PriceLayer, error) {
	entityMap, err := PriceLayerRepo.Find(price_layer.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return price_layer.PriceLayer{}, err
	}
	return price_layer.Decode(entityMap), err
}

// GetAllPriceLayers retrieves all PriceLayer entities.
func GetAllPriceLayers() ([]price_layer.PriceLayer, error) {
	entityMaps, err := PriceLayerRepo.ReadAll(price_layer.GetCollectionPath())
	if err != nil {
		return []price_layer.PriceLayer{}, err
	}
	return price_layer.DecodeAll(entityMaps), nil
}


// GetAllPriceLayersPaginated retrieves all PriceLayer entities in a paginated manner.
func GetAllPriceLayersPaginated(pagination datastore.Pagination) ([]price_layer.PriceLayer, datastore.Pagination, error) {
	entityMaps, pagination, err := PriceLayerRepo.ReadPaginated(price_layer.GetCollectionPath(), pagination)
	if err != nil {
		return []price_layer.PriceLayer{}, pagination, err
	}
	return price_layer.DecodeAll(entityMaps), pagination, nil
}

// QueryPriceLayers retrieves all PriceLayer entities according to given queries.
func QueryPriceLayers(queries []datastore.Query, pagination datastore.Pagination) ([]price_layer.PriceLayer, error) {
	entityMaps, err := PriceLayerRepo.Query(price_layer.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []price_layer.PriceLayer{}, err
	}
	return price_layer.DecodeAll(entityMaps), nil
}

// QueryPriceLayersGroup retrieves all PriceLayer entities according to given queries.
func QueryPriceLayersGroup(queries []datastore.Query) ([]price_layer.PriceLayer, error) {
	entityMaps, err := PriceLayerRepo.QueryGroup("price_layers", queries)
	if err != nil {
		return []price_layer.PriceLayer{}, err
	}
	return price_layer.DecodeAll(entityMaps), nil
}

// CreatePriceLayer creates a new PriceLayer entity.
func CreatePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error) {
	return PriceLayerRepo.Create(&entity, editorID)
}

// UpdatePriceLayer updates an existing PriceLayer entity.
func UpdatePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error) {
	return PriceLayerRepo.Update(&entity, editorID)
}

// SavePriceLayer saves a PriceLayer entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SavePriceLayer(entity price_layer.PriceLayer, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreatePriceLayer(entity, editorID)
	} else {
		return UpdatePriceLayer(entity, editorID)
	}
}

// DeletePriceLayer deletes a PriceLayer entity by its parents' path and priceLayerID.
func DeletePriceLayer(priceLayerID string) error {
	return PriceLayerRepo.Delete(price_layer.GetDocumentPath(priceLayerID))
}
