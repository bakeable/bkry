package question_bundle_localisation

import "encoding/json"

// ToMap converts a QuestionBundleLocalisation struct to a map
func ToMap(e QuestionBundleLocalisation) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a QuestionBundleLocalisation struct to a JSON string
func ToJSON(e QuestionBundleLocalisation) string {
    data, _ := json.Marshal(e)
    return string(data)
}