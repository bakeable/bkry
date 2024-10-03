package game_modus_localisation

import "encoding/json"

// Decode converts a map to GameModusLocalisation struct
func Decode(m interface{}) GameModusLocalisation {
	var entity GameModusLocalisation
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to GameModusLocalisation struct
func DecodeJSON(s string) GameModusLocalisation {
	var entity GameModusLocalisation
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of GameModusLocalisation struct
func DecodeAll(ms interface{}) []GameModusLocalisation {
	var entities []GameModusLocalisation
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
