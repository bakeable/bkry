package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/question_context_localisation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var QuestionContextLocalisationRepo = repository.NewFirestoreRepository[*question_context_localisation.QuestionContextLocalisation]()

// GetQuestionContextLocalisation retrieves a single QuestionContextLocalisation entity by its ID and questionContextLocalisationID.
func GetQuestionContextLocalisation(QuestionContextID string, questionContextLocalisationID string) (question_context_localisation.QuestionContextLocalisation, error) {
	entityMap, err := QuestionContextLocalisationRepo.Read(question_context_localisation.GetDocumentPath(QuestionContextID,  questionContextLocalisationID))
	return question_context_localisation.Decode(entityMap), err
}

// GetQuestionContextLocalisationOrNew retrieves a single QuestionContextLocalisation entity by its ID and questionContextLocalisationID.
func GetQuestionContextLocalisationOrNew(QuestionContextID string, questionContextLocalisationID string) (question_context_localisation.QuestionContextLocalisation, error) {
	entityMap, err := QuestionContextLocalisationRepo.Read(question_context_localisation.GetDocumentPath(QuestionContextID,  questionContextLocalisationID))
	if err != nil || entityMap == nil {
		return question_context_localisation.QuestionContextLocalisation{}, err
	}
	return question_context_localisation.Decode(entityMap), err
}

// GetQuestionContextLocalisation retrieves a single QuestionContextLocalisation entity by its document path.
func GetQuestionContextLocalisationByPath(path string) (question_context_localisation.QuestionContextLocalisation, error) {
	entityMap, err := QuestionContextLocalisationRepo.Read(path)
	return question_context_localisation.Decode(entityMap), err
}

// FindQuestionContextLocalisation retrieves a QuestionContextLocalisation entity according to given queries.
func FindQuestionContextLocalisation(questionContextID string, queries []database.Query) (question_context_localisation.QuestionContextLocalisation, error) {
	entityMap, err := QuestionContextLocalisationRepo.Find(question_context_localisation.GetCollectionPath(questionContextID), queries)
	if err != nil || entityMap == nil {
		return question_context_localisation.QuestionContextLocalisation{}, err
	}
	return question_context_localisation.Decode(entityMap), err
}

// GetAllQuestionContextLocalisations retrieves all QuestionContextLocalisation entities.
func GetAllQuestionContextLocalisations(questionContextID string) ([]question_context_localisation.QuestionContextLocalisation, error) {
	entityMaps, err := QuestionContextLocalisationRepo.ReadAll(question_context_localisation.GetCollectionPath(questionContextID))
	if err != nil {
		return []question_context_localisation.QuestionContextLocalisation{}, err
	}
	return question_context_localisation.DecodeAll(entityMaps), nil
}


// GetAllQuestionContextLocalisationsPaginated retrieves all QuestionContextLocalisation entities in a paginated manner.
func GetAllQuestionContextLocalisationsPaginated(questionContextID string, pagination database.Pagination) ([]question_context_localisation.QuestionContextLocalisation, database.Pagination, error) {
	entityMaps, pagination, err := QuestionContextLocalisationRepo.ReadPaginated(question_context_localisation.GetCollectionPath(questionContextID), pagination)
	if err != nil {
		return []question_context_localisation.QuestionContextLocalisation{}, pagination, err
	}
	return question_context_localisation.DecodeAll(entityMaps), pagination, nil
}

// QueryQuestionContextLocalisations retrieves all QuestionContextLocalisation entities according to given queries.
func QueryQuestionContextLocalisations(questionContextID string, queries []database.Query, pagination database.Pagination) ([]question_context_localisation.QuestionContextLocalisation, error) {
	entityMaps, err := QuestionContextLocalisationRepo.Query(question_context_localisation.GetCollectionPath(questionContextID), queries, pagination)
	if err != nil {
		return []question_context_localisation.QuestionContextLocalisation{}, err
	}
	return question_context_localisation.DecodeAll(entityMaps), nil
}

// QueryQuestionContextLocalisationsGroup retrieves all QuestionContextLocalisation entities according to given queries.
func QueryQuestionContextLocalisationsGroup(queries []database.Query) ([]question_context_localisation.QuestionContextLocalisation, error) {
	entityMaps, err := QuestionContextLocalisationRepo.QueryGroup("question_context_localisations", queries)
	if err != nil {
		return []question_context_localisation.QuestionContextLocalisation{}, err
	}
	return question_context_localisation.DecodeAll(entityMaps), nil
}

// CreateQuestionContextLocalisation creates a new QuestionContextLocalisation entity.
func CreateQuestionContextLocalisation(QuestionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error) {
	entity.QuestionContextID = QuestionContextID
	return QuestionContextLocalisationRepo.Create(&entity, editorID)
}

// UpdateQuestionContextLocalisation updates an existing QuestionContextLocalisation entity.
func UpdateQuestionContextLocalisation(QuestionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error) {
	entity.QuestionContextID = QuestionContextID
	return QuestionContextLocalisationRepo.Update(&entity, editorID)
}

// SaveQuestionContextLocalisation saves a QuestionContextLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveQuestionContextLocalisation(QuestionContextID string, entity question_context_localisation.QuestionContextLocalisation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateQuestionContextLocalisation(QuestionContextID, entity, editorID)
	} else {
		return UpdateQuestionContextLocalisation(QuestionContextID, entity, editorID)
	}
}

// DeleteQuestionContextLocalisation deletes a QuestionContextLocalisation entity by its parents' path and questionContextLocalisationID.
func DeleteQuestionContextLocalisation(QuestionContextID string, questionContextLocalisationID string) error {
	return QuestionContextLocalisationRepo.Delete(question_context_localisation.GetDocumentPath(QuestionContextID, questionContextLocalisationID))
}
