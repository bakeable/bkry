package media

import "encoding/json"

// Decode converts a map to Media struct
func Decode(m interface{}) Media {
	var entity Media
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to Media struct
func DecodeJSON(s string) Media {
	var entity Media
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of Media struct
func DecodeAll(ms interface{}) []Media {
	var entities []Media
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
