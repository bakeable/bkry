package attribute_operations

import (
	attribute "github.com/bakeable/bkry/internal/server/models/entities/attribute"
)


func afterGetAll(entities []attribute.Attribute) []attribute.Attribute {
	return entities
}