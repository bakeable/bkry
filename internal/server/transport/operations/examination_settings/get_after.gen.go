package examination_settings_operations

import (
	examination_settings "github.com/bakeable/bkry/internal/server/models/entities/examination_settings"
)

func afterGet(examinationSettingsID string, entity examination_settings.ExaminationSettings) examination_settings.ExaminationSettings {
	return entity
}