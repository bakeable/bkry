package attribute_option_operations

import (
	attribute_option "github.com/bakeable/bkry/internal/server/models/entities/attribute_option"
)


func afterGetAll(entities []attribute_option.AttributeOption) []attribute_option.AttributeOption {
	return entities
}