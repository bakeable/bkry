package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)

func afterSave(entity media.Media, editorID *string) {}