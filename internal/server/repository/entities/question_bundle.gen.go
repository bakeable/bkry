package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/question_bundle"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var QuestionBundleRepo = repository.NewFirestoreRepository[*question_bundle.QuestionBundle]()

// GetQuestionBundle retrieves a single QuestionBundle entity by its ID and questionBundleID.
func GetQuestionBundle(questionBundleID string) (question_bundle.QuestionBundle, error) {
	entityMap, err := QuestionBundleRepo.Read(question_bundle.GetDocumentPath( questionBundleID))
	return question_bundle.Decode(entityMap), err
}

// GetQuestionBundleOrNew retrieves a single QuestionBundle entity by its ID and questionBundleID.
func GetQuestionBundleOrNew(questionBundleID string) (question_bundle.QuestionBundle, error) {
	entityMap, err := QuestionBundleRepo.Read(question_bundle.GetDocumentPath( questionBundleID))
	if err != nil || entityMap == nil {
		return question_bundle.QuestionBundle{}, err
	}
	return question_bundle.Decode(entityMap), err
}

// GetQuestionBundle retrieves a single QuestionBundle entity by its document path.
func GetQuestionBundleByPath(path string) (question_bundle.QuestionBundle, error) {
	entityMap, err := QuestionBundleRepo.Read(path)
	return question_bundle.Decode(entityMap), err
}

// FindQuestionBundle retrieves a QuestionBundle entity according to given queries.
func FindQuestionBundle(queries []database.Query) (question_bundle.QuestionBundle, error) {
	entityMap, err := QuestionBundleRepo.Find(question_bundle.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return question_bundle.QuestionBundle{}, err
	}
	return question_bundle.Decode(entityMap), err
}

// GetAllQuestionBundles retrieves all QuestionBundle entities.
func GetAllQuestionBundles() ([]question_bundle.QuestionBundle, error) {
	entityMaps, err := QuestionBundleRepo.ReadAll(question_bundle.GetCollectionPath())
	if err != nil {
		return []question_bundle.QuestionBundle{}, err
	}
	return question_bundle.DecodeAll(entityMaps), nil
}


// GetAllQuestionBundlesPaginated retrieves all QuestionBundle entities in a paginated manner.
func GetAllQuestionBundlesPaginated(pagination database.Pagination) ([]question_bundle.QuestionBundle, database.Pagination, error) {
	entityMaps, pagination, err := QuestionBundleRepo.ReadPaginated(question_bundle.GetCollectionPath(), pagination)
	if err != nil {
		return []question_bundle.QuestionBundle{}, pagination, err
	}
	return question_bundle.DecodeAll(entityMaps), pagination, nil
}

// QueryQuestionBundles retrieves all QuestionBundle entities according to given queries.
func QueryQuestionBundles(queries []database.Query, pagination database.Pagination) ([]question_bundle.QuestionBundle, error) {
	entityMaps, err := QuestionBundleRepo.Query(question_bundle.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []question_bundle.QuestionBundle{}, err
	}
	return question_bundle.DecodeAll(entityMaps), nil
}

// QueryQuestionBundlesGroup retrieves all QuestionBundle entities according to given queries.
func QueryQuestionBundlesGroup(queries []database.Query) ([]question_bundle.QuestionBundle, error) {
	entityMaps, err := QuestionBundleRepo.QueryGroup("question_bundles", queries)
	if err != nil {
		return []question_bundle.QuestionBundle{}, err
	}
	return question_bundle.DecodeAll(entityMaps), nil
}

// CreateQuestionBundle creates a new QuestionBundle entity.
func CreateQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error) {
	return QuestionBundleRepo.Create(&entity, editorID)
}

// UpdateQuestionBundle updates an existing QuestionBundle entity.
func UpdateQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error) {
	return QuestionBundleRepo.Update(&entity, editorID)
}

// SaveQuestionBundle saves a QuestionBundle entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveQuestionBundle(entity question_bundle.QuestionBundle, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateQuestionBundle(entity, editorID)
	} else {
		return UpdateQuestionBundle(entity, editorID)
	}
}

// DeleteQuestionBundle deletes a QuestionBundle entity by its parents' path and questionBundleID.
func DeleteQuestionBundle(questionBundleID string) error {
	return QuestionBundleRepo.Delete(question_bundle.GetDocumentPath(questionBundleID))
}
