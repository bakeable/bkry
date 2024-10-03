package question_bundle_localisation

import "encoding/json"

// Decode converts a map to QuestionBundleLocalisation struct
func Decode(m interface{}) QuestionBundleLocalisation {
	var entity QuestionBundleLocalisation
	data, err := json.Marshal(m)
	if err != nil {
		return entity
	}
	json.Unmarshal(data, &entity)
	return entity
}

// DecodeJSON converts a JSON string to QuestionBundleLocalisation struct
func DecodeJSON(s string) QuestionBundleLocalisation {
	var entity QuestionBundleLocalisation
	json.Unmarshal([]byte(s), &entity)
	return entity
}

// DecodeAll converts a slice of interfaces to a slice of QuestionBundleLocalisation struct
func DecodeAll(ms interface{}) []QuestionBundleLocalisation {
	var entities []QuestionBundleLocalisation
	data, err := json.Marshal(ms)
	if err != nil {
		return entities
	}
	json.Unmarshal(data, &entities)
	return entities
}
