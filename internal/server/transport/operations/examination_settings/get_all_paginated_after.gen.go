package examination_settings_operations

import (
	examination_settings "github.com/bakeable/bkry/internal/server/models/entities/examination_settings"
)


func afterGetAllPaginated(entities []examination_settings.ExaminationSettings, pageSize int, pageNumber int, orderBy string, ascending bool) []examination_settings.ExaminationSettings {
	return entities
}