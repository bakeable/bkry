package tag_localisation

import "encoding/json"

// Decode converts a map to TagLocalisation struct
func Decode(m interface{}) TagLocalisation {
	var entity TagLocalisation
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to TagLocalisation struct
func DecodeJSON(s string) TagLocalisation {
	var entity TagLocalisation
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of TagLocalisation struct
func DecodeAll(ms interface{}) []TagLocalisation {
	var entities []TagLocalisation
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
