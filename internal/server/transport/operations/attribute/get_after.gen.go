package attribute_operations

import (
	attribute "github.com/bakeable/bkry/internal/server/models/entities/attribute"
)

func afterGet(attributeID string, entity attribute.Attribute) attribute.Attribute {
	return entity
}