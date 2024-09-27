package api_key_operations

import (
	api_key "github.com/bakeable/bkry/internal/server/models/entities/api_key"
)

func afterGet(apiKeyID string, entity api_key.ApiKey) api_key.ApiKey {
	return entity
}