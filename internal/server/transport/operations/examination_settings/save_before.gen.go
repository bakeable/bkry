package examination_settings_operations

import (
	examination_settings "github.com/bakeable/bkry/internal/server/models/entities/examination_settings"
)

func beforeSave(entity examination_settings.ExaminationSettings, editorID *string) examination_settings.ExaminationSettings {
	// Return ExaminationSettings
	return entity
}