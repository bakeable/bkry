package attribute_operations

import (
	attribute "github.com/bakeable/bkry/data/entities/attribute"
)

func beforeSave(entity attribute.Attribute, editorID *string) attribute.Attribute {
	// Return Attribute
	return entity
}