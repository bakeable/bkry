package attribute_option_operations

import (
	attribute_option "github.com/bakeable/bkry/internal/server/models/entities/attribute_option"
)


func afterGetAllPaginated(entities []attribute_option.AttributeOption, pageSize int, pageNumber int, orderBy string, ascending bool) []attribute_option.AttributeOption {
	return entities
}