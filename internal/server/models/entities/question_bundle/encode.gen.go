package question_bundle

import "encoding/json"

// ToMap converts a QuestionBundle struct to a map
func ToMap(e QuestionBundle) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a QuestionBundle struct to a JSON string
func ToJSON(e QuestionBundle) string {
    data, _ := json.Marshal(e)
    return string(data)
}