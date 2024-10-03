package author_operations

import (
	author "github.com/bakeable/bkry/internal/server/models/entities/author"
)

func afterSave(entity author.Author, editorID *string) {}