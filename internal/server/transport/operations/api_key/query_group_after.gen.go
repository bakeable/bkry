package api_key_operations

import (
	api_key "github.com/bakeable/bkry/internal/server/models/entities/api_key"
)


func afterQueryGroup(entities []api_key.ApiKey) []api_key.ApiKey {
	return entities
}