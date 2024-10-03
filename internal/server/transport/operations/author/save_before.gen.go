package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)

func beforeSave(entity author.Author, editorID *string) author.Author {
	// Return Author
	return entity
}