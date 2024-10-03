package question_bundle

import "encoding/json"

// Decode converts a map to QuestionBundle struct
func Decode(m interface{}) QuestionBundle {
	var entity QuestionBundle
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to QuestionBundle struct
func DecodeJSON(s string) QuestionBundle {
	var entity QuestionBundle
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of QuestionBundle struct
func DecodeAll(ms interface{}) []QuestionBundle {
	var entities []QuestionBundle
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
