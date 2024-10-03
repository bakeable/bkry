package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/question_context"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var QuestionContextRepo = repository.NewFirestoreRepository[*question_context.QuestionContext]()

// GetQuestionContext retrieves a single QuestionContext entity by its ID and questionContextID.
func GetQuestionContext(questionContextID string) (question_context.QuestionContext, error) {
	entityMap, err := QuestionContextRepo.Read(question_context.GetDocumentPath( questionContextID))
	return question_context.Decode(entityMap), err
}

// GetQuestionContextOrNew retrieves a single QuestionContext entity by its ID and questionContextID.
func GetQuestionContextOrNew(questionContextID string) (question_context.QuestionContext, error) {
	entityMap, err := QuestionContextRepo.Read(question_context.GetDocumentPath( questionContextID))
	if err != nil || entityMap == nil {
		return question_context.QuestionContext{}, err
	}
	return question_context.Decode(entityMap), err
}

// GetQuestionContext retrieves a single QuestionContext entity by its document path.
func GetQuestionContextByPath(path string) (question_context.QuestionContext, error) {
	entityMap, err := QuestionContextRepo.Read(path)
	return question_context.Decode(entityMap), err
}

// FindQuestionContext retrieves a QuestionContext entity according to given queries.
func FindQuestionContext(queries []database.Query) (question_context.QuestionContext, error) {
	entityMap, err := QuestionContextRepo.Find(question_context.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return question_context.QuestionContext{}, err
	}
	return question_context.Decode(entityMap), err
}

// GetAllQuestionContexts retrieves all QuestionContext entities.
func GetAllQuestionContexts() ([]question_context.QuestionContext, error) {
	entityMaps, err := QuestionContextRepo.ReadAll(question_context.GetCollectionPath())
	if err != nil {
		return []question_context.QuestionContext{}, err
	}
	return question_context.DecodeAll(entityMaps), nil
}


// GetAllQuestionContextsPaginated retrieves all QuestionContext entities in a paginated manner.
func GetAllQuestionContextsPaginated(pagination database.Pagination) ([]question_context.QuestionContext, database.Pagination, error) {
	entityMaps, pagination, err := QuestionContextRepo.ReadPaginated(question_context.GetCollectionPath(), pagination)
	if err != nil {
		return []question_context.QuestionContext{}, pagination, err
	}
	return question_context.DecodeAll(entityMaps), pagination, nil
}

// QueryQuestionContexts retrieves all QuestionContext entities according to given queries.
func QueryQuestionContexts(queries []database.Query, pagination database.Pagination) ([]question_context.QuestionContext, error) {
	entityMaps, err := QuestionContextRepo.Query(question_context.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []question_context.QuestionContext{}, err
	}
	return question_context.DecodeAll(entityMaps), nil
}

// QueryQuestionContextsGroup retrieves all QuestionContext entities according to given queries.
func QueryQuestionContextsGroup(queries []database.Query) ([]question_context.QuestionContext, error) {
	entityMaps, err := QuestionContextRepo.QueryGroup("question_contexts", queries)
	if err != nil {
		return []question_context.QuestionContext{}, err
	}
	return question_context.DecodeAll(entityMaps), nil
}

// CreateQuestionContext creates a new QuestionContext entity.
func CreateQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error) {
	return QuestionContextRepo.Create(&entity, editorID)
}

// UpdateQuestionContext updates an existing QuestionContext entity.
func UpdateQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error) {
	return QuestionContextRepo.Update(&entity, editorID)
}

// SaveQuestionContext saves a QuestionContext entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveQuestionContext(entity question_context.QuestionContext, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateQuestionContext(entity, editorID)
	} else {
		return UpdateQuestionContext(entity, editorID)
	}
}

// DeleteQuestionContext deletes a QuestionContext entity by its parents' path and questionContextID.
func DeleteQuestionContext(questionContextID string) error {
	return QuestionContextRepo.Delete(question_context.GetDocumentPath(questionContextID))
}
