package attribute_option_set_operations

import (
	attribute_option_set "github.com/bakeable/bkry/internal/server/models/entities/attribute_option_set"
)

func afterGet(attributeOptionSetID string, entity attribute_option_set.AttributeOptionSet) attribute_option_set.AttributeOptionSet {
	return entity
}