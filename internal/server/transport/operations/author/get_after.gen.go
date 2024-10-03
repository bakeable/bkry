package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)

func afterGet(authorID string, entity author.Author) author.Author {
	return entity
}