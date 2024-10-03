package author

import "encoding/json"

// Decode converts a map to Author struct
func Decode(m interface{}) Author {
	var entity Author
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to Author struct
func DecodeJSON(s string) Author {
	var entity Author
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of Author struct
func DecodeAll(ms interface{}) []Author {
	var entities []Author
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
