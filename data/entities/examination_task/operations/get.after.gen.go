package examination_task_operations

import (
	examination_task "github.com/bakeable/bkry/data/entities/examination_task"
)

func afterGet(examinationTaskID string, entity examination_task.ExaminationTask) examination_task.ExaminationTask {
	return entity
}