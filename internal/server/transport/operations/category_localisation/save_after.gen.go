package category_localisation_operations

import (
	category_localisation "github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
)

func afterSave(categoryID string, entity category_localisation.CategoryLocalisation, editorID *string) {}