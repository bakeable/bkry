package category_operations

import (
	category "github.com/bakeable/bkry/internal/server/models/entities/category"
)

func afterGet(categoryID string, entity category.Category) category.Category {
	return entity
}