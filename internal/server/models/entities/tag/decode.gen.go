package tag

import "encoding/json"

// Decode converts a map to Tag struct
func Decode(m interface{}) Tag {
	var entity Tag
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to Tag struct
func DecodeJSON(s string) Tag {
	var entity Tag
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of Tag struct
func DecodeAll(ms interface{}) []Tag {
	var entities []Tag
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
