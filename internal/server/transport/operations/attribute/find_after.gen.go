package attribute_operations

import (
	attribute "github.com/bakeable/bkry/data/entities/attribute"
)

func afterFind(entity attribute.Attribute) attribute.Attribute {
	return entity
}