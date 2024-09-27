package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/batch_export"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var BatchExportRepo = repository.NewRepository[*batch_export.BatchExport]()

// GetBatchExport retrieves a single BatchExport entity by its ID and batchExportID.
func GetBatchExport(batchExportID string) (batch_export.BatchExport, error) {
	entityMap, err := BatchExportRepo.Read(batch_export.GetDocumentPath( batchExportID))
	return batch_export.Decode(entityMap), err
}

// GetBatchExportOrNew retrieves a single BatchExport entity by its ID and batchExportID.
func GetBatchExportOrNew(batchExportID string) (batch_export.BatchExport, error) {
	entityMap, err := BatchExportRepo.Read(batch_export.GetDocumentPath( batchExportID))
	if err != nil || entityMap == nil {
		return batch_export.BatchExport{}, err
	}
	return batch_export.Decode(entityMap), err
}

// GetBatchExport retrieves a single BatchExport entity by its document path.
func GetBatchExportByPath(path string) (batch_export.BatchExport, error) {
	entityMap, err := BatchExportRepo.Read(path)
	return batch_export.Decode(entityMap), err
}

// FindBatchExport retrieves a BatchExport entity according to given queries.
func FindBatchExport(queries []datastore.Query) (batch_export.BatchExport, error) {
	entityMap, err := BatchExportRepo.Find(batch_export.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return batch_export.BatchExport{}, err
	}
	return batch_export.Decode(entityMap), err
}

// GetAllBatchExports retrieves all BatchExport entities.
func GetAllBatchExports() ([]batch_export.BatchExport, error) {
	entityMaps, err := BatchExportRepo.ReadAll(batch_export.GetCollectionPath())
	if err != nil {
		return []batch_export.BatchExport{}, err
	}
	return batch_export.DecodeAll(entityMaps), nil
}


// GetAllBatchExportsPaginated retrieves all BatchExport entities in a paginated manner.
func GetAllBatchExportsPaginated(pagination datastore.Pagination) ([]batch_export.BatchExport, datastore.Pagination, error) {
	entityMaps, pagination, err := BatchExportRepo.ReadPaginated(batch_export.GetCollectionPath(), pagination)
	if err != nil {
		return []batch_export.BatchExport{}, pagination, err
	}
	return batch_export.DecodeAll(entityMaps), pagination, nil
}

// QueryBatchExports retrieves all BatchExport entities according to given queries.
func QueryBatchExports(queries []datastore.Query, pagination datastore.Pagination) ([]batch_export.BatchExport, error) {
	entityMaps, err := BatchExportRepo.Query(batch_export.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []batch_export.BatchExport{}, err
	}
	return batch_export.DecodeAll(entityMaps), nil
}

// QueryBatchExportsGroup retrieves all BatchExport entities according to given queries.
func QueryBatchExportsGroup(queries []datastore.Query) ([]batch_export.BatchExport, error) {
	entityMaps, err := BatchExportRepo.QueryGroup("batch_exports", queries)
	if err != nil {
		return []batch_export.BatchExport{}, err
	}
	return batch_export.DecodeAll(entityMaps), nil
}

// CreateBatchExport creates a new BatchExport entity.
func CreateBatchExport(entity batch_export.BatchExport, editorID *string) (string, error) {
	return BatchExportRepo.Create(&entity, editorID)
}

// UpdateBatchExport updates an existing BatchExport entity.
func UpdateBatchExport(entity batch_export.BatchExport, editorID *string) (string, error) {
	return BatchExportRepo.Update(&entity, editorID)
}

// SaveBatchExport saves a BatchExport entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveBatchExport(entity batch_export.BatchExport, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateBatchExport(entity, editorID)
	} else {
		return UpdateBatchExport(entity, editorID)
	}
}

// DeleteBatchExport deletes a BatchExport entity by its parents' path and batchExportID.
func DeleteBatchExport(batchExportID string) error {
	return BatchExportRepo.Delete(batch_export.GetDocumentPath(batchExportID))
}
