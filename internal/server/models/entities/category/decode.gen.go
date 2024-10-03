package category

import "encoding/json"

// Decode converts a map to Category struct
func Decode(m interface{}) Category {
	var entity Category
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to Category struct
func DecodeJSON(s string) Category {
	var entity Category
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of Category struct
func DecodeAll(ms interface{}) []Category {
	var entities []Category
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
