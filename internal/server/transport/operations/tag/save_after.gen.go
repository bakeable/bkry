package tag_operations

import (
	tag "github.com/bakeable/bkry/internal/server/models/entities/tag"
)

func afterSave(entity tag.Tag, editorID *string) {}