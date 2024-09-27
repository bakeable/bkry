package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
)


func afterGetAllPaginated(entities []attribute_option_set.AttributeOptionSet, pageSize int, pageNumber int, orderBy string, ascending bool) []attribute_option_set.AttributeOptionSet {
	return entities
}