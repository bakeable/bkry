package question_context

import "encoding/json"

// Decode converts a map to QuestionContext struct
func Decode(m interface{}) QuestionContext {
	var entity QuestionContext
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to QuestionContext struct
func DecodeJSON(s string) QuestionContext {
	var entity QuestionContext
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of QuestionContext struct
func DecodeAll(ms interface{}) []QuestionContext {
	var entities []QuestionContext
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
