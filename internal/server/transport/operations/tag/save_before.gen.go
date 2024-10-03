package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)

func beforeSave(entity tag.Tag, editorID *string) tag.Tag {
	// Return Tag
	return entity
}