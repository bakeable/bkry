package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
)

func beforeSave(entity attribute_option_set.AttributeOptionSet, editorID *string) attribute_option_set.AttributeOptionSet {
	// Return AttributeOptionSet
	return entity
}