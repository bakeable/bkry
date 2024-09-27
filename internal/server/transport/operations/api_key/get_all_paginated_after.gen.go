package api_key_operations

import (
	api_key "github.com/bakeable/bkry/internal/server/models/entities/api_key"
)


func afterGetAllPaginated(entities []api_key.ApiKey, pageSize int, pageNumber int, orderBy string, ascending bool) []api_key.ApiKey {
	return entities
}