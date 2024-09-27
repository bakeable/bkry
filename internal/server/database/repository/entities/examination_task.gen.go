package repo

import (
	"github.com/bakeable/bkry/internal/server/models/entities/examination_task"
	"github.com/bakeable/bkry/internal/server/database/repository"
	"github.com/bakeable/bkry/third_party/gcloud/datastore"
	
)

//// THIS FILE IS AUTO-GENERATED. DO NOT EDIT.

var ExaminationTaskRepo = repository.NewRepository[*examination_task.ExaminationTask]()

// GetExaminationTask retrieves a single ExaminationTask entity by its ID and examinationTaskID.
func GetExaminationTask(examinationTaskID string) (examination_task.ExaminationTask, error) {
	entityMap, err := ExaminationTaskRepo.Read(examination_task.GetDocumentPath( examinationTaskID))
	return examination_task.Decode(entityMap), err
}

// GetExaminationTaskOrNew retrieves a single ExaminationTask entity by its ID and examinationTaskID.
func GetExaminationTaskOrNew(examinationTaskID string) (examination_task.ExaminationTask, error) {
	entityMap, err := ExaminationTaskRepo.Read(examination_task.GetDocumentPath( examinationTaskID))
	if err != nil || entityMap == nil {
		return examination_task.ExaminationTask{}, err
	}
	return examination_task.Decode(entityMap), err
}

// GetExaminationTask retrieves a single ExaminationTask entity by its document path.
func GetExaminationTaskByPath(path string) (examination_task.ExaminationTask, error) {
	entityMap, err := ExaminationTaskRepo.Read(path)
	return examination_task.Decode(entityMap), err
}

// FindExaminationTask retrieves a ExaminationTask entity according to given queries.
func FindExaminationTask(queries []datastore.Query) (examination_task.ExaminationTask, error) {
	entityMap, err := ExaminationTaskRepo.Find(examination_task.GetCollectionPath(), queries)
	if err != nil || entityMap == nil {
		return examination_task.ExaminationTask{}, err
	}
	return examination_task.Decode(entityMap), err
}

// GetAllExaminationTasks retrieves all ExaminationTask entities.
func GetAllExaminationTasks() ([]examination_task.ExaminationTask, error) {
	entityMaps, err := ExaminationTaskRepo.ReadAll(examination_task.GetCollectionPath())
	if err != nil {
		return []examination_task.ExaminationTask{}, err
	}
	return examination_task.DecodeAll(entityMaps), nil
}


// GetAllExaminationTasksPaginated retrieves all ExaminationTask entities in a paginated manner.
func GetAllExaminationTasksPaginated(pagination datastore.Pagination) ([]examination_task.ExaminationTask, datastore.Pagination, error) {
	entityMaps, pagination, err := ExaminationTaskRepo.ReadPaginated(examination_task.GetCollectionPath(), pagination)
	if err != nil {
		return []examination_task.ExaminationTask{}, pagination, err
	}
	return examination_task.DecodeAll(entityMaps), pagination, nil
}

// QueryExaminationTasks retrieves all ExaminationTask entities according to given queries.
func QueryExaminationTasks(queries []datastore.Query, pagination datastore.Pagination) ([]examination_task.ExaminationTask, error) {
	entityMaps, err := ExaminationTaskRepo.Query(examination_task.GetCollectionPath(), queries, pagination)
	if err != nil {
		return []examination_task.ExaminationTask{}, err
	}
	return examination_task.DecodeAll(entityMaps), nil
}

// QueryExaminationTasksGroup retrieves all ExaminationTask entities according to given queries.
func QueryExaminationTasksGroup(queries []datastore.Query) ([]examination_task.ExaminationTask, error) {
	entityMaps, err := ExaminationTaskRepo.QueryGroup("examination_tasks", queries)
	if err != nil {
		return []examination_task.ExaminationTask{}, err
	}
	return examination_task.DecodeAll(entityMaps), nil
}

// CreateExaminationTask creates a new ExaminationTask entity.
func CreateExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error) {
	return ExaminationTaskRepo.Create(&entity, editorID)
}

// UpdateExaminationTask updates an existing ExaminationTask entity.
func UpdateExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error) {
	return ExaminationTaskRepo.Update(&entity, editorID)
}

// SaveExaminationTask saves a ExaminationTask entity. If the entity has an ID, it updates the existing entity; otherwise, it creates a new entity.
func SaveExaminationTask(entity examination_task.ExaminationTask, editorID *string) (string, error) {
	if entity.ID == "" {
		return CreateExaminationTask(entity, editorID)
	} else {
		return UpdateExaminationTask(entity, editorID)
	}
}

// DeleteExaminationTask deletes a ExaminationTask entity by its parents' path and examinationTaskID.
func DeleteExaminationTask(examinationTaskID string) error {
	return ExaminationTaskRepo.Delete(examination_task.GetDocumentPath(examinationTaskID))
}
