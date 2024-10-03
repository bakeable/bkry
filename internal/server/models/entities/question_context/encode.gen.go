package question_context

import "encoding/json"

// ToMap converts a QuestionContext struct to a map
func ToMap(e QuestionContext) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a QuestionContext struct to a JSON string
func ToJSON(e QuestionContext) string {
    data, _ := json.Marshal(e)
    return string(data)
}