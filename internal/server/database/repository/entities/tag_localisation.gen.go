package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/tag_localisation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var TagLocalisationRepo = repository.NewFirestoreRepository[*tag_localisation.TagLocalisation]()

// GetTagLocalisation retrieves a single TagLocalisation entity by its ID and tagLocalisationID.
func GetTagLocalisation(TagID string, tagLocalisationID string) (tag_localisation.TagLocalisation, error) {
	entityMap, err := TagLocalisationRepo.Read(tag_localisation.GetDocumentPath(TagID,  tagLocalisationID))
	return tag_localisation.Decode(entityMap), err
}

// GetTagLocalisationOrNew retrieves a single TagLocalisation entity by its ID and tagLocalisationID.
func GetTagLocalisationOrNew(TagID string, tagLocalisationID string) (tag_localisation.TagLocalisation, error) {
	entityMap, err := TagLocalisationRepo.Read(tag_localisation.GetDocumentPath(TagID,  tagLocalisationID))
	if err != nil || entityMap == nil {
		return tag_localisation.TagLocalisation{}, err
	}
	return tag_localisation.Decode(entityMap), err
}

// GetTagLocalisation retrieves a single TagLocalisation entity by its document path.
func GetTagLocalisationByPath(path string) (tag_localisation.TagLocalisation, error) {
	entityMap, err := TagLocalisationRepo.Read(path)
	return tag_localisation.Decode(entityMap), err
}

// FindTagLocalisation retrieves a TagLocalisation entity according to given queries.
func FindTagLocalisation(tagID string, queries []database.Query) (tag_localisation.TagLocalisation, error) {
	entityMap, err := TagLocalisationRepo.Find(tag_localisation.GetCollectionPath(tagID), queries)
	if err != nil || entityMap == nil {
		return tag_localisation.TagLocalisation{}, err
	}
	return tag_localisation.Decode(entityMap), err
}

// GetAllTagLocalisations retrieves all TagLocalisation entities.
func GetAllTagLocalisations(tagID string) ([]tag_localisation.TagLocalisation, error) {
	entityMaps, err := TagLocalisationRepo.ReadAll(tag_localisation.GetCollectionPath(tagID))
	if err != nil {
		return []tag_localisation.TagLocalisation{}, err
	}
	return tag_localisation.DecodeAll(entityMaps), nil
}


// GetAllTagLocalisationsPaginated retrieves all TagLocalisation entities in a paginated manner.
func GetAllTagLocalisationsPaginated(tagID string, pagination database.Pagination) ([]tag_localisation.TagLocalisation, database.Pagination, error) {
	entityMaps, pagination, err := TagLocalisationRepo.ReadPaginated(tag_localisation.GetCollectionPath(tagID), pagination)
	if err != nil {
		return []tag_localisation.TagLocalisation{}, pagination, err
	}
	return tag_localisation.DecodeAll(entityMaps), pagination, nil
}

// QueryTagLocalisations retrieves all TagLocalisation entities according to given queries.
func QueryTagLocalisations(tagID string, queries []database.Query, pagination database.Pagination) ([]tag_localisation.TagLocalisation, error) {
	entityMaps, err := TagLocalisationRepo.Query(tag_localisation.GetCollectionPath(tagID), queries, pagination)
	if err != nil {
		return []tag_localisation.TagLocalisation{}, err
	}
	return tag_localisation.DecodeAll(entityMaps), nil
}

// QueryTagLocalisationsGroup retrieves all TagLocalisation entities according to given queries.
func QueryTagLocalisationsGroup(queries []database.Query) ([]tag_localisation.TagLocalisation, error) {
	entityMaps, err := TagLocalisationRepo.QueryGroup("tag_localisations", queries)
	if err != nil {
		return []tag_localisation.TagLocalisation{}, err
	}
	return tag_localisation.DecodeAll(entityMaps), nil
}

// CreateTagLocalisation creates a new TagLocalisation entity.
func CreateTagLocalisation(TagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error) {
	entity.TagID = TagID
	return TagLocalisationRepo.Create(&entity, editorID)
}

// UpdateTagLocalisation updates an existing TagLocalisation entity.
func UpdateTagLocalisation(TagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error) {
	entity.TagID = TagID
	return TagLocalisationRepo.Update(&entity, editorID)
}

// SaveTagLocalisation saves a TagLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveTagLocalisation(TagID string, entity tag_localisation.TagLocalisation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateTagLocalisation(TagID, entity, editorID)
	} else {
		return UpdateTagLocalisation(TagID, entity, editorID)
	}
}

// DeleteTagLocalisation deletes a TagLocalisation entity by its parents' path and tagLocalisationID.
func DeleteTagLocalisation(TagID string, tagLocalisationID string) error {
	return TagLocalisationRepo.Delete(tag_localisation.GetDocumentPath(TagID, tagLocalisationID))
}
