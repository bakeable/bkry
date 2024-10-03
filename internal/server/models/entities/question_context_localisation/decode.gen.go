package question_context_localisation

import "encoding/json"

// Decode converts a map to QuestionContextLocalisation struct
func Decode(m interface{}) QuestionContextLocalisation {
	var entity QuestionContextLocalisation
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to QuestionContextLocalisation struct
func DecodeJSON(s string) QuestionContextLocalisation {
	var entity QuestionContextLocalisation
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of QuestionContextLocalisation struct
func DecodeAll(ms interface{}) []QuestionContextLocalisation {
	var entities []QuestionContextLocalisation
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
