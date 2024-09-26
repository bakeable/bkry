package examination_property_operations

import (
	examination_property "github.com/bakeable/bkry/data/entities/examination_property"
)

func afterGet(examinationPropertyID string, entity examination_property.ExaminationProperty) examination_property.ExaminationProperty {
	return entity
}