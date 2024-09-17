package repo

import (
	"github.com/bakeable/bkry/data/entities/cloud_function"
	"github.com/bakeable/bkry/data/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var CloudFunctionRepo = repository.NewRepository[*cloud_function.CloudFunction]()

// GetCloudFunction retrieves a single CloudFunction entity by its ID and cloudFunctionID.
func GetCloudFunction(cloudFunctionID string) (cloud_function.CloudFunction, error) {
	entityMap, err := CloudFunctionRepo.Read(cloud_function.GetDocumentPath( cloudFunctionID))
	return cloud_function.Decode(entityMap), err
}

// GetCloudFunctionOrNew retrieves a single CloudFunction entity by its ID and cloudFunctionID.
func GetCloudFunctionOrNew(cloudFunctionID string) (cloud_function.CloudFunction, error) {
	entityMap, err := CloudFunctionRepo.Read(cloud_function.GetDocumentPath( cloudFunctionID))
	if err != nil || entityMap == nil {
		return cloud_function.CloudFunction{}, err
	}
	return cloud_function.Decode(entityMap), err
}

// GetCloudFunction retrieves a single CloudFunction entity by its document path.
func GetCloudFunctionByPath(path string) (cloud_function.CloudFunction, error) {
	entityMap, err := CloudFunctionRepo.Read(path)
	return cloud_function.Decode(entityMap), err
}

// FindCloudFunction retrieves a CloudFunction entity according to given queries.
func FindCloudFunction(queries []datastore.Query) (cloud_function.CloudFunction, error) {
	entityMap, err := CloudFunctionRepo.Find(cloud_function.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return cloud_function.CloudFunction{}, err
	}
	return cloud_function.Decode(entityMap), err
}

// GetAllCloudFunctions retrieves all CloudFunction entities.
func GetAllCloudFunctions() ([]cloud_function.CloudFunction, error) {
	entityMaps, err := CloudFunctionRepo.ReadAll(cloud_function.GetCollectionPath())
	if err != nil {
		return []cloud_function.CloudFunction{}, err
	}
	return cloud_function.DecodeAll(entityMaps), nil
}


// GetAllCloudFunctionsPaginated retrieves all CloudFunction entities in a paginated manner.
func GetAllCloudFunctionsPaginated(pagination datastore.Pagination) ([]cloud_function.CloudFunction, datastore.Pagination, error) {
	entityMaps, pagination, err := CloudFunctionRepo.ReadPaginated(cloud_function.GetCollectionPath(), pagination)
	if err != nil {
		return []cloud_function.CloudFunction{}, pagination, err
	}
	return cloud_function.DecodeAll(entityMaps), pagination, nil
}

// QueryCloudFunctions retrieves all CloudFunction entities according to given queries.
func QueryCloudFunctions(queries []datastore.Query, pagination datastore.Pagination) ([]cloud_function.CloudFunction, error) {
	entityMaps, err := CloudFunctionRepo.Query(cloud_function.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []cloud_function.CloudFunction{}, err
	}
	return cloud_function.DecodeAll(entityMaps), nil
}

// QueryCloudFunctionsGroup retrieves all CloudFunction entities according to given queries.
func QueryCloudFunctionsGroup(queries []datastore.Query) ([]cloud_function.CloudFunction, error) {
	entityMaps, err := CloudFunctionRepo.QueryGroup("cloud_functions", queries)
	if err != nil {
		return []cloud_function.CloudFunction{}, err
	}
	return cloud_function.DecodeAll(entityMaps), nil
}

// CreateCloudFunction creates a new CloudFunction entity.
func CreateCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error) {
	return CloudFunctionRepo.Create(&entity, editorID)
}

// UpdateCloudFunction updates an existing CloudFunction entity.
func UpdateCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error) {
	return CloudFunctionRepo.Update(&entity, editorID)
}

// SaveCloudFunction saves a CloudFunction entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveCloudFunction(entity cloud_function.CloudFunction, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateCloudFunction(entity, editorID)
	} else {
		return UpdateCloudFunction(entity, editorID)
	}
}

// DeleteCloudFunction deletes a CloudFunction entity by its parents' path and cloudFunctionID.
func DeleteCloudFunction(cloudFunctionID string) error {
	return CloudFunctionRepo.Delete(cloud_function.GetDocumentPath(cloudFunctionID))
}
