package media_operations

import (
	media "github.com/bakeable/bkry/internal/server/models/entities/media"
)

func beforeSave(entity media.Media, editorID *string) media.Media {
	// Return Media
	return entity
}