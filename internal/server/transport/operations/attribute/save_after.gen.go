package attribute_operations

import (
	attribute "github.com/bakeable/bkry/internal/server/models/entities/attribute"
)

func afterSave(entity attribute.Attribute, editorID *string) {}