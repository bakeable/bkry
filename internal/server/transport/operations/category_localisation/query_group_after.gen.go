package category_localisation_operations

import (
	category_localisation "github.com/bakeable/bkry/internal/server/models/entities/category_localisation"
)


func afterQueryGroup(entities []category_localisation.CategoryLocalisation) []category_localisation.CategoryLocalisation {
	return entities
}