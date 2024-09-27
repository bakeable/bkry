package attribute_operations

import (
	attribute "github.com/bakeable/bkry/internal/server/models/entities/attribute"
)

func beforeSave(entity attribute.Attribute, editorID *string) attribute.Attribute {
	// Return Attribute
	return entity
}