package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)

func afterFind(entity author.Author) author.Author {
	return entity
}