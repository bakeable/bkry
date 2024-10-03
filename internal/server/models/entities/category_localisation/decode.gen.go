package category_localisation

import "encoding/json"

// Decode converts a map to CategoryLocalisation struct
func Decode(m interface{}) CategoryLocalisation {
	var entity CategoryLocalisation
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to CategoryLocalisation struct
func DecodeJSON(s string) CategoryLocalisation {
	var entity CategoryLocalisation
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of CategoryLocalisation struct
func DecodeAll(ms interface{}) []CategoryLocalisation {
	var entities []CategoryLocalisation
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
