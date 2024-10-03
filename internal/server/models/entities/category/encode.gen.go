package category

import "encoding/json"

// ToMap converts a Category struct to a map
func ToMap(e Category) map[string]interface{} {
    data, _ := json.Marshal(e)
    var m map[string]interface{}
    json.Unmarshal(data, &m)
    return m
}

// ToJSON converts a Category struct to a JSON string
func ToJSON(e Category) string {
    data, _ := json.Marshal(e)
    return string(data)
}