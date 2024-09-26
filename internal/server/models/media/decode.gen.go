package media

import (
	"github.com/bakeable/bkry/tools"
	"github.com/bakeable/bkry/data/general/auditinfo"
)

// Decode converts a map to Media struct
func Decode(m interface{}) Media {
	if m, ok := m.(map[string]interface{}); ok {
		return Media{
			ID: utils.TryDecodeString(m["id"]),
			Created: auditinfo.Decode(m["created"]),
			Modified: auditinfo.Decode(m["modified"]),
			ContentType: ContentType(utils.DecodeString(m["contentType"], "Image")),
			Description: utils.DecodeString(m["description"], ""),
			Extension: utils.DecodeString(m["extension"], ""),
			Filename: utils.DecodeString(m["filename"], ""),
			MimeType: utils.DecodeString(m["mimeType"], ""),
			Size: utils.DecodeString(m["size"], ""),
			StoragePath: utils.DecodeString(m["storagePath"], ""),
			Url: utils.DecodeString(m["url"], ""),
		}
	}

	return Media{}
}

// DecodeAll converts a slice of maps to a slice of Media struct
func DecodeAll(ms interface{}) []Media {
	var entities []Media

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

