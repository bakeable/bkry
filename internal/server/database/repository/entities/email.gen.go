package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/email"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var EmailRepo = repository.NewRepository[*email.Email]()

// GetEmail retrieves a single Email entity by its ID and emailID.
func GetEmail(emailID string) (email.Email, error) {
	entityMap, err := EmailRepo.Read(email.GetDocumentPath( emailID))
	return email.Decode(entityMap), err
}

// GetEmailOrNew retrieves a single Email entity by its ID and emailID.
func GetEmailOrNew(emailID string) (email.Email, error) {
	entityMap, err := EmailRepo.Read(email.GetDocumentPath( emailID))
	if err != nil || entityMap == nil {
		return email.Email{}, err
	}
	return email.Decode(entityMap), err
}

// GetEmail retrieves a single Email entity by its document path.
func GetEmailByPath(path string) (email.Email, error) {
	entityMap, err := EmailRepo.Read(path)
	return email.Decode(entityMap), err
}

// FindEmail retrieves a Email entity according to given queries.
func FindEmail(queries []datastore.Query) (email.Email, error) {
	entityMap, err := EmailRepo.Find(email.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return email.Email{}, err
	}
	return email.Decode(entityMap), err
}

// GetAllEmails retrieves all Email entities.
func GetAllEmails() ([]email.Email, error) {
	entityMaps, err := EmailRepo.ReadAll(email.GetCollectionPath())
	if err != nil {
		return []email.Email{}, err
	}
	return email.DecodeAll(entityMaps), nil
}


// GetAllEmailsPaginated retrieves all Email entities in a paginated manner.
func GetAllEmailsPaginated(pagination datastore.Pagination) ([]email.Email, datastore.Pagination, error) {
	entityMaps, pagination, err := EmailRepo.ReadPaginated(email.GetCollectionPath(), pagination)
	if err != nil {
		return []email.Email{}, pagination, err
	}
	return email.DecodeAll(entityMaps), pagination, nil
}

// QueryEmails retrieves all Email entities according to given queries.
func QueryEmails(queries []datastore.Query, pagination datastore.Pagination) ([]email.Email, error) {
	entityMaps, err := EmailRepo.Query(email.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []email.Email{}, err
	}
	return email.DecodeAll(entityMaps), nil
}

// QueryEmailsGroup retrieves all Email entities according to given queries.
func QueryEmailsGroup(queries []datastore.Query) ([]email.Email, error) {
	entityMaps, err := EmailRepo.QueryGroup("emails", queries)
	if err != nil {
		return []email.Email{}, err
	}
	return email.DecodeAll(entityMaps), nil
}

// CreateEmail creates a new Email entity.
func CreateEmail(entity email.Email, editorID *string) (string, error) {
	return EmailRepo.Create(&entity, editorID)
}

// UpdateEmail updates an existing Email entity.
func UpdateEmail(entity email.Email, editorID *string) (string, error) {
	return EmailRepo.Update(&entity, editorID)
}

// SaveEmail saves a Email entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveEmail(entity email.Email, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateEmail(entity, editorID)
	} else {
		return UpdateEmail(entity, editorID)
	}
}

// DeleteEmail deletes a Email entity by its parents' path and emailID.
func DeleteEmail(emailID string) error {
	return EmailRepo.Delete(email.GetDocumentPath(emailID))
}
