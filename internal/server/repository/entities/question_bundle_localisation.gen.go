package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/question_bundle_localisation"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/internal/server/database"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var QuestionBundleLocalisationRepo = repository.NewFirestoreRepository[*question_bundle_localisation.QuestionBundleLocalisation]()

// GetQuestionBundleLocalisation retrieves a single QuestionBundleLocalisation entity by its ID and questionBundleLocalisationID.
func GetQuestionBundleLocalisation(QuestionBundleID string, questionBundleLocalisationID string) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMap, err := QuestionBundleLocalisationRepo.Read(question_bundle_localisation.GetDocumentPath(QuestionBundleID,  questionBundleLocalisationID))
	return question_bundle_localisation.Decode(entityMap), err
}

// GetQuestionBundleLocalisationOrNew retrieves a single QuestionBundleLocalisation entity by its ID and questionBundleLocalisationID.
func GetQuestionBundleLocalisationOrNew(QuestionBundleID string, questionBundleLocalisationID string) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMap, err := QuestionBundleLocalisationRepo.Read(question_bundle_localisation.GetDocumentPath(QuestionBundleID,  questionBundleLocalisationID))
	if err != nil || entityMap == nil {
		return question_bundle_localisation.QuestionBundleLocalisation{}, err
	}
	return question_bundle_localisation.Decode(entityMap), err
}

// GetQuestionBundleLocalisation retrieves a single QuestionBundleLocalisation entity by its document path.
func GetQuestionBundleLocalisationByPath(path string) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMap, err := QuestionBundleLocalisationRepo.Read(path)
	return question_bundle_localisation.Decode(entityMap), err
}

// FindQuestionBundleLocalisation retrieves a QuestionBundleLocalisation entity according to given queries.
func FindQuestionBundleLocalisation(questionBundleID string, queries []database.Query) (question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMap, err := QuestionBundleLocalisationRepo.Find(question_bundle_localisation.GetCollectionPath(questionBundleID), queries)
	if err != nil || entityMap == nil {
		return question_bundle_localisation.QuestionBundleLocalisation{}, err
	}
	return question_bundle_localisation.Decode(entityMap), err
}

// GetAllQuestionBundleLocalisations retrieves all QuestionBundleLocalisation entities.
func GetAllQuestionBundleLocalisations(questionBundleID string) ([]question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMaps, err := QuestionBundleLocalisationRepo.ReadAll(question_bundle_localisation.GetCollectionPath(questionBundleID))
	if err != nil {
		return []question_bundle_localisation.QuestionBundleLocalisation{}, err
	}
	return question_bundle_localisation.DecodeAll(entityMaps), nil
}


// GetAllQuestionBundleLocalisationsPaginated retrieves all QuestionBundleLocalisation entities in a paginated manner.
func GetAllQuestionBundleLocalisationsPaginated(questionBundleID string, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, database.Pagination, error) {
	entityMaps, pagination, err := QuestionBundleLocalisationRepo.ReadPaginated(question_bundle_localisation.GetCollectionPath(questionBundleID), pagination)
	if err != nil {
		return []question_bundle_localisation.QuestionBundleLocalisation{}, pagination, err
	}
	return question_bundle_localisation.DecodeAll(entityMaps), pagination, nil
}

// QueryQuestionBundleLocalisations retrieves all QuestionBundleLocalisation entities according to given queries.
func QueryQuestionBundleLocalisations(questionBundleID string, queries []database.Query, pagination database.Pagination) ([]question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMaps, err := QuestionBundleLocalisationRepo.Query(question_bundle_localisation.GetCollectionPath(questionBundleID), queries, pagination)
	if err != nil {
		return []question_bundle_localisation.QuestionBundleLocalisation{}, err
	}
	return question_bundle_localisation.DecodeAll(entityMaps), nil
}

// QueryQuestionBundleLocalisationsGroup retrieves all QuestionBundleLocalisation entities according to given queries.
func QueryQuestionBundleLocalisationsGroup(queries []database.Query) ([]question_bundle_localisation.QuestionBundleLocalisation, error) {
	entityMaps, err := QuestionBundleLocalisationRepo.QueryGroup("question_bundle_localisations", queries)
	if err != nil {
		return []question_bundle_localisation.QuestionBundleLocalisation{}, err
	}
	return question_bundle_localisation.DecodeAll(entityMaps), nil
}

// CreateQuestionBundleLocalisation creates a new QuestionBundleLocalisation entity.
func CreateQuestionBundleLocalisation(QuestionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error) {
	entity.QuestionBundleID = QuestionBundleID
	return QuestionBundleLocalisationRepo.Create(&entity, editorID)
}

// UpdateQuestionBundleLocalisation updates an existing QuestionBundleLocalisation entity.
func UpdateQuestionBundleLocalisation(QuestionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error) {
	entity.QuestionBundleID = QuestionBundleID
	return QuestionBundleLocalisationRepo.Update(&entity, editorID)
}

// SaveQuestionBundleLocalisation saves a QuestionBundleLocalisation entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveQuestionBundleLocalisation(QuestionBundleID string, entity question_bundle_localisation.QuestionBundleLocalisation, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateQuestionBundleLocalisation(QuestionBundleID, entity, editorID)
	} else {
		return UpdateQuestionBundleLocalisation(QuestionBundleID, entity, editorID)
	}
}

// DeleteQuestionBundleLocalisation deletes a QuestionBundleLocalisation entity by its parents' path and questionBundleLocalisationID.
func DeleteQuestionBundleLocalisation(QuestionBundleID string, questionBundleLocalisationID string) error {
	return QuestionBundleLocalisationRepo.Delete(question_bundle_localisation.GetDocumentPath(QuestionBundleID, questionBundleLocalisationID))
}
