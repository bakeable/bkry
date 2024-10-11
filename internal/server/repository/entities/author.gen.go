package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/author"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var AuthorRepo = repository.NewFirestoreRepository[*author.Author]()

// GetAuthor retrieves a single Author entity by its ID and authorID.
func GetAuthor(authorID string) (author.Author, error) {
	entityMap, err := AuthorRepo.Read(author.GetDocumentPath( authorID))
	return author.Decode(entityMap), err
}

// GetAuthorOrNew retrieves a single Author entity by its ID and authorID.
func GetAuthorOrNew(authorID string) (author.Author, error) {
	entityMap, err := AuthorRepo.Read(author.GetDocumentPath( authorID))
	if err != nil || entityMap == nil {
		return author.Author{}, err
	}
	return author.Decode(entityMap), err
}

// GetAuthor retrieves a single Author entity by its document path.
func GetAuthorByPath(path string) (author.Author, error) {
	entityMap, err := AuthorRepo.Read(path)
	return author.Decode(entityMap), err
}

// FindAuthor retrieves a Author entity according to given queries.
func FindAuthor(queries []database.Query) (author.Author, error) {
	entityMap, err := AuthorRepo.Find(author.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return author.Author{}, err
	}
	return author.Decode(entityMap), err
}

// GetAllAuthors retrieves all Author entities.
func GetAllAuthors() ([]author.Author, error) {
	entityMaps, err := AuthorRepo.ReadAll(author.GetCollectionPath())
	if err != nil {
		return []author.Author{}, err
	}
	return author.DecodeAll(entityMaps), nil
}


// GetAllAuthorsPaginated retrieves all Author entities in a paginated manner.
func GetAllAuthorsPaginated(pagination database.Pagination) ([]author.Author, database.Pagination, error) {
	entityMaps, pagination, err := AuthorRepo.ReadPaginated(author.GetCollectionPath(), pagination)
	if err != nil {
		return []author.Author{}, pagination, err
	}
	return author.DecodeAll(entityMaps), pagination, nil
}

// QueryAuthors retrieves all Author entities according to given queries.
func QueryAuthors(queries []database.Query, pagination database.Pagination) ([]author.Author, error) {
	entityMaps, err := AuthorRepo.Query(author.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []author.Author{}, err
	}
	return author.DecodeAll(entityMaps), nil
}

// QueryAuthorsGroup retrieves all Author entities according to given queries.
func QueryAuthorsGroup(queries []database.Query) ([]author.Author, error) {
	entityMaps, err := AuthorRepo.QueryGroup("authors", queries)
	if err != nil {
		return []author.Author{}, err
	}
	return author.DecodeAll(entityMaps), nil
}

// CreateAuthor creates a new Author entity.
func CreateAuthor(entity author.Author, editorID *string) (string, error) {
	return AuthorRepo.Create(&entity, editorID)
}

// UpdateAuthor updates an existing Author entity.
func UpdateAuthor(entity author.Author, editorID *string) (string, error) {
	return AuthorRepo.Update(&entity, editorID)
}

// SaveAuthor saves a Author entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveAuthor(entity author.Author, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateAuthor(entity, editorID)
	} else {
		return UpdateAuthor(entity, editorID)
	}
}

// DeleteAuthor deletes a Author entity by its parents' path and authorID.
func DeleteAuthor(authorID string) error {
	return AuthorRepo.Delete(author.GetDocumentPath(authorID))
}
