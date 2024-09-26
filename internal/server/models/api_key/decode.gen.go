package api_key

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to ApiKey struct
func Decode(m interface{}) ApiKey {
	if m, ok := m.(map[string]interface{}); ok {
		return ApiKey{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Kind: utils.DecodeString(m["_kind"], "ApiKey"),
			publicKey: utils.DecodeString(m["publicKey"], ""),
			secretKey: utils.DecodeString(m["secretKey"], ""),
			userId: utils.DecodeString(m["userId"], ""),
		}
	}

	return ApiKey{}
}

// DecodeAll converts a slice of maps to a slice of ApiKey struct
func DecodeAll(ms interface{}) []ApiKey {
	var entities []ApiKey

	if arr, ok := ms.([]map[string]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	if arr, ok := ms.([]interface{}); ok {
		for _, m := range arr {
			entities = append(entities, Decode(m))
		}
		return entities
	}

	return entities
}

