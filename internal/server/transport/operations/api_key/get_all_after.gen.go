package api_key_operations

import (
	api_key "github.com/bakeable/bkry/data/entities/api_key"
)


func afterGetAll(entities []api_key.ApiKey) []api_key.ApiKey {
	return entities
}