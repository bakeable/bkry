package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/tag"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var TagRepo = repository.NewFirestoreRepository[*tag.Tag]()

// GetTag retrieves a single Tag entity by its ID and tagID.
func GetTag(tagID string) (tag.Tag, error) {
	entityMap, err := TagRepo.Read(tag.GetDocumentPath( tagID))
	return tag.Decode(entityMap), err
}

// GetTagOrNew retrieves a single Tag entity by its ID and tagID.
func GetTagOrNew(tagID string) (tag.Tag, error) {
	entityMap, err := TagRepo.Read(tag.GetDocumentPath( tagID))
	if err != nil || entityMap == nil {
		return tag.Tag{}, err
	}
	return tag.Decode(entityMap), err
}

// GetTag retrieves a single Tag entity by its document path.
func GetTagByPath(path string) (tag.Tag, error) {
	entityMap, err := TagRepo.Read(path)
	return tag.Decode(entityMap), err
}

// FindTag retrieves a Tag entity according to given queries.
func FindTag(queries []database.Query) (tag.Tag, error) {
	entityMap, err := TagRepo.Find(tag.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return tag.Tag{}, err
	}
	return tag.Decode(entityMap), err
}

// GetAllTags retrieves all Tag entities.
func GetAllTags() ([]tag.Tag, error) {
	entityMaps, err := TagRepo.ReadAll(tag.GetCollectionPath())
	if err != nil {
		return []tag.Tag{}, err
	}
	return tag.DecodeAll(entityMaps), nil
}


// GetAllTagsPaginated retrieves all Tag entities in a paginated manner.
func GetAllTagsPaginated(pagination database.Pagination) ([]tag.Tag, database.Pagination, error) {
	entityMaps, pagination, err := TagRepo.ReadPaginated(tag.GetCollectionPath(), pagination)
	if err != nil {
		return []tag.Tag{}, pagination, err
	}
	return tag.DecodeAll(entityMaps), pagination, nil
}

// QueryTags retrieves all Tag entities according to given queries.
func QueryTags(queries []database.Query, pagination database.Pagination) ([]tag.Tag, error) {
	entityMaps, err := TagRepo.Query(tag.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []tag.Tag{}, err
	}
	return tag.DecodeAll(entityMaps), nil
}

// QueryTagsGroup retrieves all Tag entities according to given queries.
func QueryTagsGroup(queries []database.Query) ([]tag.Tag, error) {
	entityMaps, err := TagRepo.QueryGroup("tags", queries)
	if err != nil {
		return []tag.Tag{}, err
	}
	return tag.DecodeAll(entityMaps), nil
}

// CreateTag creates a new Tag entity.
func CreateTag(entity tag.Tag, editorID *string) (string, error) {
	return TagRepo.Create(&entity, editorID)
}

// UpdateTag updates an existing Tag entity.
func UpdateTag(entity tag.Tag, editorID *string) (string, error) {
	return TagRepo.Update(&entity, editorID)
}

// SaveTag saves a Tag entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveTag(entity tag.Tag, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateTag(entity, editorID)
	} else {
		return UpdateTag(entity, editorID)
	}
}

// DeleteTag deletes a Tag entity by its parents' path and tagID.
func DeleteTag(tagID string) error {
	return TagRepo.Delete(tag.GetDocumentPath(tagID))
}
