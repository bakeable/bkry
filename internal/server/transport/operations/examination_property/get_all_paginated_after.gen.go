package examination_property_operations

import (
	examination_property "github.com/bakeable/bkry/data/entities/examination_property"
)


func afterGetAllPaginated(entities []examination_property.ExaminationProperty, pageSize int, pageNumber int, orderBy string, ascending bool) []examination_property.ExaminationProperty {
	return entities
}