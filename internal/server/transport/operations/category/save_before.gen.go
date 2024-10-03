package category_operations

import (
	category "github.com/bakeable/bkry/internal/server/models/entities/category"
)

func beforeSave(entity category.Category, editorID *string) category.Category {
	// Return Category
	return entity
}