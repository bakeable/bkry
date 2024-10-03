package tag

import "encoding/json"

// ToMap converts a Tag struct to a map
func ToMap(e Tag) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a Tag struct to a JSON string
func ToJSON(e Tag) string {
    data, _ := json.Marshal(e)
    return string(data)
}