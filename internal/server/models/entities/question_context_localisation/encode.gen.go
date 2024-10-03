package question_context_localisation

import "encoding/json"

// ToMap converts a QuestionContextLocalisation struct to a map
func ToMap(e QuestionContextLocalisation) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a QuestionContextLocalisation struct to a JSON string
func ToJSON(e QuestionContextLocalisation) string {
    data, _ := json.Marshal(e)
    return string(data)
}