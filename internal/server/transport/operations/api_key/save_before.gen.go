package api_key_operations

import (
	api_key "github.com/bakeable/bkry/internal/server/models/entities/api_key"
)

func beforeSave(entity api_key.ApiKey, editorID *string) api_key.ApiKey {
	// Return ApiKey
	return entity
}