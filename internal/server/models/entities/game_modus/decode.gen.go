package game_modus

import "encoding/json"

// Decode converts a map to GameModus struct
func Decode(m interface{}) GameModus {
	var entity GameModus
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to GameModus struct
func DecodeJSON(s string) GameModus {
	var entity GameModus
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of GameModus struct
func DecodeAll(ms interface{}) []GameModus {
	var entities []GameModus
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
