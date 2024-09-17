package email

import (
	"github.com/bakeable/bkry/data/general/auditinfo"
	utils "github.com/bakeable/bkry/tools"
)

// Decode converts a map to Email struct
func Decode(m interface{}) Email {
	if m, ok := m.(map[string]interface{}); ok {
		return Email{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			Kind: utils.DecodeString(m["_kind"], "Email"),
		}
	}

	return Email{}
}

// DecodeAll converts a slice of maps to a slice of Email struct
func DecodeAll(ms interface{}) []Email {
	var entities []Email

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

