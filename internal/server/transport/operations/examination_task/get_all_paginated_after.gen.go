package examination_task_operations

import (
	examination_task "github.com/bakeable/bkry/internal/server/models/entities/examination_task"
)


func afterGetAllPaginated(entities []examination_task.ExaminationTask, pageSize int, pageNumber int, orderBy string, ascending bool) []examination_task.ExaminationTask {
	return entities
}