package examination_task_operations

import (
	examination_task "github.com/bakeable/bkry/data/entities/examination_task"
)

func beforeSave(entity examination_task.ExaminationTask, editorID *string) examination_task.ExaminationTask {
	// Return ExaminationTask
	return entity
}