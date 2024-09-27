package examination_task_operations

import (
	examination_task "github.com/bakeable/bkry/internal/server/models/entities/examination_task"
)


func afterQueryGroup(entities []examination_task.ExaminationTask) []examination_task.ExaminationTask {
	return entities
}