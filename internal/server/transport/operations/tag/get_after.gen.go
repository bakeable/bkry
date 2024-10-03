package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)

func afterGet(tagID string, entity tag.Tag) tag.Tag {
	return entity
}