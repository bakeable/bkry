package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/api_key"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ApiKeyRepo = repository.NewRepository[*api_key.ApiKey]()

// GetApiKey retrieves a single ApiKey entity by its ID and apiKeyID.
func GetApiKey(apiKeyID string) (api_key.ApiKey, error) {
	entityMap, err := ApiKeyRepo.Read(api_key.GetDocumentPath( apiKeyID))
	return api_key.Decode(entityMap), err
}

// GetApiKeyOrNew retrieves a single ApiKey entity by its ID and apiKeyID.
func GetApiKeyOrNew(apiKeyID string) (api_key.ApiKey, error) {
	entityMap, err := ApiKeyRepo.Read(api_key.GetDocumentPath( apiKeyID))
	if err != nil || entityMap == nil {
		return api_key.ApiKey{}, err
	}
	return api_key.Decode(entityMap), err
}

// GetApiKey retrieves a single ApiKey entity by its document path.
func GetApiKeyByPath(path string) (api_key.ApiKey, error) {
	entityMap, err := ApiKeyRepo.Read(path)
	return api_key.Decode(entityMap), err
}

// FindApiKey retrieves a ApiKey entity according to given queries.
func FindApiKey(queries []datastore.Query) (api_key.ApiKey, error) {
	entityMap, err := ApiKeyRepo.Find(api_key.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return api_key.ApiKey{}, err
	}
	return api_key.Decode(entityMap), err
}

// GetAllApiKeys retrieves all ApiKey entities.
func GetAllApiKeys() ([]api_key.ApiKey, error) {
	entityMaps, err := ApiKeyRepo.ReadAll(api_key.GetCollectionPath())
	if err != nil {
		return []api_key.ApiKey{}, err
	}
	return api_key.DecodeAll(entityMaps), nil
}


// GetAllApiKeysPaginated retrieves all ApiKey entities in a paginated manner.
func GetAllApiKeysPaginated(pagination datastore.Pagination) ([]api_key.ApiKey, datastore.Pagination, error) {
	entityMaps, pagination, err := ApiKeyRepo.ReadPaginated(api_key.GetCollectionPath(), pagination)
	if err != nil {
		return []api_key.ApiKey{}, pagination, err
	}
	return api_key.DecodeAll(entityMaps), pagination, nil
}

// QueryApiKeys retrieves all ApiKey entities according to given queries.
func QueryApiKeys(queries []datastore.Query, pagination datastore.Pagination) ([]api_key.ApiKey, error) {
	entityMaps, err := ApiKeyRepo.Query(api_key.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []api_key.ApiKey{}, err
	}
	return api_key.DecodeAll(entityMaps), nil
}

// QueryApiKeysGroup retrieves all ApiKey entities according to given queries.
func QueryApiKeysGroup(queries []datastore.Query) ([]api_key.ApiKey, error) {
	entityMaps, err := ApiKeyRepo.QueryGroup("api_keys", queries)
	if err != nil {
		return []api_key.ApiKey{}, err
	}
	return api_key.DecodeAll(entityMaps), nil
}

// CreateApiKey creates a new ApiKey entity.
func CreateApiKey(entity api_key.ApiKey, editorID *string) (string, error) {
	return ApiKeyRepo.Create(&entity, editorID)
}

// UpdateApiKey updates an existing ApiKey entity.
func UpdateApiKey(entity api_key.ApiKey, editorID *string) (string, error) {
	return ApiKeyRepo.Update(&entity, editorID)
}

// SaveApiKey saves a ApiKey entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveApiKey(entity api_key.ApiKey, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateApiKey(entity, editorID)
	} else {
		return UpdateApiKey(entity, editorID)
	}
}

// DeleteApiKey deletes a ApiKey entity by its parents' path and apiKeyID.
func DeleteApiKey(apiKeyID string) error {
	return ApiKeyRepo.Delete(api_key.GetDocumentPath(apiKeyID))
}
