package media_operations

import (
	media "github.com/bakeable/bkry/data/entities/media"
)

func afterFind(entity media.Media) media.Media {
	return entity
}