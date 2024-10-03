package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/media"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var MediaRepo = repository.NewFirestoreRepository[*media.Media]()

// GetMedia retrieves a single Media entity by its ID and mediaID.
func GetMedia(mediaID string) (media.Media, error) {
	entityMap, err := MediaRepo.Read(media.GetDocumentPath( mediaID))
	return media.Decode(entityMap), err
}

// GetMediaOrNew retrieves a single Media entity by its ID and mediaID.
func GetMediaOrNew(mediaID string) (media.Media, error) {
	entityMap, err := MediaRepo.Read(media.GetDocumentPath( mediaID))
	if err != nil || entityMap == nil {
		return media.Media{}, err
	}
	return media.Decode(entityMap), err
}

// GetMedia retrieves a single Media entity by its document path.
func GetMediaByPath(path string) (media.Media, error) {
	entityMap, err := MediaRepo.Read(path)
	return media.Decode(entityMap), err
}

// FindMedia retrieves a Media entity according to given queries.
func FindMedia(queries []database.Query) (media.Media, error) {
	entityMap, err := MediaRepo.Find(media.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return media.Media{}, err
	}
	return media.Decode(entityMap), err
}

// GetAllMedia retrieves all Media entities.
func GetAllMedia() ([]media.Media, error) {
	entityMaps, err := MediaRepo.ReadAll(media.GetCollectionPath())
	if err != nil {
		return []media.Media{}, err
	}
	return media.DecodeAll(entityMaps), nil
}


// GetAllMediaPaginated retrieves all Media entities in a paginated manner.
func GetAllMediaPaginated(pagination database.Pagination) ([]media.Media, database.Pagination, error) {
	entityMaps, pagination, err := MediaRepo.ReadPaginated(media.GetCollectionPath(), pagination)
	if err != nil {
		return []media.Media{}, pagination, err
	}
	return media.DecodeAll(entityMaps), pagination, nil
}

// QueryMedia retrieves all Media entities according to given queries.
func QueryMedia(queries []database.Query, pagination database.Pagination) ([]media.Media, error) {
	entityMaps, err := MediaRepo.Query(media.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []media.Media{}, err
	}
	return media.DecodeAll(entityMaps), nil
}

// QueryMediaGroup retrieves all Media entities according to given queries.
func QueryMediaGroup(queries []database.Query) ([]media.Media, error) {
	entityMaps, err := MediaRepo.QueryGroup("media", queries)
	if err != nil {
		return []media.Media{}, err
	}
	return media.DecodeAll(entityMaps), nil
}

// CreateMedia creates a new Media entity.
func CreateMedia(entity media.Media, editorID *string) (string, error) {
	return MediaRepo.Create(&entity, editorID)
}

// UpdateMedia updates an existing Media entity.
func UpdateMedia(entity media.Media, editorID *string) (string, error) {
	return MediaRepo.Update(&entity, editorID)
}

// SaveMedia saves a Media entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveMedia(entity media.Media, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateMedia(entity, editorID)
	} else {
		return UpdateMedia(entity, editorID)
	}
}

// DeleteMedia deletes a Media entity by its parents' path and mediaID.
func DeleteMedia(mediaID string) error {
	return MediaRepo.Delete(media.GetDocumentPath(mediaID))
}
