package attribute_operations

import (
	attribute "github.com/bakeable/bkry/internal/server/models/entities/attribute"
)


func afterGetAllPaginated(entities []attribute.Attribute, pageSize int, pageNumber int, orderBy string, ascending bool) []attribute.Attribute {
	return entities
}