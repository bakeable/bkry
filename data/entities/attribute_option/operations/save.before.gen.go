package attribute_option_operations

import (
	attribute_option "github.com/bakeable/bkry/data/entities/attribute_option"
)

func beforeSave(entity attribute_option.AttributeOption, editorID *string) attribute_option.AttributeOption {
	// Return AttributeOption
	return entity
}