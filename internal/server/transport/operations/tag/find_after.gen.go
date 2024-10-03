package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)

func afterFind(entity tag.Tag) tag.Tag {
	return entity
}