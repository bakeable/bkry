package examination_property_operations

import (
	examination_property "github.com/bakeable/bkry/internal/server/models/entities/examination_property"
)

func afterFind(entity examination_property.ExaminationProperty) examination_property.ExaminationProperty {
	return entity
}