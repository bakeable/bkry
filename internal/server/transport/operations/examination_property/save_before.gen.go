package examination_property_operations

import (
	examination_property "github.com/bakeable/bkry/internal/server/models/entities/examination_property"
)

func beforeSave(entity examination_property.ExaminationProperty, editorID *string) examination_property.ExaminationProperty {
	// Return ExaminationProperty
	return entity
}